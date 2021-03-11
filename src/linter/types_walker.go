package linter

import (
	"fmt"
	"strings"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/meta"
	"github.com/VKCOM/noverify/src/solver"
)

type Type struct {
	typ     map[string]struct{}
	precise bool
}

func NewType(typesMap meta.TypesMap) Type {
	return Type{
		typ:     typesMap.Raw(),
		precise: typesMap.IsPrecise(),
	}
}

type TypesWalker struct {
	Types map[ir.Node]Type

	state  *meta.ClassParseState
	scope  *meta.Scope
	custom []solver.CustomType // TODO: initialize
}

func NewTypesWalker(state *meta.ClassParseState) *TypesWalker {
	return &TypesWalker{Types: map[ir.Node]Type{}, state: state, scope: meta.NewScope()}
}

func (t *TypesWalker) TypeOf(n ir.Node) Type {
	if n == nil {
		return Type{typ: map[string]struct{}{"mixed": {}}}
	}

	return t.Types[n]
}

func (t *TypesWalker) EnterNode(n ir.Node) bool {
	switch n := n.(type) {
	case *ir.Lnumber:
		t.Types[n] = NewType(meta.PreciseIntType)
	case *ir.SimpleVar:
		return t.handleVariable(n)
	case *ir.StaticStmt:
		return t.handleStaticStmt(n)
	case *ir.ArrayExpr:
		return t.handleArrayExpr(n)
	case *ir.ForeachStmt:
		return t.handleForeach(n)
	case *ir.FunctionCallExpr:
		typ := solver.ExprType(t.scope, t.state, n)
		t.Types[n] = NewType(typ)
	case *ir.Assign:
		return t.handleAssign(n)
	}

	return true
}

func (t *TypesWalker) handleForeach(s *ir.ForeachStmt) bool {
	// expression is always executed and is executed in base context
	if s.Expr != nil {
		s.Expr.Walk(t)
	}

	// foreach body can do 0 cycles so we need a separate context for that
	if s.Stmt != nil {
		scope := t.withNewContext(func() {
			exprType := solver.ExprTypeCustom(t.scope, t.state, s.Expr, t.custom)
			t.Types[s.Expr] = NewType(exprType)

			t.handleVariableNode(s.Variable, exprType.Map(meta.WrapElemOf))
			t.handleVariableNode(s.Key, arrayKeyType)

			if list, ok := s.Variable.(*ir.ListExpr); ok {
				for _, item := range list.Items {
					t.handleVariableNode(item.Val, meta.MixedType)
				}
			}

			s.Stmt.Walk(t)
		})

		t.maybeAddAllVars(scope)
	}

	return false
}

func (t *TypesWalker) handleArrayExpr(n *ir.ArrayExpr) bool {
	typ := solver.ExprType(t.scope, t.state, n)
	t.Types[n] = NewType(typ)

	for _, item := range n.Items {
		if item.Val != nil {
			item.Val.Walk(t)
		}
		if item.Key != nil {
			item.Key.Walk(t)
		}
	}

	return false
}

func (t *TypesWalker) handleStaticStmt(n *ir.StaticStmt) bool {
	for _, vv := range n.Vars {
		v := vv.(*ir.StaticVarStmt)
		ev := v.Variable
		typ := solver.ExprTypeLocalCustom(t.scope, t.state, v.Expr, t.custom)
		// Static vars can be assigned below and preserve the type of
		// the previously assigned value.
		typ.MarkAsImprecise()
		t.addVarName(ev.Name, typ)
		// b.addNonLocalVarName(ev.Name, varStatic)
		t.Types[vv] = NewType(typ)
		t.Types[ev] = NewType(typ)

		if v.Expr != nil {
			v.Expr.Walk(t)
		}
	}

	return false
}

func (t *TypesWalker) handleVar(n *ir.SimpleVar) {
	typ, ok := t.scope.GetVarNameType(n.Name)
	if ok {
		t.Types[n] = NewType(typ)
	}
}

func (t *TypesWalker) handleAssign(n *ir.Assign) bool {
	n.Expr.Walk(t)

	typ := solver.ExprType(t.scope, t.state, n.Expr)

	tp := NewType(typ)
	t.Types[n.Expr] = tp
	t.Types[n.Variable] = tp

	switch v := n.Variable.(type) {
	case *ir.ArrayDimFetchExpr:
		t.handleDimFetchLValue(v, typ)
	case *ir.SimpleVar, *ir.Var:
		t.replaceVar(v, typ)
	case *ir.ListExpr:
		t.handleAssignList(v, typ)
	// case *ir.PropertyFetchExpr:
	// 	v.Property.Walk(t)
	// 	sv, ok := v.Variable.(*ir.SimpleVar)
	// 	if !ok {
	// 		v.Variable.Walk(t)
	// 		break
	// 	}
	//
	// 	if sv.Name != "this" {
	// 		break
	// 	}
	//
	// 	if t.state.CurrentClass == "" {
	// 		break
	// 	}
	//
	// 	propertyName, ok := v.Property.(*ir.Identifier)
	// 	if !ok {
	// 		break
	// 	}
	//
	// 	cls := b.r.getClass()
	//
	// 	p := cls.Properties[propertyName.Value]
	// 	p.Typ = p.Typ.Append(solver.ExprTypeLocalCustom(b.ctx.sc, b.r.ctx.st, a.Expr, b.ctx.customTypes))
	// 	cls.Properties[propertyName.Value] = p
	//
	// case *ir.StaticPropertyFetchExpr:
	// 	sv, ok := v.Property.(*ir.SimpleVar)
	// 	if !ok {
	// 		vv := v.Property.(*ir.Var)
	// 		vv.Expr.Walk(b)
	// 		break
	// 	}
	//
	// 	if b.r.ctx.st.CurrentClass == "" {
	// 		break
	// 	}
	//
	// 	className, ok := solver.GetClassName(b.r.ctx.st, v.Class)
	// 	if !ok || className != b.r.ctx.st.CurrentClass {
	// 		break
	// 	}
	//
	// 	cls := b.r.getClass()
	//
	// 	p := cls.Properties["$"+sv.Name]
	// 	p.Typ = p.Typ.Append(solver.ExprTypeLocalCustom(b.ctx.sc, b.r.ctx.st, a.Expr, b.ctx.customTypes))
	// 	cls.Properties["$"+sv.Name] = p
	default:
		n.Variable.Walk(t)
	}

	return false
}

func (t *TypesWalker) handleAssignList(list *ir.ListExpr, rhsType meta.TypesMap) {
	typ := rhsType

	// Hint: only const (literal) size hints work for this.
	// Hint: check the compiler output to see whether elemTypes "escape" or not.
	//
	// We store meta.Type instead of string to avoid the need to do strings.Join
	// when we want to create a TypesMap.
	var elemTypes []meta.Type
	var shapeType string
	typ.Iterate(func(typ string) {
		switch {
		case meta.IsShapeType(typ):
			shapeType = typ
		case meta.IsArrayType(typ):
			elemType := strings.TrimSuffix(typ, "[]")
			elemTypes = append(elemTypes, meta.Type{Elem: elemType})
		}
	})

	// Try to handle it as a shape assignment.
	if shapeType != "" {
		class, ok := t.state.Info.GetClass(shapeType)
		if ok {
			t.handleAssignShapeToList(list.Items, class)
			return
		}
	}

	// Try to handle it as an array assignment.
	if len(elemTypes) != 0 {
		elemTypeMap := meta.NewTypesMapFromTypes(elemTypes).Immutable()
		for _, item := range list.Items {
			t.handleVariableNode(item.Val, elemTypeMap)
		}
		return
	}

	// Fallback: define vars with unknown types.
	for _, item := range list.Items {
		t.handleVariableNode(item.Val, meta.NewTypesMap("unknown_from_list"))
	}
}

func (t *TypesWalker) handleAssignShapeToList(items []*ir.ArrayItemExpr, info meta.ClassInfo) {
	for i, item := range items {
		var prop meta.PropertyInfo
		var ok bool

		if item.Key != nil {
			var key string
			switch keyNode := item.Key.(type) {
			case *ir.String:
				key = keyNode.Value
			case *ir.Lnumber:
				key = keyNode.Value
			case *ir.Dnumber:
				key = keyNode.Value
			}

			if key != "" {
				prop, ok = info.Properties[key]
			}
		} else {
			prop, ok = info.Properties[fmt.Sprint(i)]
		}

		var tp meta.TypesMap
		if !ok {
			tp = meta.NewTypesMap("unknown_from_list")
		} else {
			tp = prop.Typ
		}
		t.handleVariableNode(item.Val, tp)
	}
}

func (t *TypesWalker) handleDimFetchLValue(e *ir.ArrayDimFetchExpr, typ meta.TypesMap) {
	switch v := e.Variable.(type) {
	case *ir.Var, *ir.SimpleVar:
		arrTyp := typ.Map(meta.WrapArrayOf)
		t.addVar(v, arrTyp)
		t.handleVariable(v)
	case *ir.ArrayDimFetchExpr:
		t.handleDimFetchLValue(v, meta.MixedType)
	default:
		// probably not assignable?
		v.Walk(t)
	}

	if e.Dim != nil {
		e.Dim.Walk(t)
	}
}

func (t *TypesWalker) handleVariableNode(n ir.Node, typ meta.TypesMap) {
	if n == nil {
		return
	}

	t.Types[n] = NewType(typ)

	var vv ir.Node
	switch n := n.(type) {
	case *ir.Var, *ir.SimpleVar:
		vv = n
	case *ir.ReferenceExpr:
		vv = n.Variable
		t.Types[vv] = NewType(typ)
	default:
		return
	}

	t.addVar(vv, typ)
}

func (t *TypesWalker) handleVariable(v ir.Node) bool {
	have := t.scope.HaveVar(v)

	if !have /*&& !t.inArrowFunction*/ {
		t.addVar(v, meta.NewTypesMap("undefined"))
		t.Types[v] = NewType(meta.NewTypesMap("undefined"))
		return false
	}

	switch v := v.(type) {
	case *ir.SimpleVar:
		typ, ok := t.scope.GetVarNameType(v.Name)
		if ok {
			t.Types[v] = NewType(typ)
		}
	case *ir.Var:
		typ := solver.ExprType(t.scope, t.state, v)
		t.Types[v] = NewType(typ)
	}

	return false
}

func (t *TypesWalker) addVarName(nm string, typ meta.TypesMap) {
	t.scope.AddVarName(nm, typ, "", meta.VarAlwaysDefined)
}

func (t *TypesWalker) addVar(v ir.Node, typ meta.TypesMap) {
	t.Types[v] = NewType(typ)
	t.scope.AddVar(v, typ, "", meta.VarAlwaysDefined)
}

func (t *TypesWalker) replaceVar(v ir.Node, typ meta.TypesMap) {
	t.Types[v] = NewType(typ)
	t.scope.ReplaceVar(v, typ, "", meta.VarAlwaysDefined)
}

func (t *TypesWalker) maybeAddAllVars(sc *meta.Scope) {
	sc.Iterate(func(varName string, typ meta.TypesMap, flags meta.VarFlags) {
		flags &^= meta.VarAlwaysDefined
		t.scope.AddVarName(varName, typ, "", flags)
	})
}

func (t *TypesWalker) withNewContext(action func()) *meta.Scope {
	oldCtx := t.scope
	newCtx := t.scope.Clone()

	t.scope = newCtx
	action()
	t.scope = oldCtx

	return newCtx
}

func (t *TypesWalker) LeaveNode(n ir.Node) {

}
