package linter

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/VKCOM/noverify/src/baseline"
	"github.com/VKCOM/noverify/src/constfold"
	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/ir/irutil"
	"github.com/VKCOM/noverify/src/meta"
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/phpdoc"
	"github.com/VKCOM/noverify/src/phpgrep"
	"github.com/VKCOM/noverify/src/quickfix"
	"github.com/VKCOM/noverify/src/rules"
	"github.com/VKCOM/noverify/src/solver"
	"github.com/VKCOM/noverify/src/state"
	"github.com/VKCOM/noverify/src/workspace"
)

const (
	maxFunctionLines = 150
)

// rootWalker is used to analyze root scope. Mostly defines, function and class definitions are analyzed.
type rootWalker struct {
	// An associated file that is traversed by the current Root Walker.
	file *workspace.File

	custom      []RootChecker
	customBlock []BlockCheckerCreateFunc
	customState map[string]interface{}

	rootRset  *rules.ScopedSet
	localRset *rules.ScopedSet
	anyRset   *rules.ScopedSet

	ctx rootContext

	// nodeSet is a reusable node set for both root and block walkers.
	nodeSet irutil.NodeSet

	reSimplifier *regexpSimplifier
	reVet        *regexpVet

	// internal state
	meta fileMeta

	currentClassNode ir.Node

	allowDisabledRegexp *regexp.Regexp // user-defined flag that files suitable for this regular expression should not be linted
	linterDisabled      bool           // flag indicating whether linter is disabled. Flag is set to true only if the file
	// name matches the pattern and @linter disable was encountered

	// strictTypes is true if file contains `declare(strict_types=1)`.
	strictTypes bool

	reports []*Report

	config *Config
}

type phpDocParamEl struct {
	optional bool
	typ      meta.TypesMap
}

type phpDocParamsMap map[string]phpDocParamEl

// InitCustom is needed to initialize walker state
func (d *rootWalker) InitCustom() {
	d.custom = nil
	for _, createFn := range customRootLinters {
		d.custom = append(d.custom, createFn(&RootContext{w: d}))
	}

	d.customBlock = customBlockLinters
}

// scope returns root-level variable scope if applicable.
func (d *rootWalker) scope() *meta.Scope {
	if d.meta.Scope == nil {
		d.meta.Scope = meta.NewScope()
	}
	return d.meta.Scope
}

// state allows for custom hooks to store state between entering root context and block context.
func (d *rootWalker) state() map[string]interface{} {
	if d.customState == nil {
		d.customState = make(map[string]interface{})
	}
	return d.customState
}

// File returns file for current root walker.
func (d *rootWalker) File() *workspace.File {
	return d.file
}

// EnterNode is invoked at every node in hierarchy
func (d *rootWalker) EnterNode(n ir.Node) (res bool) {
	res = true

	for _, c := range d.custom {
		c.BeforeEnterNode(n)
	}

	if ffs := n.GetFreeFloating(); ffs != nil {
		for _, cs := range *ffs {
			for _, c := range cs {
				if c.StringType == freefloating.CommentType {
					d.handleComment(c)
				}
			}
		}
	}

	if _, ok := n.(*ir.AnonClassExpr); ok {
		// TODO: remove when #62 and anon class support in general is ready.
		return false // Don't walk nor enter anon classes
	}

	state.EnterNode(d.ctx.st, n)

	switch n := n.(type) {
	case *ir.DeclareStmt:
		for _, c := range n.Consts {
			c, ok := c.(*ir.ConstantStmt)
			if !ok {
				continue
			}
			if c.ConstantName.Value == "strict_types" {
				v, ok := c.Expr.(*ir.Lnumber)
				if ok && v.Value == "1" {
					d.strictTypes = true
				}
			}
		}

	case *ir.InterfaceStmt:
		d.currentClassNode = n

		if meta.IsIndexingComplete() {
			d.checkInterface(n)
		} else {
			d.handleInterface(n)
		}

	case *ir.ClassStmt:
		d.currentClassNode = n

		if meta.IsIndexingComplete() {
			d.checkClass(n)
		} else {
			d.handleClass(n)
		}

	case *ir.PropertyListStmt:
		if meta.IsIndexingComplete() {
			d.checkPropertyList(n)
		} else {
			d.handlePropertyList(n)
		}

	case *ir.ClassConstListStmt:
		if meta.IsIndexingComplete() {
			d.checkClassConstList(n)
		} else {
			d.handleClassConstList(n)
		}

	case *ir.ClassMethodStmt:
		if meta.IsIndexingComplete() {
			d.checkClassMethod(n)
		} else {
			d.handleClassMethod(n)
		}
		res = false

	case *ir.TraitStmt:
		d.currentClassNode = n

		if meta.IsIndexingComplete() {
			d.checkTrait(n)
		} else {
			d.handleTrait(n)
		}

	case *ir.TraitUseStmt:
		if meta.IsIndexingComplete() {
			d.checkTraitUse(n)
		} else {
			d.handleTraitUse(n)
		}

	case *ir.Assign:
		v, ok := n.Variable.(*ir.SimpleVar)
		if !ok {
			break
		}

		d.scope().AddVar(v, solver.ExprTypeLocal(d.scope(), d.ctx.st, n.Expression), "global variable", meta.VarAlwaysDefined)
	case *ir.FunctionStmt:
		res = d.enterFunction(n)
		d.checkKeywordCase(n, "function")
	case *ir.FunctionCallExpr:
		res = d.enterFunctionCall(n)
	case *ir.ConstListStmt:
		res = d.enterConstList(n)

	case *ir.NamespaceStmt:
		d.checkKeywordCase(n, "namespace")
	}

	for _, c := range d.custom {
		c.AfterEnterNode(n)
	}

	if meta.IsIndexingComplete() && d.rootRset != nil {
		kind := ir.GetNodeKind(n)
		d.runRules(n, d.scope(), d.rootRset.RulesByKind[kind])
	}

	if !res {
		// If we're not returning true from this method,
		// LeaveNode will not be called for this node.
		// But we still need to "leave" them if they
		// were entered in the ClassParseState.
		state.LeaveNode(d.ctx.st, n)
	}
	return res
}

func (d *rootWalker) handleTrait(n *ir.TraitStmt) {
	if d.meta.Traits.H == nil {
		d.meta.Traits = meta.NewClassesMap()
	}

	name := n.TraitName.Value
	namespace := d.ctx.st.Namespace
	fullName := namespace + `\` + name

	trait := meta.ClassInfo{
		Pos:              d.getElementPos(d.currentClassNode),
		Name:             fullName,
		Parent:           "",
		ParentInterfaces: d.ctx.st.CurrentParentInterfaces,
		Methods:          meta.NewFunctionsMap(),
		Properties:       make(meta.PropertiesMap),
		Traits:           make(map[string]struct{}),
		Constants:        nil, // not need
		Interfaces:       nil, // not need
	}

	d.meta.Traits.Set(fullName, trait)
}

func (d *rootWalker) handleTraitUse(n *ir.TraitUseStmt) {
	class, ok := d.getCurrentClass()
	if !ok {
		return
	}

	for _, tr := range n.Traits {
		traitName, ok := solver.GetClassName(d.ctx.st, tr)
		if !ok {
			continue
		}

		class.Traits[traitName] = struct{}{}
	}
}

func (d *rootWalker) handleInterface(n *ir.InterfaceStmt) {
	if d.meta.Classes.H == nil {
		d.meta.Classes = meta.NewClassesMap()
	}

	name := n.InterfaceName.Value
	namespace := d.ctx.st.Namespace
	fullName := namespace + `\` + name

	extends := d.handleInterfaceExtends(n)

	iface := meta.ClassInfo{
		Pos:              d.getElementPos(d.currentClassNode),
		Name:             fullName,
		Parent:           "",
		ParentInterfaces: extends,
		Methods:          meta.NewFunctionsMap(),
		Constants:        make(meta.ConstantsMap),
		Interfaces:       nil, // not need
		Traits:           nil, // not need
		Properties:       nil, // not need
	}

	d.meta.Classes.Set(fullName, iface)
}

func (d *rootWalker) handleInterfaceExtends(n *ir.InterfaceStmt) []string {
	var extends []string
	if n.Extends != nil {
		for _, iface := range n.Extends.InterfaceNames {
			ifaceName, ok := solver.GetClassName(d.ctx.st, iface)
			if !ok {
				continue
			}

			extends = append(extends, ifaceName)
		}
	}
	return extends
}

func (d *rootWalker) handleClass(classNode *ir.ClassStmt) {
	if d.meta.Classes.H == nil {
		d.meta.Classes = meta.NewClassesMap()
	}

	name := classNode.ClassName.Value
	namespace := d.ctx.st.Namespace
	fullName := namespace + `\` + name

	flags := d.handleClassModifiers(classNode.Modifiers)
	ifaces := d.handleClassImplements(classNode.Implements)
	parent := d.handleClassExtends(classNode.Extends)

	doc := d.parseClassPHPDoc(classNode, classNode.PhpDoc)
	docProps, docMethods, mixins := d.handleClassPhpDoc(doc)

	class := meta.ClassInfo{
		Pos:              d.getElementPos(classNode),
		Flags:            flags,
		Name:             fullName,
		Parent:           parent,
		ParentInterfaces: d.ctx.st.CurrentParentInterfaces,
		Interfaces:       ifaces,
		Traits:           make(map[string]struct{}),
		Methods:          docMethods,
		Properties:       docProps,
		Constants:        make(meta.ConstantsMap),
		Mixins:           mixins,
	}

	d.meta.Classes.Set(fullName, class)
}

func (d *rootWalker) handleClassModifiers(modifiers []*ir.Identifier) meta.ClassFlags {
	var kind meta.ClassFlags
	for _, m := range modifiers {
		switch {
		case strings.EqualFold("abstract", m.Value):
			kind |= meta.ClassAbstract
		case strings.EqualFold("final", m.Value):
			kind |= meta.ClassFinal
		}
	}
	return kind
}

func (d *rootWalker) handleClassImplements(impl *ir.ClassImplementsStmt) map[string]struct{} {
	if impl == nil {
		return map[string]struct{}{}
	}
	ifaces := make(map[string]struct{}, len(impl.InterfaceNames))

	for _, tr := range impl.InterfaceNames {
		interfaceName, ok := solver.GetClassName(d.ctx.st, tr)
		if !ok {
			continue
		}

		ifaces[interfaceName] = struct{}{}
	}

	return ifaces
}

func (d *rootWalker) handleClassExtends(exts *ir.ClassExtendsStmt) (parent string) {
	if exts == nil {
		return ""
	}

	className, ok := solver.GetClassName(d.ctx.st, exts.ClassName)
	if !ok {
		return ""
	}

	return className
}

func (d *rootWalker) handleClassPhpDoc(doc classPhpDocParseResult) (props meta.PropertiesMap, methods meta.FunctionsMap, mixins []string) {
	props = make(meta.PropertiesMap, len(doc.properties))
	methods = meta.NewFunctionsMap()

	// If we ever need to distinguish @property-annotated and real properties,
	// more work will be required here.
	for name, prop := range doc.properties {
		props[name] = prop
	}
	for name, method := range doc.methods.H {
		methods.Set(string(name), method)
	}

	mixins = doc.mixins

	return props, methods, mixins
}

func (d *rootWalker) parseClassPHPDoc(n ir.Node, doc []phpdoc.CommentPart) classPhpDocParseResult {
	var result classPhpDocParseResult

	if len(doc) == 0 {
		return result
	}

	// TODO: allocate maps lazily.
	// Class may not have any @property or @method annotations.
	// In that case we can handle avoid map allocations.
	result.properties = make(meta.PropertiesMap)
	result.methods = meta.NewFunctionsMap()

	for _, part := range doc {
		switch part.Name() {
		case "property", "property-read", "property-write":
			parseClassPHPDocProperty(&d.ctx, &result, part.(*phpdoc.TypeVarCommentPart))
		case "method":
			parseClassPHPDocMethod(&d.ctx, &result, part.(*phpdoc.RawCommentPart))
		case "mixin":
			parseClassPHPDocMixin(d.ctx.st, &result, part.(*phpdoc.RawCommentPart))
		}
	}

	return result
}

func (d *rootWalker) parseStartPos(pos *position.Position) (startLn []byte, startChar int) {
	if pos.StartLine >= 1 && d.file.NumLines() > pos.StartLine {
		startLn = d.file.Line(pos.StartLine - 1)
		p := d.file.LinePosition(pos.StartLine - 1)
		if pos.StartPos > p {
			startChar = pos.StartPos - p
		}
	}

	return startLn, startChar
}

func (d *rootWalker) report(n ir.Node, lineNumber int, level int, checkName, msg string, args ...interface{}) {
	if !meta.IsIndexingComplete() {
		return
	}
	if d.file.AutoGenerated() && !d.config.CheckAutoGenerated {
		return
	}
	// We don't report anything if linter was disabled by a
	// successful @linter disable, unless it's the linterError.
	if d.linterDisabled && checkName != "linterError" {
		return
	}

	isReportForNode := lineNumber == 0
	isReportForLine := !isReportForNode

	var pos position.Position
	var startLn []byte
	var startChar int
	var endLn []byte
	var endChar int

	if isReportForNode {
		if n == nil {
			// Hack to parse syntax error message from php-parser.
			// When in language server mode, do not map syntax errors in order not to
			// complain about unfinished piece of code that user is currently writing.
			if strings.Contains(msg, "syntax error") && strings.Contains(msg, " at line ") {
				// it is in form "Syntax error: syntax error: unexpected '*' at line 4"
				if lastIdx := strings.LastIndexByte(msg, ' '); lastIdx > 0 {
					lineNumStr := msg[lastIdx+1:]
					lineNum, err := strconv.Atoi(lineNumStr)
					if err == nil {
						pos.StartLine = lineNum
						pos.EndLine = lineNum
						msg = msg[0:lastIdx]
						msg = strings.TrimSuffix(msg, " at line")
					}
				}
			}
		} else {
			pos = *ir.GetPosition(n)
		}

		startLn, startChar = d.parseStartPos(&pos)

		if pos.EndLine >= 1 && d.file.NumLines() > pos.EndLine {
			endLn = d.file.Line(pos.EndLine - 1)
			p := d.file.LinePosition(pos.EndLine - 1)
			if pos.EndPos > p {
				endChar = pos.EndPos - p
			}
		} else {
			endLn = startLn
		}

		if endChar == 0 {
			endChar = len(endLn)
		}
	} else if isReportForLine {
		if lineNumber < 1 || lineNumber > d.file.NumLines() {
			return
		}

		startLn = d.file.Line(lineNumber - 1)
		startChar = 0
		endChar = len(startLn)

		if strings.HasSuffix(string(startLn), "\r") {
			endChar--
		}

		pos = position.Position{
			StartLine: lineNumber,
			EndLine:   lineNumber,
		}
	}

	// Replace Unused with Info (Notice) in non-LSP mode.
	if level == LevelUnused {
		level = LevelInformation
	}
	msg = fmt.Sprintf(msg, args...)
	var hash uint64
	if d.config.BaselineProfile != nil {
		// If baseline is not enabled, don't waste time on hash computations.
		hash = d.reportHash(&pos, startLn, checkName, msg)
		if count := d.ctx.baseline.Count(hash); count >= 1 {
			if d.ctx.hashCounters == nil {
				d.ctx.hashCounters = make(map[uint64]int)
			}
			d.ctx.hashCounters[hash]++
			if d.ctx.hashCounters[hash] <= count {
				return
			}
		}
	}

	d.reports = append(d.reports, &Report{
		CheckName: checkName,
		Context:   string(startLn),
		StartChar: startChar,
		EndChar:   endChar,
		Line:      pos.StartLine,
		Level:     level,
		Filename:  strings.ReplaceAll(d.ctx.st.CurrentFile, "\\", "/"), // To make output stable between platforms, see #572
		Message:   msg,
		Hash:      hash,
	})
}

// Report registers a single report message about some found problem.
func (d *rootWalker) Report(n ir.Node, level int, checkName, msg string, args ...interface{}) {
	d.report(n, 0, level, checkName, msg, args...)
}

// ReportByLine registers a single report message about some found problem in lineNumber code line.
func (d *rootWalker) ReportByLine(lineNumber int, level int, checkName, msg string, args ...interface{}) {
	d.report(nil, lineNumber, level, checkName, msg, args...)
}

// reportHash computes the report signature hash for the baseline.
func (d *rootWalker) reportHash(pos *position.Position, startLine []byte, checkName, msg string) uint64 {
	// Since we store class::method scope, renaming a class would cause baseline
	// invalidation for the entire class. But this is not an issue, since in
	// a modern PHP class name always should map to a filename.
	// If we renamed a class, we probably renamed the file as well, so
	// the baseline would be invalidated anyway.
	//
	// We still get all the benefits by using method prefix: it's common
	// for different classes to have methods with similar name. We don't
	// want such methods to be considered as a single "scope".
	scope := "file"
	switch {
	case d.ctx.st.CurrentClass != "" && d.ctx.st.CurrentFunction != "":
		scope = d.ctx.st.CurrentClass + "::" + d.ctx.st.CurrentFunction
	case d.ctx.st.CurrentFunction != "":
		scope = d.ctx.st.CurrentFunction
	}

	var prevLine []byte
	var nextLine []byte
	// Adjacent lines are only included when using non-conservative baseline.
	if !d.config.ConservativeBaseline {
		// Lines are 1-based, indexes are 0-based.
		// If this function is called, we expect that lines[index] exists.
		index := pos.StartLine - 1
		if index >= 1 {
			prevLine = d.file.Line(index - 1)
		}
		if d.file.NumLines() > index+1 {
			nextLine = d.file.Line(index + 1)
		}
	}

	// Observation: using a base file name instead of its full name does not
	// introduce any "bad collisions", because we have a "scope" part
	// that cuts the risk by a big margin.
	//
	// One benefit is that it makes the baseline contents more position-independent.
	// We don't need to know a source root folder to truncate it.
	//
	// Moves like Foo/Bar/User.php -> Common/User.php do not invalidate the
	// suppress base. This is not an obvious win, but it may be a good
	// compromise to solve a use case where a class file is being moved.
	// For classes, filename is unlikely to be changed during this file move operation.
	//
	// It can't handle file duplication when Foo/Bar/User.php is *copied* to a new location.
	// It may be useful to give warnings to both *old* and *new* copies of the duplicated
	// code since some bugs should be fixed in both places.
	// We'll see how it goes.
	filename := filepath.Base(d.ctx.st.CurrentFile)

	d.ctx.scratchBuf.Reset()
	return baseline.ReportHash(&d.ctx.scratchBuf, baseline.HashFields{
		Filename:  filename,
		PrevLine:  bytes.TrimSuffix(prevLine, []byte("\r")),
		StartLine: bytes.TrimSuffix(startLine, []byte("\r")),
		NextLine:  bytes.TrimSuffix(nextLine, []byte("\r")),
		CheckName: checkName,
		Message:   msg,
		Scope:     scope,
	})
}

func (d *rootWalker) reportUndefinedVariable(v ir.Node, maybeHave bool) {
	sv, ok := v.(*ir.SimpleVar)
	if !ok {
		d.Report(v, LevelInformation, "undefined", "Unknown variable variable %s used",
			meta.NameNodeToString(v))
		return
	}

	if _, ok := superGlobals[sv.Name]; ok {
		return
	}

	if maybeHave {
		d.Report(sv, LevelInformation, "undefined", "Variable might have not been defined: %s", sv.Name)
	} else {
		d.Report(sv, LevelError, "undefined", "Undefined variable: %s", sv.Name)
	}
}

func (d *rootWalker) handleComment(c freefloating.String) {
	if c.StringType != freefloating.CommentType {
		return
	}
	str := c.Value

	if !phpdoc.IsPHPDoc(str) {
		return
	}

	for _, ln := range phpdoc.Parse(d.ctx.phpdocTypeParser, str) {
		if ln.Name() != "linter" {
			continue
		}

		for _, p := range ln.(*phpdoc.RawCommentPart).Params {
			if p != "disable" {
				continue
			}
			if d.linterDisabled {
				needleLine := ln.Line() + c.Position.StartLine - 1
				d.ReportByLine(needleLine, LevelInformation, "linterError", "Linter is already disabled for this file")
				continue
			}
			canDisable := false
			if d.allowDisabledRegexp != nil {
				canDisable = d.allowDisabledRegexp.MatchString(d.ctx.st.CurrentFile)
			}
			d.linterDisabled = canDisable
			if !canDisable {
				needleLine := ln.Line() + c.Position.StartLine - 1
				d.ReportByLine(needleLine, LevelInformation, "linterError", "You are not allowed to disable linter")
			}
		}
	}
}

type handleFuncResult struct {
	returnTypes            meta.TypesMap
	prematureExitFlags     int
	callsParentConstructor bool
}

func (d *rootWalker) handleArrowFuncExpr(params []meta.FuncParam, expr ir.Node, sc *meta.Scope, parentBlockWalker *blockWalker) handleFuncResult {
	b := newBlockWalker(d, sc)
	b.inArrowFunction = true
	parentBlockWalker.parentBlockWalkers = append(parentBlockWalker.parentBlockWalkers, parentBlockWalker)
	b.parentBlockWalkers = parentBlockWalker.parentBlockWalkers

	for _, p := range params {
		if p.IsRef {
			b.nonLocalVars[p.Name] = varRef
		}
	}

	b.addStatement(expr)
	expr.Walk(b)

	b.flushUnused()

	return handleFuncResult{
		returnTypes: b.returnTypes,
	}
}

func (d *rootWalker) handleFuncStmts(params []meta.FuncParam, uses, stmts []ir.Node, sc *meta.Scope) handleFuncResult {
	b := newBlockWalker(d, sc)
	for _, createFn := range d.customBlock {
		b.custom = append(b.custom, createFn(&BlockContext{w: b}))
	}

	for _, useExpr := range uses {
		var byRef bool
		var v *ir.SimpleVar
		switch u := useExpr.(type) {
		case *ir.ReferenceExpr:
			v = u.Variable.(*ir.SimpleVar)
			byRef = true
		case *ir.SimpleVar:
			v = u
		}

		typ, ok := sc.GetVarNameType(v.Name)
		if !ok {
			typ = meta.NewTypesMap("TODO_use_var")
		}

		sc.AddVar(v, typ, "use", meta.VarAlwaysDefined)

		if !byRef {
			b.unusedVars[v.Name] = append(b.unusedVars[v.Name], v)
		} else {
			b.nonLocalVars[v.Name] = varRef
		}
	}

	// It's OK to read from and delete from a nil map.
	// If a func/method has 0 params, don't allocate a map for it.
	if len(params) != 0 {
		b.unusedParams = make(map[string]struct{}, len(params))
	}
	for _, p := range params {
		if p.IsRef {
			b.nonLocalVars[p.Name] = varRef
		}
		if !p.IsRef && !d.config.IsDiscardVar(p.Name) {
			b.unusedParams[p.Name] = struct{}{}
		}
	}
	for _, s := range stmts {
		b.addStatement(s)
		s.Walk(b)
	}
	b.flushUnused()

	// we can mark function as exiting abnormally if and only if
	// it only exits with die; or throw; and does not exit
	// using return; or any other control structure
	cleanFlags := b.ctx.exitFlags & (FlagDie | FlagThrow)

	var prematureExitFlags int
	if b.ctx.exitFlags == cleanFlags && (b.ctx.containsExitFlags&FlagReturn) == 0 {
		prematureExitFlags = cleanFlags
	}

	switch {
	case b.bareReturn && b.returnsValue:
		b.returnTypes = meta.MergeTypeMaps(b.returnTypes, meta.NullType)
	case b.returnTypes.IsEmpty() && b.returnsValue:
		b.returnTypes = meta.MixedType
	}

	return handleFuncResult{
		returnTypes:            b.returnTypes,
		prematureExitFlags:     prematureExitFlags,
		callsParentConstructor: b.callsParentConstructor,
	}
}

func (d *rootWalker) checkParentConstructorCall(n *ir.Identifier, parentConstructorCalled bool) {
	if !meta.IsIndexingComplete() {
		return
	}
	if n.Value != "__construct" {
		return
	}

	class, ok := d.currentClassNode.(*ir.ClassStmt)
	if !ok || class.Extends == nil {
		return
	}
	m, ok := solver.FindMethod(d.ctx.st.CurrentParentClass, `__construct`)
	if !ok || m.Info.AccessLevel == meta.Private || m.Info.IsAbstract() {
		return
	}

	if !parentConstructorCalled {
		d.Report(n, LevelWarning, "parentConstructor", "Missing parent::__construct() call")
	}
}

func (d *rootWalker) getElementPos(n ir.Node) meta.ElementPosition {
	pos := ir.GetPosition(n)
	_, startChar := d.parseStartPos(pos)

	return meta.ElementPosition{
		Filename:  d.ctx.st.CurrentFile,
		Character: int32(startChar),
		Line:      int32(pos.StartLine),
		EndLine:   int32(pos.EndLine),
		Length:    int32(pos.EndPos - pos.StartPos),
	}
}

type methodModifiers struct {
	abstract    bool
	static      bool
	final       bool
	accessLevel meta.AccessLevel
}

func (d *rootWalker) parseMethodModifiers(meth *ir.ClassMethodStmt) (res methodModifiers) {
	res.accessLevel = meta.Public

	for _, m := range meth.Modifiers {
		switch strings.ToLower(m.Value) {
		case "abstract":
			res.abstract = true
		case "static":
			res.static = true
		case "public":
			res.accessLevel = meta.Public
		case "private":
			res.accessLevel = meta.Private
		case "protected":
			res.accessLevel = meta.Protected
		case "final":
			res.final = true
		}
	}

	return res
}

func (d *rootWalker) getCurrentClass() (meta.ClassInfo, bool) {
	var classes meta.ClassesMap

	if d.ctx.st.IsTrait {
		classes = d.meta.Traits
	} else {
		classes = d.meta.Classes
	}

	cl, ok := classes.Get(d.ctx.st.CurrentClass)
	if !ok {
		if meta.IsIndexingComplete() {
			if d.ctx.st.IsTrait {
				cl, ok = meta.Info.GetTrait(d.ctx.st.CurrentClass)
			} else {
				cl, ok = meta.Info.GetClass(d.ctx.st.CurrentClass)
			}
			if !ok {
				return cl, false
			}
		}
	}

	return cl, true
}

func (d *rootWalker) checkLowerCaseModifier(m *ir.Identifier) string {
	lcase := strings.ToLower(m.Value)
	if lcase != m.Value {
		d.Report(m, LevelWarning, "keywordCase", "Use %s instead of %s",
			lcase, m.Value)
	}
	return lcase
}

func (d *rootWalker) handlePropertyList(pl *ir.PropertyListStmt) {
	class, ok := d.getCurrentClass()
	if !ok {
		return
	}

	isStatic, accessLevel := d.handlePropertyModifiers(pl)
	hintType := d.handleTypeHint(pl.Type)

	for _, pNode := range pl.Properties {
		prop := pNode.(*ir.PropertyStmt)
		propName := prop.Variable.Name

		typ := d.parsePHPDocVar(prop, prop.PhpDoc)
		if prop.Expr != nil {
			typ = typ.Append(solver.ExprTypeLocal(d.scope(), d.ctx.st, prop.Expr))
		}
		typ = typ.Append(hintType)

		if isStatic {
			propName = "$" + propName
		}

		// TODO: handle duplicate property
		class.Properties[propName] = meta.PropertyInfo{
			Pos:         d.getElementPos(prop),
			Typ:         typ.Immutable(),
			AccessLevel: accessLevel,
		}
	}
}

func (d *rootWalker) handlePropertyModifiers(pl *ir.PropertyListStmt) (bool, meta.AccessLevel) {
	isStatic := false
	accessLevel := meta.Public

	for _, m := range pl.Modifiers {
		switch strings.ToLower(m.Value) {
		case "public":
			accessLevel = meta.Public
		case "protected":
			accessLevel = meta.Protected
		case "private":
			accessLevel = meta.Private
		case "static":
			isStatic = true
		}
	}

	return isStatic, accessLevel
}

func (d *rootWalker) handleClassConstList(s *ir.ClassConstListStmt) {
	class, ok := d.getCurrentClass()
	if !ok {
		return
	}

	accessLevel := d.handleConstantAccessLevel(s)

	for _, cNode := range s.Consts {
		c := cNode.(*ir.ConstantStmt)

		constantName := c.ConstantName.Value
		typ := solver.ExprTypeLocal(d.scope(), d.ctx.st, c.Expr)
		value := constfold.Eval(d.ctx.st, c.Expr)

		// TODO: handle duplicate constant
		class.Constants[constantName] = meta.ConstInfo{
			Pos:         d.getElementPos(c),
			Typ:         typ.Immutable(),
			AccessLevel: accessLevel,
			Value:       value,
		}
	}
}

func (d *rootWalker) handleConstantAccessLevel(s *ir.ClassConstListStmt) meta.AccessLevel {
	level := meta.Public

	for _, m := range s.Modifiers {
		switch strings.ToLower(m.Value) {
		case "public":
			level = meta.Public
		case "protected":
			level = meta.Protected
		case "private":
			level = meta.Private
		}
	}

	return level
}

func (d *rootWalker) addClassMethodThisVariableToScope(modif methodModifiers, sc *meta.Scope) {
	if modif.static {
		return
	}

	sc.AddVarName("this", meta.NewTypesMap(d.ctx.st.CurrentClass).Immutable(), "instance method", meta.VarAlwaysDefined)
	sc.SetInInstanceMethod(true)
}

func (d *rootWalker) addClassMethodParamsToScope(method meta.FuncInfo, sc *meta.Scope) {
	for _, param := range method.Params {
		sc.AddVarName(param.Name, param.Typ, "param", meta.VarAlwaysDefined)
	}
}

func (d *rootWalker) handleClassMethod(m *ir.ClassMethodStmt) {
	class, ok := d.getCurrentClass()
	if !ok {
		return
	}

	name := m.MethodName.Value
	_, insideInterface := d.currentClassNode.(*ir.InterfaceStmt)

	sc := meta.NewScope()

	modif := d.parseMethodModifiers(m)
	hintReturnType := d.handleTypeHint(m.ReturnType)

	doc := d.parsePHPDoc(m.MethodName, m.PhpDoc, m.Params)
	phpdocReturnType := doc.returnType
	phpDocParamTypes := doc.types

	// need to proper return type
	d.addClassMethodThisVariableToScope(modif, sc)

	params, minParamsCnt := d.parseFuncArgs(m.Params, phpDocParamTypes, sc, nil)

	if len(class.Interfaces) != 0 {
		// If we implement interfaces, methods that take a part in this
		// can borrow types information from them.
		// Programmers sometimes leave implementing methods without a
		// comment or use @inheritdoc there.
		//
		// If method params are properly documented, it's possible to
		// derive that information, but we need to know in which
		// interface we can find that method.
		//
		// Since we don't have all interfaces during the indexing phase
		// and shouldn't update meta after it, we defer type resolving by
		// using BaseMethodParam here. We would have to lookup
		// matching interface during the type resolving.

		// Find params without type and annotate them with special
		// type that will force solver to walk interface types that
		// current class implements to have a chance of finding relevant type info.
		for i, p := range params {
			if !p.Typ.IsEmpty() {
				continue // Already has a type
			}

			if i > math.MaxUint8 {
				break // Current implementation limit reached
			}

			res := make(map[string]struct{})
			res[meta.WrapBaseMethodParam(i, d.ctx.st.CurrentClass, name)] = struct{}{}
			params[i].Typ = meta.NewTypesMapFromMap(res)
			// need to proper return type
			sc.AddVarName(p.Name, params[i].Typ, "param", meta.VarAlwaysDefined)
		}
	}

	stmts := convertNodeToStmts(m.Stmt)
	funcInfo := d.handleFuncStmts(params, nil, stmts, sc)
	actualReturnTypes := funcInfo.returnTypes
	exitFlags := funcInfo.prematureExitFlags

	returnTypes := functionReturnType(phpdocReturnType, hintReturnType, actualReturnTypes)
	funcFlags := d.transformClassMethodModifiersToFuncFlags(modif, insideInterface, stmts)

	// TODO: handle duplicate method
	class.Methods.Set(name, meta.FuncInfo{
		Params:       params,
		Name:         name,
		Pos:          d.getElementPos(m),
		Typ:          returnTypes.Immutable(),
		MinParamsCnt: minParamsCnt,
		AccessLevel:  modif.accessLevel,
		Flags:        funcFlags,
		ExitFlags:    exitFlags,
		Doc:          doc.info,
	})
}

func (d *rootWalker) transformClassMethodModifiersToFuncFlags(modif methodModifiers, insideInterface bool, stmts []ir.Node) meta.FuncFlags {
	var funcFlags meta.FuncFlags
	if modif.static {
		funcFlags |= meta.FuncStatic
	}
	if modif.abstract {
		funcFlags |= meta.FuncAbstract
	}
	if modif.final {
		funcFlags |= meta.FuncFinal
	}
	if !insideInterface && !modif.abstract && solver.SideEffectFreeFunc(d.scope(), d.ctx.st, nil, stmts) {
		funcFlags |= meta.FuncPure
	}
	return funcFlags
}

func (d *rootWalker) handleTypeHint(n ir.Node) meta.TypesMap {
	var hintReturnType meta.TypesMap
	if typ, ok := d.parseTypeNode(n); ok {
		hintReturnType = typ
	}
	return hintReturnType
}

func (d *rootWalker) reportPhpdocErrors(n ir.Node, errs phpdocErrors) {
	for _, err := range errs.phpdocLint {
		d.Report(n, LevelInformation, "phpdocLint", "%s", err)
	}
	for _, err := range errs.phpdocType {
		d.Report(n, LevelInformation, "phpdocType", "%s", err)
	}
}

func (d *rootWalker) parsePHPDocVar(n ir.Node, doc []phpdoc.CommentPart) (m meta.TypesMap) {
	for _, part := range doc {
		part, ok := part.(*phpdoc.TypeVarCommentPart)
		if ok && part.Name() == "var" {
			types, _ := typesFromPHPDoc(&d.ctx, part.Type)
			m = newTypesMap(&d.ctx, types)
		}
	}

	return m
}

func (d *rootWalker) checkPHPDocVar(n ir.Node, doc []phpdoc.CommentPart) {
	for _, part := range doc {
		d.checkPHPDocRef(n, part)
		part, ok := part.(*phpdoc.TypeVarCommentPart)
		if !ok || part.Name() != "var" {
			continue
		}

		_, warning := typesFromPHPDoc(&d.ctx, part.Type)
		if warning != "" {
			d.Report(n, LevelInformation, "phpdocType", "%s on line %d", warning, part.Line())
		}
	}
}

type phpDocParseResult struct {
	returnType meta.TypesMap
	types      phpDocParamsMap
	info       meta.PhpDocInfo
	errs       phpdocErrors
}

func (d *rootWalker) isValidPHPDocRef(n ir.Node, ref string) bool {
	// Skip:
	// - URLs
	// - Things that can be a filename (e.g. "foo.php")
	// - Wildcards (e.g. "self::FOO*")
	// - Issue references (e.g. "#1393" "BACK-103")
	if strings.Contains(ref, "http:") || strings.Contains(ref, "https:") {
		return true // OK: URL?
	}
	if strings.ContainsAny(ref, ".*-#") {
		return true
	}

	// expandName tries to convert s symbol into fully qualified form.
	expandName := func(s string) string {
		s, ok := solver.GetClassName(d.ctx.st, &ir.Name{Value: s})
		if !ok {
			return s
		}
		return s
	}

	isValidGlobalVar := func(ref string) bool {
		// Since we don't have an exhaustive list of globals,
		// we can't tell for sure whether a variable ref is correct.
		return true
	}

	isValidClassSymbol := func(ref string) bool {
		parts := strings.Split(ref, "::")
		if len(parts) != 2 {
			return false
		}
		typeName, symbolName := expandName(parts[0]), parts[1]
		if symbolName == "class" {
			_, ok := meta.Info.GetClass(typeName)
			return ok
		}
		if strings.HasPrefix(symbolName, "$") {
			return classHasProp(typeName, symbolName)
		}
		if _, ok := solver.FindMethod(typeName, symbolName); ok {
			return true
		}
		if _, _, ok := solver.FindConstant(typeName, symbolName); ok {
			return true
		}
		return false
	}

	isValidSymbol := func(ref string) bool {
		if !strings.HasPrefix(ref, `\`) {
			if d.currentClassNode != nil {
				if _, ok := solver.FindMethod(d.ctx.st.CurrentClass, ref); ok {
					return true // OK: class method reference
				}
				if classHasProp(d.ctx.st.CurrentClass, ref) {
					return true // OK: class prop reference
				}
			}

			// Functions and constants fall back in global namespace resolving.
			// See https://www.php.net/manual/en/language.namespaces.fallback.php
			globalRef := `\` + ref
			if _, ok := meta.Info.GetFunction(globalRef); ok {
				return true // OK: function reference
			}
			if _, ok := meta.Info.GetConstant(globalRef); ok {
				return true // OK: const reference
			}
		}
		fqnRef := expandName(ref)
		if _, ok := meta.Info.GetFunction(fqnRef); ok {
			return true // OK: FQN function reference
		}
		if _, ok := meta.Info.GetClass(fqnRef); ok {
			return true // OK: FQN class reference
		}
		if _, ok := meta.Info.GetConstant(fqnRef); ok {
			return true // OK: FQN const reference
		}
		return false
	}

	switch {
	case strings.Contains(ref, "::"):
		return isValidClassSymbol(ref)
	case strings.HasPrefix(ref, "$"):
		return isValidGlobalVar(ref)
	default:
		return isValidSymbol(ref)
	}
}

func (d *rootWalker) checkPHPDocRef(n ir.Node, part phpdoc.CommentPart) {
	if !meta.IsIndexingComplete() {
		return
	}

	switch part.Name() {
	case "mixin":
		d.checkPHPDocMixinRef(n, part)
	case "see":
		d.checkPHPDocSeeRef(n, part)
	}
}

func (d *rootWalker) checkPHPDocSeeRef(n ir.Node, part phpdoc.CommentPart) {
	params := part.(*phpdoc.RawCommentPart).Params
	if len(params) == 0 {
		return
	}

	// @see supports a comma-separated list of refs.
	var refs []string
	for _, p := range params {
		refs = append(refs, strings.TrimSuffix(p, ","))
		if !strings.HasSuffix(p, ",") {
			break
		}
	}

	for _, ref := range refs {
		// Sometimes people write references like `foo()` `foo...` `foo@`.
		ref = strings.TrimRight(ref, "().;@")
		if !d.isValidPHPDocRef(n, ref) {
			d.Report(n, LevelWarning, "phpdocRef", "line %d: @see tag refers to unknown symbol %s",
				part.Line(), ref)
		}
	}
}

func (d *rootWalker) checkPHPDocMixinRef(n ir.Node, part phpdoc.CommentPart) {
	rawPart, ok := part.(*phpdoc.RawCommentPart)
	if !ok {
		return
	}

	params := rawPart.Params
	if len(params) == 0 {
		return
	}

	name, ok := solver.GetClassName(d.ctx.st, &ir.Name{
		Value: params[0],
	})

	if !ok {
		return
	}

	if _, ok := meta.Info.GetClass(name); !ok {
		d.Report(n, LevelWarning, "phpdocRef", "line %d: @mixin tag refers to unknown class %s", part.Line(), name)
	}
}

func (d *rootWalker) parsePHPDoc(n ir.Node, doc []phpdoc.CommentPart, actualParams []ir.Node) phpDocParseResult {
	var result phpDocParseResult

	if len(doc) == 0 {
		return result
	}

	actualParamNames := make(map[string]struct{}, len(actualParams))
	for _, p := range actualParams {
		p := p.(*ir.Parameter)
		actualParamNames[p.Variable.Name] = struct{}{}
	}

	result.types = make(phpDocParamsMap, len(actualParams))

	var curParam int

	for _, part := range doc {
		d.checkPHPDocRef(n, part)

		if part.Name() == "deprecated" {
			part := part.(*phpdoc.RawCommentPart)
			result.info.Deprecated = true
			result.info.DeprecationNote = part.ParamsText
			continue
		}

		if part.Name() == "return" {
			part := part.(*phpdoc.TypeCommentPart)
			types, warning := typesFromPHPDoc(&d.ctx, part.Type)
			if warning != "" {
				result.errs.pushType("%s on line %d", warning, part.Line())
			}
			result.returnType = newTypesMap(&d.ctx, types)
			continue
		}

		// Rest is for @param handling.

		if part.Name() != "param" {
			continue
		}

		part := part.(*phpdoc.TypeVarCommentPart)
		optional := strings.Contains(part.Rest, "[optional]")
		switch {
		case part.Var == "":
			result.errs.pushLint("malformed @param tag (maybe var is missing?) on line %d", part.Line())
		case part.Type.IsEmpty():
			result.errs.pushLint("malformed @param %s tag (maybe type is missing?) on line %d",
				part.Var, part.Line())
			continue
		}

		if part.VarIsFirst {
			// Phpstorm gives the same message.
			result.errs.pushLint("non-canonical order of variable and type on line %d", part.Line())
		}

		variable := part.Var
		if !strings.HasPrefix(variable, "$") {
			if len(actualParams) > curParam {
				variable = actualParams[curParam].(*ir.Parameter).Variable.Name
			}
		}
		if _, ok := actualParamNames[strings.TrimPrefix(variable, "$")]; !ok {
			result.errs.pushLint("@param for non-existing argument %s", variable)
			continue
		}

		curParam++

		var param phpDocParamEl
		types, warning := typesFromPHPDoc(&d.ctx, part.Type)
		if warning != "" {
			result.errs.pushType("%s on line %d", warning, part.Line())
		}
		param.typ = newTypesMap(&d.ctx, types)
		param.typ.Iterate(func(t string) {
			if t == "void" {
				result.errs.pushType("void is not a valid type for input parameter")
			}
		})
		param.optional = optional

		variable = strings.TrimPrefix(variable, "$")
		result.types[variable] = param
	}

	result.returnType = result.returnType.Immutable()
	return result
}

// parse type info, e.g. "string" in "someFunc() : string { ... }"
func (d *rootWalker) parseTypeNode(n ir.Node) (typ meta.TypesMap, ok bool) {
	if n == nil {
		return meta.TypesMap{}, false
	}

	types := typesFromNode(n)
	tm := newTypesMap(&d.ctx, types)
	return tm, !tm.IsEmpty()
}

// callbackParamByIndex returns the description of the parameter for the function by its index.
func (d *rootWalker) callbackParamByIndex(param ir.Node, argType meta.TypesMap) meta.FuncParam {
	p := param.(*ir.Parameter)
	v := p.Variable

	var typ meta.TypesMap
	tp, ok := d.parseTypeNode(p.VariableType)
	if ok {
		typ = tp
	} else {
		typ = argType.Map(meta.WrapElemOf)
	}

	arg := meta.FuncParam{
		IsRef: p.ByRef,
		Name:  v.Name,
		Typ:   typ,
	}
	return arg
}

func (d *rootWalker) parseFuncArgsForCallback(params []ir.Node, sc *meta.Scope, closureSolver *solver.ClosureCallerInfo) (args []meta.FuncParam, minArgs int) {
	countParams := len(params)
	minArgs = countParams
	if countParams == 0 {
		return nil, 0
	}
	args = make([]meta.FuncParam, countParams)

	switch closureSolver.Name {
	case `\usort`, `\uasort`, `\array_reduce`:
		args[0] = d.callbackParamByIndex(params[0], closureSolver.ArgTypes[0])
		if countParams > 1 {
			args[1] = d.callbackParamByIndex(params[1], closureSolver.ArgTypes[0])
		}
	case `\array_walk`, `\array_walk_recursive`, `\array_filter`:
		args[0] = d.callbackParamByIndex(params[0], closureSolver.ArgTypes[0])
	case `\array_map`:
		args[0] = d.callbackParamByIndex(params[0], closureSolver.ArgTypes[1])
	}

	for i, param := range params {
		p := param.(*ir.Parameter)
		v := p.Variable
		var typ meta.TypesMap
		if i < len(args) {
			typ = args[i].Typ
		} else {
			typ = meta.MixedType
		}

		sc.AddVarName(v.Name, typ, "param", meta.VarAlwaysDefined)
	}

	return args, minArgs
}

func (d *rootWalker) parseFuncArgs(params []ir.Node, parTypes phpDocParamsMap, sc *meta.Scope, closureSolver *solver.ClosureCallerInfo) (args []meta.FuncParam, minArgs int) {
	if len(params) == 0 {
		return nil, 0
	}

	args = make([]meta.FuncParam, 0, len(params))

	if closureSolver != nil && solver.IsSupportedFunction(closureSolver.Name) {
		return d.parseFuncArgsForCallback(params, sc, closureSolver)
	}

	for _, param := range params {
		p := param.(*ir.Parameter)
		v := p.Variable
		parTyp := parTypes[v.Name]

		if !parTyp.typ.IsEmpty() {
			sc.AddVarName(v.Name, parTyp.typ, "param", meta.VarAlwaysDefined)
		}

		typ := parTyp.typ

		if p.DefaultValue == nil && !parTyp.optional && !p.Variadic {
			minArgs++
		}

		if p.VariableType != nil {
			if varTyp, ok := d.parseTypeNode(p.VariableType); ok {
				typ = varTyp
			}
		} else if typ.IsEmpty() && p.DefaultValue != nil {
			typ = solver.ExprTypeLocal(sc, d.ctx.st, p.DefaultValue)
			// For the type resolver default value can look like a
			// precise source of information (e.g. "false" is a precise bool),
			// but it's not assigned unconditionally.
			// If explicit argument is provided, that parameter can have
			// almost any type possible.
			typ.MarkAsImprecise()
		}

		if p.Variadic {
			typ = typ.Map(meta.WrapArrayOf)
		}

		sc.AddVarName(v.Name, typ, "param", meta.VarAlwaysDefined)

		par := meta.FuncParam{
			Typ:   typ.Immutable(),
			IsRef: p.ByRef,
		}

		par.Name = v.Name
		args = append(args, par)
	}
	return args, minArgs
}

func (d *rootWalker) checkCommentMisspellings(n ir.Node, s string) {
	// Try to avoid checking for symbol names and references.
	d.checkMisspellings(n, s, "misspellComment", isCapitalized)
}

func (d *rootWalker) checkVarnameMisspellings(n ir.Node, s string) {
	d.checkMisspellings(n, s, "misspellName", func(string) bool {
		return false
	})
}

func (d *rootWalker) checkIdentMisspellings(n *ir.Identifier) {
	d.checkMisspellings(n, n.Value, "misspellName", func(s string) bool {
		// Before PHP got context-sensitive lexer, it was common to use
		// method names to avoid parsing errors.
		// We can't suggest a fix that leads to a parsing error.
		// To avoid false positives, skip PHP keywords.
		return phpKeywords[s]
	})
}

func (d *rootWalker) checkMisspellings(n ir.Node, s string, label string, skip func(string) bool) {
	if !meta.IsIndexingComplete() {
		return
	}
	if d.config.TypoFixer == nil {
		return
	}
	_, changes := d.config.TypoFixer.Replace(s)
	for _, c := range changes {
		if skip(c.Corrected) || skip(c.Original) {
			continue
		}
		d.Report(n, LevelDoNotReject, label, `"%s" is a misspelling of "%s"`, c.Original, c.Corrected)
	}
}

func (d *rootWalker) enterFunction(fun *ir.FunctionStmt) bool {
	nm := d.ctx.st.Namespace + `\` + fun.FunctionName.Value
	pos := ir.GetPosition(fun)

	if funcSize := pos.EndLine - pos.StartLine; funcSize > maxFunctionLines {
		d.Report(fun.FunctionName, LevelDoNotReject, "complexity", "Too big function: more than %d lines", maxFunctionLines)
	}

	var hintReturnType meta.TypesMap
	if typ, ok := d.parseTypeNode(fun.ReturnType); ok {
		hintReturnType = typ
	}

	d.checkCommentMisspellings(fun.FunctionName, fun.PhpDocComment)
	d.checkIdentMisspellings(fun.FunctionName)
	for _, param := range fun.Params {
		d.checkVarnameMisspellings(param, param.(*ir.Parameter).Variable.Name)
		d.checkFuncParam(param.(*ir.Parameter))
	}

	doc := d.parsePHPDoc(fun.FunctionName, fun.PhpDoc, fun.Params)
	d.reportPhpdocErrors(fun.FunctionName, doc.errs)
	phpdocReturnType := doc.returnType
	phpDocParamTypes := doc.types

	if d.meta.Functions.H == nil {
		d.meta.Functions = meta.NewFunctionsMap()
	}

	sc := meta.NewScope()

	params, minParamsCnt := d.parseFuncArgs(fun.Params, phpDocParamTypes, sc, nil)

	funcInfo := d.handleFuncStmts(params, nil, fun.Stmts, sc)
	actualReturnTypes := funcInfo.returnTypes
	exitFlags := funcInfo.prematureExitFlags

	returnTypes := functionReturnType(phpdocReturnType, hintReturnType, actualReturnTypes)

	var funcFlags meta.FuncFlags
	if solver.SideEffectFreeFunc(d.scope(), d.ctx.st, nil, fun.Stmts) {
		funcFlags |= meta.FuncPure
	}
	d.meta.Functions.Set(nm, meta.FuncInfo{
		Params:       params,
		Name:         nm,
		Pos:          d.getElementPos(fun),
		Typ:          returnTypes.Immutable(),
		MinParamsCnt: minParamsCnt,
		Flags:        funcFlags,
		ExitFlags:    exitFlags,
		Doc:          doc.info,
	})

	return false
}

func (d *rootWalker) checkFuncParam(p *ir.Parameter) {
	// TODO(quasilyte): DefaultValue can only contain constant expressions.
	// Could run special check over them to detect the potential fatal errors.
	walkNode(p.DefaultValue, func(w ir.Node) bool {
		if n, ok := w.(*ir.ArrayExpr); ok && !n.ShortSyntax {
			d.Report(n, LevelDoNotReject, "arraySyntax", "Use of old array syntax (use short form instead)")
		}
		return true
	})
}

func (d *rootWalker) enterFunctionCall(s *ir.FunctionCallExpr) bool {
	nm, ok := s.Function.(*ir.Name)
	if !ok {
		return true
	}

	if d.ctx.st.Namespace == `\PHPSTORM_META` && nm.Value == `override` {
		return d.handleOverride(s)
	}

	if nm.Value != `define` || len(s.Args) < 2 {
		// TODO: actually we could warn about bogus defines
		return true
	}

	arg := s.Arg(0)

	str, ok := arg.Expr.(*ir.String)
	if !ok {
		return true
	}

	valueArg := s.Arg(1)

	if d.meta.Constants == nil {
		d.meta.Constants = make(meta.ConstantsMap)
	}

	value := constfold.Eval(d.ctx.st, valueArg)

	d.meta.Constants[`\`+strings.TrimFunc(str.Value, isQuote)] = meta.ConstInfo{
		Pos:   d.getElementPos(s),
		Typ:   solver.ExprTypeLocal(d.scope(), d.ctx.st, valueArg.Expr),
		Value: value,
	}
	return true
}

// Handle e.g. "override(\array_shift(0), elementType(0));"
// which means "return type of array_shift() is the type of element of first function parameter"
func (d *rootWalker) handleOverride(s *ir.FunctionCallExpr) bool {
	if len(s.Args) != 2 {
		return true
	}

	arg0 := s.Arg(0)
	arg1 := s.Arg(1)

	fc0, ok := arg0.Expr.(*ir.FunctionCallExpr)
	if !ok {
		return true
	}

	fc1, ok := arg1.Expr.(*ir.FunctionCallExpr)
	if !ok {
		return true
	}

	fnNameNode, ok := fc0.Function.(*ir.Name)
	if !ok || !fnNameNode.IsFullyQualified() {
		return true
	}

	overrideNameNode, ok := fc1.Function.(*ir.Name)
	if !ok {
		return true
	}

	if len(fc1.Args) != 1 {
		return true
	}

	fc1Arg0 := fc1.Arg(0)

	argNumNode, ok := fc1Arg0.Expr.(*ir.Lnumber)
	if !ok {
		return true
	}

	argNum, err := strconv.Atoi(argNumNode.Value)
	if err != nil {
		return true
	}

	var overrideTyp meta.OverrideType
	switch {
	case overrideNameNode.Value == `type`:
		overrideTyp = meta.OverrideArgType
	case overrideNameNode.Value == `elementType`:
		overrideTyp = meta.OverrideElementType
	default:
		return true
	}

	fnName := fnNameNode.Value

	if d.meta.FunctionOverrides == nil {
		d.meta.FunctionOverrides = make(meta.FunctionsOverrideMap)
	}

	d.meta.FunctionOverrides[fnName] = meta.FuncInfoOverride{
		OverrideType: overrideTyp,
		ArgNum:       argNum,
	}

	return true
}

func (d *rootWalker) enterConstList(lst *ir.ConstListStmt) bool {
	if d.meta.Constants == nil {
		d.meta.Constants = make(meta.ConstantsMap)
	}

	for _, sNode := range lst.Consts {
		s := sNode.(*ir.ConstantStmt)

		value := constfold.Eval(d.ctx.st, s.Expr)

		id := s.ConstantName
		nm := d.ctx.st.Namespace + `\` + id.Value

		d.meta.Constants[nm] = meta.ConstInfo{
			Pos:   d.getElementPos(s),
			Typ:   solver.ExprTypeLocal(d.scope(), d.ctx.st, s.Expr),
			Value: value,
		}
	}

	return false
}

// LeaveNode is invoked after node process
func (d *rootWalker) LeaveNode(n ir.Node) {
	for _, c := range d.custom {
		c.BeforeLeaveNode(n)
	}

	switch n.(type) {
	case *ir.ClassStmt, *ir.InterfaceStmt, *ir.TraitStmt:
		d.currentClassNode = nil
	}

	state.LeaveNode(d.ctx.st, n)

	for _, c := range d.custom {
		c.AfterLeaveNode(n)
	}
}

func (d *rootWalker) runRules(n ir.Node, sc *meta.Scope, rlist []rules.Rule) {
	for i := range rlist {
		rule := &rlist[i]
		if d.runRule(n, sc, rule) {
			// Stop at the first matched rule per IR node.
			// Sometimes it's useful to report more, but we rely on the rules definition
			// order so we can report more specific issues instead of the
			// more generic ones whether possible.
			// This also makes rules execution faster.
			break
		}
	}
}

// nodeText is used to get the string that represents the
// passed node more efficiently than irutil.FmtNode.
func (d *rootWalker) nodeText(n ir.Node) string {
	pos := ir.GetPosition(n)
	from, to := pos.StartPos, pos.EndPos
	src := d.file.Contents()
	// Taking a node from the source code preserves the original formatting
	// and is more efficient than printing it.
	if (from >= 0 && from < len(src)) && (to >= 0 && to < len(src)) {
		return string(src[from:to])
	}
	// If we can't take node out of the source text, print it.
	return irutil.FmtNode(n)
}

func (d *rootWalker) renderRuleMessage(msg string, n ir.Node, m phpgrep.MatchData, truncate bool) string {
	// "$$" stands for the entire matched node, like $0 in regexp.
	if strings.Contains(msg, "$$") {
		msg = strings.ReplaceAll(msg, "$$", d.nodeText(n))
	}

	if len(m.Capture) == 0 {
		return msg // No variables to interpolate, we're done
	}
	for _, c := range m.Capture {
		key := "$" + c.Name
		if !strings.Contains(msg, key) {
			continue
		}
		nodeString := d.nodeText(c.Node)
		// Don't interpolate strings that are too long
		// or contain a newline.
		var replacement string
		if truncate && (len(nodeString) > 60 || strings.Contains(nodeString, "\n")) {
			replacement = key
		} else {
			replacement = nodeString
		}
		msg = strings.ReplaceAll(msg, key, replacement)
	}
	return msg
}

func (d *rootWalker) runRule(n ir.Node, sc *meta.Scope, rule *rules.Rule) bool {
	m, ok := rule.Matcher.Match(n)
	if !ok {
		return false
	}

	matched := false
	if len(rule.Filters) == 0 {
		matched = true
	} else {
		for _, filterSet := range rule.Filters {
			if d.checkFilterSet(&m, sc, filterSet) {
				matched = true
				break
			}
		}
	}

	// If location is explicitly set, use named match set.
	// Otherwise peek the root target node.
	var location ir.Node
	switch {
	case matched && rule.Location != "":
		named, _ := m.CapturedByName(rule.Location)
		location = named
	case matched:
		location = n
	}

	if location == nil {
		return false
	}

	message := d.renderRuleMessage(rule.Message, n, m, true)
	d.Report(location, rule.Level, rule.Name, "%s", message)

	if d.config.ApplyQuickFixes && rule.Fix != "" {
		// As rule sets contain only enabled rules,
		// we should be OK without any filtering here.
		pos := ir.GetPosition(n)
		d.ctx.fixes = append(d.ctx.fixes, quickfix.TextEdit{
			StartPos:    pos.StartPos,
			EndPos:      pos.EndPos,
			Replacement: d.renderRuleMessage(rule.Fix, n, m, false),
		})
	}

	return true
}

func (d *rootWalker) checkTypeFilter(wantType *phpdoc.Type, sc *meta.Scope, nn ir.Node) bool {
	if wantType == nil {
		return true
	}

	// TODO: compare without converting a TypesMap into TypeExpr?
	// Or maybe store TypeExpr inside a TypesMap instead of strings?
	// Can we use `meta.Type` for this?
	typ := solver.ExprType(sc, d.ctx.st, nn)
	haveType := typesMapToTypeExpr(d.ctx.phpdocTypeParser, typ)
	return rules.TypeIsCompatible(wantType.Expr, haveType.Expr)
}

func (d *rootWalker) checkFilterSet(m *phpgrep.MatchData, sc *meta.Scope, filterSet map[string]rules.Filter) bool {
	// TODO: pass custom types here, so both @type and @pure predicates can use it.

	for name, filter := range filterSet {
		nn, ok := m.CapturedByName(name)
		if !ok {
			continue
		}

		if !d.checkTypeFilter(filter.Type, sc, nn) {
			return false
		}
		if filter.Pure && !solver.SideEffectFree(d.scope(), d.ctx.st, nil, nn) {
			return false
		}
	}

	return true
}

func (d *rootWalker) checkTraitImplemented(n ir.Node, nameUsed string) {
	trait, ok := meta.Info.GetTrait(nameUsed)
	if !ok {
		d.reportUndefinedType(n, nameUsed)
		return
	}
	d.checkImplemented(n, nameUsed, trait)
}

func (d *rootWalker) checkClassImplemented(n ir.Node, nameUsed string) {
	class, ok := meta.Info.GetClass(nameUsed)
	if !ok {
		d.reportUndefinedType(n, nameUsed)
		return
	}
	d.checkImplemented(n, nameUsed, class)
}

func (d *rootWalker) checkIfaceImplemented(n ir.Node, nameUsed string) {
	d.checkClassImplemented(n, nameUsed)
}

func (d *rootWalker) checkImplemented(n ir.Node, nameUsed string, otherClass meta.ClassInfo) {
	class, ok := d.getCurrentClass()
	if !ok {
		return
	}

	if d.ctx.st.IsTrait || class.IsAbstract() {
		return
	}
	d.checkNameCase(n, nameUsed, otherClass.Name)
	visited := make(map[string]struct{}, 4)
	d.checkImplementedStep(n, nameUsed, otherClass, visited)
}

func (d *rootWalker) checkImplementedStep(n ir.Node, className string, otherClass meta.ClassInfo, visited map[string]struct{}) {
	// TODO: check that method signatures are compatible?
	if _, ok := visited[className]; ok {
		return
	}
	visited[className] = struct{}{}
	for _, ifaceMethod := range otherClass.Methods.H {
		m, ok := solver.FindMethod(d.ctx.st.CurrentClass, ifaceMethod.Name)
		if !ok || !m.Implemented {
			d.Report(n, LevelError, "unimplemented", "Class %s must implement %s::%s method",
				d.ctx.st.CurrentClass, className, ifaceMethod.Name)
			continue
		}
		if m.Info.Name != ifaceMethod.Name {
			d.Report(n, LevelDoNotReject, "nameCase", "%s::%s should be spelled as %s::%s",
				d.ctx.st.CurrentClass, m.Info.Name, className, ifaceMethod.Name)
		}
	}
	for _, ifaceName := range otherClass.ParentInterfaces {
		iface, ok := meta.Info.GetClass(ifaceName)
		if ok {
			d.checkImplementedStep(n, ifaceName, iface, visited)
		}
	}
	if otherClass.Parent != "" {
		class, ok := meta.Info.GetClass(otherClass.Parent)
		if ok {
			d.checkImplementedStep(n, otherClass.Parent, class, visited)
		}
	}
}

func (d *rootWalker) reportUndefinedType(n ir.Node, name string) {
	d.Report(n, LevelError, "undefined", "Type %s not found", name)
}

func (d *rootWalker) checkNameCase(n ir.Node, nameUsed, nameExpected string) {
	if nameUsed == "" || nameExpected == "" {
		return
	}
	if nameUsed != nameExpected {
		d.Report(n, LevelInformation, "nameCase", "%s should be spelled %s",
			nameUsed, nameExpected)
	}
}

func (d *rootWalker) checkKeywordCasePos(n ir.Node, begin int, keyword string) {
	from := begin
	to := from + len(keyword)

	wantKwd := keyword
	haveKwd := d.file.Contents()[from:to]
	if wantKwd != string(haveKwd) {
		d.Report(n, LevelWarning, "keywordCase", "Use %s instead of %s",
			wantKwd, haveKwd)
	}
}

func (d *rootWalker) checkKeywordCase(n ir.Node, keyword string) {
	// Only works for nodes that have a keyword of interest
	// as the leftmost token.
	d.checkKeywordCasePos(n, ir.GetPosition(n).StartPos, keyword)
}

func (d *rootWalker) beforeEnterFile() {
	for _, c := range d.custom {
		c.BeforeEnterFile()
	}
}

func (d *rootWalker) afterLeaveFile() {
	for _, c := range d.custom {
		c.AfterLeaveFile()
	}

	if !meta.IsIndexingComplete() {
		for _, shape := range d.ctx.shapes {
			props := make(meta.PropertiesMap)
			for _, p := range shape.props {
				props[p.key] = meta.PropertyInfo{
					Typ:         newTypesMap(&d.ctx, p.types).Immutable(),
					AccessLevel: meta.Public,
				}
			}
			cl := meta.ClassInfo{
				Name:       shape.name,
				Properties: props,
				Flags:      meta.ClassShape,
			}
			if d.meta.Classes.H == nil {
				d.meta.Classes = meta.NewClassesMap()
			}
			d.meta.Classes.Set(shape.name, cl)
		}
	}
}
