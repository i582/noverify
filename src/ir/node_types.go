package ir

import (
	"github.com/VKCOM/noverify/src/php/parser/pkg/position"
	"github.com/VKCOM/noverify/src/php/parser/pkg/token"
	"github.com/VKCOM/noverify/src/phpdoc"
)

// TODO: make Alt and AltSyntax field names consistent.

// Assign is a `$Variable = $Expression` expression.
type Assign struct {
	EqualTkn   *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
}

// AssignBitwiseAnd is a `$Variable &= $Expression` expression.
type AssignBitwiseAnd struct {
	EqualTkn   *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
}

// AssignBitwiseOr is a `$Variable |= $Expression` expression.
type AssignBitwiseOr struct {
	EqualTkn   *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
}

// AssignBitwiseXor is a `$Variable ^= $Expression` expression.
type AssignBitwiseXor struct {
	EqualTkn   *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
}

// AssignCoalesce is a `$Variable ??= $Expression` expression.
type AssignCoalesce struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignConcat is a `$Variable .= $Expression` expression.
type AssignConcat struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignDiv is a `$Variable /= $Expression` expression.
type AssignDiv struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignMinus is a `$Variable -= $Expression` expression.
type AssignMinus struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignMod is a `$Variable %= $Expression` expression.
type AssignMod struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignMul is a `$Variable *= $Expression` expression.
type AssignMul struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignPlus is a `$Variable += $Expression` expression.
type AssignPlus struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignPow is a `$Variable **= $Expression` expression.
type AssignPow struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignReference is a `$Variable &= $Expression` expression.
type AssignReference struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignShiftLeft is a `$Variable <<= $Expression` expression.
type AssignShiftLeft struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AssignShiftRight is a `$Variable >>= $Expression` expression.
type AssignShiftRight struct {
	Token      *token.Token
	Position   *position.Position
	Variable   Node
	Expression Node
	EqualTkn   *token.Token
}

// AnonClassExpr is an anonymous class expression.
// $Args may contain constructor call arguments `new class ($Args...) {}`.
type AnonClassExpr struct {
	Token    *token.Token
	Position *position.Position
	Class

	ArgsToken *token.Token
	Args      []Node
}

// BitwiseAndExpr is a `$Left & $Right` expression.
type BitwiseAndExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// BitwiseOrExpr is a `$Left | $Right` expression.
type BitwiseOrExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// BitwiseXorExpr is a `$Left ^ $Right` expression.
type BitwiseXorExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// BooleanAndExpr is a `$Left && $Right` expression.
type BooleanAndExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// BooleanOrExpr is a `$Left || $Right` expression.
type BooleanOrExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// CoalesceExpr is a `$Left ?? $Right` expression.
type CoalesceExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// ConcatExpr is a `$Left . $Right` expression.
type ConcatExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// DivExpr is a `$Left / $Right` expression.
type DivExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// EqualExpr is a `$Left == $Right` expression.
type EqualExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// GreaterExpr is a `$Left > $Right` expression.
type GreaterExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// GreaterOrEqualExpr is a `$Left >= $Right` expression.
type GreaterOrEqualExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// IdenticalExpr is a `$Left === $Right` expression.
type IdenticalExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// LogicalAndExpr is a `$Left and $Right` expression.
type LogicalAndExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// LogicalOrExpr is a `$Left or $Right` expression.
type LogicalOrExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// LogicalXorExpr is a `$Left xor $Right` expression.
type LogicalXorExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// MinusExpr is a `$Left - $Right` expression.
type MinusExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// ModExpr is a `$Left % $Right` expression.
type ModExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// MulExpr is a `$Left * $Right` expression.
type MulExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// NotEqualExpr is a `$Left != $Right` expression.
type NotEqualExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// NotIdenticalExpr is a `$Left !== $Right` expression.
type NotIdenticalExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// PlusExpr is a `$Left + $Right` expression.
type PlusExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// PowExpr is a `$Left ** $Right` expression.
type PowExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// ShiftLeftExpr is a `$Left << $Right` expression.
type ShiftLeftExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// ShiftRightExpr is a `$Left >> $Right` expression.
type ShiftRightExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// SmallerExpr is a `$Left < $Right` expression.
type SmallerExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// SmallerOrEqualExpr is a `$Left <= $Right` expression.
type SmallerOrEqualExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// SpaceshipExpr is a `$Left <=> $Right` expression.
type SpaceshipExpr struct {
	Token    *token.Token
	Position *position.Position
	Left     Node
	Right    Node
	OpTkn    *token.Token
}

// TypeCastExpr is a `($Type)$Expr` expression.
type TypeCastExpr struct {
	Token    *token.Token
	Position *position.Position
	Type     string // "array" "bool" "int" "float" "object" "string"
	Expr     Node
}

// UnsetCastExpr is a `(unset)$Expr` expression.
type UnsetCastExpr struct {
	CastTkn  *token.Token
	Position *position.Position
	Expr     Node
}

// ArrayExpr is a `array($Items...)` expression.
// If $ShortSyntax is true, it's `[$Items...]`.
type ArrayExpr struct {
	ArrayTkn        *token.Token
	OpenBracketTkn  *token.Token
	Position        *position.Position
	Items           []*ArrayItemExpr
	ShortSyntax     bool
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
}

// ArrayDimFetchExpr is a `$Variable[$Dim]` expression.
type ArrayDimFetchExpr struct {
	Position        *position.Position
	Variable        Node
	OpenBracketTkn  *token.Token
	Dim             Node
	CloseBracketTkn *token.Token
}

// ArrayItemExpr is a `$Key => $Val` expression.
// If $Unpack is true, it's `...$Val` ($Key is nil).
//
// TODO: make unpack a separate node?
type ArrayItemExpr struct {
	Position       *position.Position
	EllipsisTkn    *token.Token
	Key            Node
	DoubleArrowTkn *token.Token
	AmpersandTkn   *token.Token
	Val            Node
	Unpack         bool
}

// ArrowFunctionExpr is a `fn($Params...): $ReturnType => $Expr` expression.
// If $ReturnsRef is true, it's `fn&($Params...): $ReturnType => $Expr`.
// If $Static is true, it's `static fn($Params...): $ReturnType => $Expr`.
// $ReturnType is optional, without it we have `fn($Params...) => $Expr` syntax.
type ArrowFunctionExpr struct {
	Position           *position.Position
	StaticTkn          *token.Token
	FnTkn              *token.Token
	AmpersandTkn       *token.Token
	OpenParenthesisTkn *token.Token

	ReturnsRef    bool
	Static        bool
	PhpDocComment string

	PhpDoc []phpdoc.CommentPart
	Params []Node

	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token

	ReturnType     Node
	DoubleArrowTkn *token.Token
	Expr           Node
}

// BitwiseNotExpr is a `~$Expr` expression.
type BitwiseNotExpr struct {
	TildaTkn *token.Token
	Position *position.Position
	Expr     Node
}

// BooleanNotExpr is a `!$Expr` expression.
type BooleanNotExpr struct {
	ExclamationTkn *token.Token
	Position       *position.Position
	Expr           Node
}

// ClassConstFetchExpr is a `$Class::$ConstantName` expression.
type ClassConstFetchExpr struct {
	DoubleColonTkn *token.Token
	Position       *position.Position
	Class          Node
	ConstantName   *Identifier
}

// CloneExpr is a `clone $Expr` expression.
type CloneExpr struct {
	CloneTkn *token.Token
	Position *position.Position
	Expr     Node
}

// ClosureExpr is a `function($Params...) use ($ClosureUse) : $ReturnType { $Stmts... }` expression.
// If $ReturnsRef is true, it's `function&($Params...) use ($ClosureUse) : $ReturnType { $Stmts... }`.
// If $Static is true, it's `static function($Params...) use ($ClosureUse) : $ReturnType { $Stmts... }`.
// $ReturnType is optional, without it we have `function($Params...) use ($ClosureUse) { $Stmts... }` syntax.
// $ClosureUse is optional, without it we have `function($Params...) : $ReturnType { $Stmts... }` syntax.
type ClosureExpr struct {
	Position *position.Position

	StaticTkn          *token.Token
	FunctionTkn        *token.Token
	AmpersandTkn       *token.Token
	OpenParenthesisTkn *token.Token

	ReturnsRef    bool
	Static        bool
	PhpDocComment string
	PhpDoc        []phpdoc.CommentPart
	Params        []Node

	SeparatorTkns         []*token.Token
	CloseParenthesisTkn   *token.Token
	UseTkn                *token.Token
	UseOpenParenthesisTkn *token.Token

	ClosureUse *ClosureUseExpr

	UseSeparatorTkns       []*token.Token
	UseCloseParenthesisTkn *token.Token
	ColonTkn               *token.Token

	ReturnType Node

	OpenCurlyBracketTkn *token.Token

	Stmts []Node

	CloseCurlyBracketTkn *token.Token
}

// ClosureUseExpr is a `use ($Uses...)` expression.
// TODO: it's not a expression really.
type ClosureUseExpr struct {
	Token    *token.Token
	Position *position.Position
	Uses     []Node
}

// ConstFetchExpr is a `$Constant` expression.
type ConstFetchExpr struct {
	Position *position.Position
	Constant *Name
}

// EmptyExpr is a `empty($Expr)` expression.
type EmptyExpr struct {
	Position            *position.Position
	EmptyTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Node
	CloseParenthesisTkn *token.Token
}

// ErrorSuppressExpr is a `@$Expr` expression.
type ErrorSuppressExpr struct {
	Position *position.Position
	AtTkn    *token.Token
	Expr     Node
}

// EvalExpr is a `eval($Expr)` expression.
type EvalExpr struct {
	Token               *token.Token
	Position            *position.Position
	EvalTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Node
	CloseParenthesisTkn *token.Token
}

// ExitExpr is a `exit($Expr)` expression.
// If $Die is true, it's `die($Expr)`.
type ExitExpr struct {
	Token               *token.Token
	Die                 bool
	Position            *position.Position
	ExitTkn             *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Node
	CloseParenthesisTkn *token.Token
}

// FunctionCallExpr is a `$Function($Args...)` expression.
type FunctionCallExpr struct {
	Position            *position.Position
	Function            Node
	OpenParenthesisTkn  *token.Token
	Args                []Node
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// ImportExpr is a `$Func $Expr` expression.
// It could be `include $Expr`, `require $Expr` and so on.
type ImportExpr struct {
	Token    *token.Token
	Position *position.Position
	Func     string // "include" "include_once" "require" "require_once"
	Expr     Node
}

// InstanceOfExpr is a `$Expr instanceof $Class` expression.
type InstanceOfExpr struct {
	Position      *position.Position
	InstanceOfTkn *token.Token
	Expr          Node
	Class         Node
}

// IssetExpr is a `isset($Variables...)` expression.
type IssetExpr struct {
	Token               *token.Token
	Position            *position.Position
	IssetTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Variables           []Node
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// ListExpr is a `list($Items...)` expression.
// Note that it may appear not only in assignments as LHS, but
// also in foreach value expressions.
type ListExpr struct {
	Position        *position.Position
	ListTkn         *token.Token
	OpenBracketTkn  *token.Token
	Items           []*ArrayItemExpr
	SeparatorTkns   []*token.Token
	CloseBracketTkn *token.Token
	ShortSyntax     bool
}

// MethodCallExpr is a `$Variable->$Method($Args...)` expression.
type MethodCallExpr struct {
	Position             *position.Position
	Variable             Node
	ObjectOperatorTkn    *token.Token
	OpenCurlyBracketTkn  *token.Token
	Method               Node
	CloseCurlyBracketTkn *token.Token
	OpenParenthesisTkn   *token.Token
	Args                 []Node
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
}

// NewExpr is a `new $Class($Args...)` expression.
// If $Args is nil, it's `new $Class`.
type NewExpr struct {
	Position            *position.Position
	NewTkn              *token.Token
	Class               Node
	OpenParenthesisTkn  *token.Token
	Args                []Node
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
}

// ParenExpr is a `($Expr)` expression.
type ParenExpr struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// PostDecExpr is a `$Variable--` expression.
type PostDecExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
}

// PostIncExpr is a `$Variable++` expression.
type PostIncExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
}

// PreDecExpr is a `--$Variable` expression.
type PreDecExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
}

// PreIncExpr is a `++$Variable` expression.
type PreIncExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
}

// PrintExpr is a `print $Expr` expression.
type PrintExpr struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// PropertyFetchExpr is a `$Variable->$Property` expression.
type PropertyFetchExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
	Property Node
}

// ReferenceExpr is a `&$Variable` expression.
type ReferenceExpr struct {
	Token    *token.Token
	Position *position.Position
	Variable Node
}

// ShellExecExpr is a ``-quoted string.
type ShellExecExpr struct {
	Token    *token.Token
	Position *position.Position
	Parts    []Node
}

// StaticCallExpr is a `$Class::$Call($Args...)` expression.
type StaticCallExpr struct {
	Token    *token.Token
	Position *position.Position
	Class    Node
	Call     Node

	ArgsToken *token.Token
	Args      []Node
}

// StaticPropertyFetchExpr is a `$Class::$Property` expression.
type StaticPropertyFetchExpr struct {
	Token    *token.Token
	Position *position.Position
	Class    Node
	Property Node
}

// TernaryExpr is a `$Condition ? $IfTrue : $IfFalse` expression.
// If $IfTrue is nil, it's `$Condition ?: $IfFalse`.
type TernaryExpr struct {
	Token     *token.Token
	Position  *position.Position
	Condition Node
	IfTrue    Node
	IfFalse   Node
}

// UnaryMinusExpr is a `-$Expr` expression.
type UnaryMinusExpr struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// UnaryPlusExpr is a `+$Expr` expression.
type UnaryPlusExpr struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// YieldExpr is a `yield $Key => $Value` expression.
// If $Key is nil, it's `yield $Value`.
type YieldExpr struct {
	Token    *token.Token
	Position *position.Position
	Key      Node
	Value    Node
}

// YieldFromExpr is a `yield from $Expr` expression.
type YieldFromExpr struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// Name is either a FQN, local name or a name that may need a further resolving.
// Use Name methods to interpret the $Value correctly.
type Name struct {
	Token    *token.Token
	Position *position.Position
	Value    string
}

// Argument is a wrapper node for func/method arguments.
// If $Variadic is true, it's `...$Expr`.
// If $IsReference is true, it's `&$Expr`.
type Argument struct {
	Token       *token.Token
	Position    *position.Position
	Variadic    bool
	IsReference bool
	Expr        Node
}

// Identifier is like a name, but it's always resolved to itself.
// Identifier always consists of a single part.
type Identifier struct {
	Token    *token.Token
	Position *position.Position
	Value    string
}

// Nullable is a `?$Expr` expression.
type Nullable struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// Parameter is a function param declaration.
// Possible syntaxes:
// $VariableType $Variable = $DefaultValue
// $VariableType $Variable
// $Variable
// If $ByRef is true, it's `&$Variable`.
// If $Variadic is true, it's `...$Variable`.
type Parameter struct {
	Token        *token.Token
	Position     *position.Position
	ByRef        bool
	Variadic     bool
	VariableType Node
	Variable     *SimpleVar
	DefaultValue Node
}

// Root is a node that wraps all file statements.
type Root struct {
	EndTkn   *token.Token
	Position *position.Position
	Stmts    []Node
}

// SimpleVar is a normal PHP variable like `$foo` or `$bar`.
type SimpleVar struct {
	Token    *token.Token
	Position *position.Position
	Name     string
}

// Var is variable variable expression like `$$foo` or `${"foo"}`.
type Var struct {
	Token    *token.Token
	Position *position.Position
	Expr     Node
}

// Dnumber is a floating point literal.
type Dnumber struct {
	NumberTkn *token.Token
	Position  *position.Position
	Value     string
}

// Encapsed is a string literal with interpolated parts.
type Encapsed struct {
	Position      *position.Position
	OpenQuoteTkn  *token.Token
	Parts         []Node
	CloseQuoteTkn *token.Token
}

// EncapsedStringPart is a part of the Encapsed literal.
type EncapsedStringPart struct {
	Position       *position.Position
	EncapsedStrTkn *token.Token
	Value          string
}

// Heredoc is special PHP literal.
// Note that it may be a nowdoc, depending on the label.
type Heredoc struct {
	Position        *position.Position
	Label           string
	OpenHeredocTkn  *token.Token
	Parts           []Node
	CloseHeredocTkn *token.Token
}

// Lnumber is an integer literal.
type Lnumber struct {
	NumberTkn *token.Token
	Position  *position.Position
	Value     string
}

// MagicConstant is a special PHP constant like __FILE__ or __CLASS__.
// TODO: do we really need a separate node for these constants?
type MagicConstant struct {
	MagicConstTkn *token.Token
	Position      *position.Position
	Value         string
}

// String is a simple (no interpolation) string literal.
//
// The $Value contains interpreted string bytes, if you need a raw
// string value, use positions and fetch relevant the source bytes.
//
// $DoubleQuotes tell whether originally this string literal was ""-quoted.
type String struct {
	Token        *token.Token
	Position     *position.Position
	Value        string
	DoubleQuotes bool
}

// BadString is a string that we couldn't interpret correctly.
// The $Value contains uninterpreted (raw) string bytes.
// $Error contains the reason why this string is "bad".
type BadString struct {
	Token        *token.Token
	Position     *position.Position
	Value        string
	DoubleQuotes bool
	Error        string
}

// BreakStmt is a `break $Expr` statement.
type BreakStmt struct {
	Position     *position.Position
	BreakTkn     *token.Token
	Expr         Node
	SemiColonTkn *token.Token
}

// CaseStmt is a `case $Cond: $Stmts...` statement.
type CaseStmt struct {
	Position         *position.Position
	CaseTkn          *token.Token
	Cond             Node
	CaseSeparatorTkn *token.Token
	Stmts            []Node
}

// CaseListStmt is a wrapper node that contains all switch statement cases.
// TODO: can we get rid of it?
type CaseListStmt struct {
	Token    *token.Token
	Position *position.Position
	Cases    []Node
}

// CatchStmt is a `catch ($Types... $Variable) { $Stmts... }` statement.
// Note that $Types are |-separated, like in `T1 | T2`.
type CatchStmt struct {
	Position             *position.Position
	CatchTkn             *token.Token
	OpenParenthesisTkn   *token.Token
	Types                []Node
	SeparatorTkns        []*token.Token
	Variable             *SimpleVar
	CloseParenthesisTkn  *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Node
	CloseCurlyBracketTkn *token.Token
}

// ClassStmt is a named class declaration.
// $Modifiers consist of identifiers like `final` and `abstract`.
type ClassStmt struct {
	Token     *token.Token
	Position  *position.Position
	ClassName *Identifier
	Modifiers []*Identifier
	Class
}

// ClassConstListStmt is a `$Modifiers... const $Consts...` statement.
// $Modifiers may specify the constant access level.
// Every element in $Consts is a *ConstantStmt.
type ClassConstListStmt struct {
	Position      *position.Position
	Modifiers     []*Identifier
	ConstTkn      *token.Token
	Consts        []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// ClassExtendsStmt is a `extends $ClassName` statement.
type ClassExtendsStmt struct {
	Token     *token.Token
	Position  *position.Position
	ClassName *Name
}

// ClassImplementsStmt is a `implements $InterfaceNames...` statement.
// TODO: shouldn't every InterfaceName be a *Name?
type ClassImplementsStmt struct {
	Token          *token.Token
	Position       *position.Position
	InterfaceNames []Node
}

// ClassMethodStmt is a class method declaration.
type ClassMethodStmt struct {
	Position            *position.Position
	Modifiers           []*Identifier
	FunctionTkn         *token.Token
	AmpersandTkn        *token.Token
	MethodName          *Identifier
	OpenParenthesisTkn  *token.Token
	Params              []Node
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	ReturnType          Node
	Stmt                Node
	ReturnsRef          bool
	PhpDocComment       string
}

// ConstListStmt is a `const $Consts` statement.
// Every element in $Consts is a *ConstantStmt.
type ConstListStmt struct {
	Position      *position.Position
	ConstTkn      *token.Token
	Consts        []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// ConstantStmt is a `$ConstantName = $Expr` statement.
// It's a part of the *ConstListStmt, *ClassConstListStmt and *DeclareStmt.
type ConstantStmt struct {
	Position      *position.Position
	PhpDocComment string
	ConstantName  *Identifier
	EqualTkn      *token.Token
	Expr          Node
}

// ContinueStmt is a `continue $Expe` statement.
type ContinueStmt struct {
	Position     *position.Position
	ContinueTkn  *token.Token
	Expr         Node
	SemiColonTkn *token.Token
}

// DeclareStmt is a `declare ($Consts...) $Stmt` statement.
// $Stmt can be an empty statement, like in `declare ($Consts...);`,
// but it can also be a block like in `declare ($Consts...) {}`.
// If $Alt is true, the block will begin with `:` and end with `enddeclare`.
// Every element in $Consts is a *ConstantStmt.
type DeclareStmt struct {
	Position            *position.Position
	DeclareTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Consts              []Node
	SeparatorTkns       []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Node
	EndDeclareTkn       *token.Token
	SemiColonTkn        *token.Token
	Alt                 bool
}

// DefaultStmt is a `default: $Stmts...` statement.
type DefaultStmt struct {
	Position         *position.Position
	DefaultTkn       *token.Token
	CaseSeparatorTkn *token.Token
	Stmts            []Node
}

// DoStmt is a `do $Stmt while ($Cond)` statement.
type DoStmt struct {
	Position            *position.Position
	DoTkn               *token.Token
	Stmt                Node
	WhileTkn            *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Node
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

// EchoStmt is a `echo $Exprs...` statement.
type EchoStmt struct {
	Position      *position.Position
	EchoTkn       *token.Token
	Exprs         []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// ElseStmt is a `else $Stmt` statement.
// If $AltSyntax is true, the block will begin with `:`.
type ElseStmt struct {
	Position  *position.Position
	ElseTkn   *token.Token
	ColonTkn  *token.Token
	Stmt      Node
	AltSyntax bool
}

// ElseIfStmt is a `elseif ($Cond) $Stmt` statement.
// If $AltSyntax is true, the block will begin with `:` and end with `endif`.
// $Merged tells whether this elseif is a merged `else if` statement.
type ElseIfStmt struct {
	Token               *token.Token
	Position            *position.Position
	ElseIfTkn           *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Node
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Node
	AltSyntax           bool
	Merged              bool
}

// ExpressionStmt is an expression $Expr that is evaluated for side-effects only.
// When expression is used in a place where statement is expected, it
// becomes an ExpressionStmt.
type ExpressionStmt struct {
	Position     *position.Position
	Expr         Node
	SemiColonTkn *token.Token
}

// FinallyStmt is a `finally { $Stmts... }` statement.
type FinallyStmt struct {
	Position             *position.Position
	FinallyTkn           *token.Token
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Node
	CloseCurlyBracketTkn *token.Token
}

// ForStmt is a `for ($Init; $Cond; $Loop) $Stmt` statement.
// If $AltSyntax is true, the block will begin with `:` and end with `endfor`.
type ForStmt struct {
	Position            *position.Position
	ForTkn              *token.Token
	OpenParenthesisTkn  *token.Token
	Init                []Node
	InitSeparatorTkns   []*token.Token
	InitSemiColonTkn    *token.Token
	Cond                []Node
	CondSeparatorTkns   []*token.Token
	CondSemiColonTkn    *token.Token
	Loop                []Node
	LoopSeparatorTkns   []*token.Token
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Node
	EndForTkn           *token.Token
	SemiColonTkn        *token.Token
	AltSyntax           bool
}

// ForeachStmt is a `foreach ($Expr as $Key => $Variable) $Stmt` statement.
// If $AltSyntax is true, the block will begin with `:` and end with `endforeach`.
type ForeachStmt struct {
	Position            *position.Position
	ForeachTkn          *token.Token
	OpenParenthesisTkn  *token.Token
	Expr                Node
	AsTkn               *token.Token
	Key                 Node
	DoubleArrowTkn      *token.Token
	AmpersandTkn        *token.Token
	Variable            Node
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Node
	EndForeachTkn       *token.Token
	SemiColonTkn        *token.Token
	AltSyntax           bool
}

// FunctionStmt is a named function declaration.
type FunctionStmt struct {
	Position             *position.Position
	FunctionTkn          *token.Token
	AmpersandTkn         *token.Token
	FunctionName         *Identifier
	OpenParenthesisTkn   *token.Token
	Params               []Node
	SeparatorTkns        []*token.Token
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	ReturnType           Node
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Node
	CloseCurlyBracketTkn *token.Token
	ReturnsRef           bool
	PhpDocComment        string
	PhpDoc               []phpdoc.CommentPart
}

// GlobalStmt is a `global $Vars` statement.
type GlobalStmt struct {
	Position      *position.Position
	GlobalTkn     *token.Token
	Vars          []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// GotoStmt is a `goto $Label` statement.
type GotoStmt struct {
	Token        *token.Token
	Position     *position.Position
	GotoTkn      *token.Token
	Label        *Identifier
	SemiColonTkn *token.Token
}

// GroupUseStmt is a `use $UseType $Prefix\{ $UseList }` statement.
// $UseType is a "function" or "const".
// TODO: change $UseType type to *Identifier?
type GroupUseStmt struct {
	Position              *position.Position
	UseTkn                *token.Token
	UseType               *Identifier
	LeadingNsSeparatorTkn *token.Token
	Prefix                *Name
	NsSeparatorTkn        *token.Token
	OpenCurlyBracketTkn   *token.Token
	UseList               []Node
	SeparatorTkns         []*token.Token
	CloseCurlyBracketTkn  *token.Token
	SemiColonTkn          *token.Token
}

// HaltCompilerStmt is a `__halt_compiler()` statement.
type HaltCompilerStmt struct {
	Position            *position.Position
	HaltCompilerTkn     *token.Token
	OpenParenthesisTkn  *token.Token
	CloseParenthesisTkn *token.Token
	SemiColonTkn        *token.Token
}

// IfStmt is a `if ($Cond) $Stmt` statement.
// $ElseIf contains an entire elseif chain, if any.
// $Else may contain an else part of the statement.
type IfStmt struct {
	Position            *position.Position
	IfTkn               *token.Token
	OpenParenthesisTkn  *token.Token
	Cond                Node
	CloseParenthesisTkn *token.Token
	ColonTkn            *token.Token
	Stmt                Node
	ElseIf              []Node
	Else                Node
	EndIfTkn            *token.Token
	SemiColonTkn        *token.Token
	AltSyntax           bool
}

// InlineHtmlStmt is a part of the script that will not be interpreted
// as a PHP script. In other words, it's everything outside of
// the <? and ?> tags.
type InlineHtmlStmt struct {
	Position      *position.Position
	InlineHtmlTkn *token.Token
	Value         string
}

// InterfaceStmt is an interface declaration.
type InterfaceStmt struct {
	Token         *token.Token
	Position      *position.Position
	PhpDocComment string
	InterfaceName *Identifier
	Extends       *InterfaceExtendsStmt
	Stmts         []Node
}

// InterfaceExtendsStmt is a `extends $InterfaceNames...` statement.
// TODO: do we need this wrapper node?
// TODO: InterfaceNames could be a []*Name.
type InterfaceExtendsStmt struct {
	Token          *token.Token
	Position       *position.Position
	InterfaceNames []Node
}

// LabelStmt is a `$LabelName:` statement.
type LabelStmt struct {
	Position  *position.Position
	LabelName *Identifier
	ColonTkn  *token.Token
}

// NamespaceStmt is a `namespace $NamespaceName` statement.
// If $Stmts is not nil, it's `namespace $NamespaceName { $Stmts... }`.
type NamespaceStmt struct {
	Token         *token.Token
	Position      *position.Position
	NamespaceName *Name
	Stmts         []Node
}

// NopStmt is a `;` statement.
// It's also known as "empty statement".
type NopStmt struct {
	Position     *position.Position
	SemiColonTkn *token.Token
}

// PropertyStmt is a `$Variable = $Expr` statement.
// It's a part of the *PropertyListStmt.
type PropertyStmt struct {
	Position      *position.Position
	PhpDocComment string
	PhpDoc        []phpdoc.CommentPart
	Variable      *SimpleVar
	EqualTkn      *token.Token
	Expr          Node
}

// PropertyListStmt is a `$Modifiers $Type $Properties` statement.
// Every element in $Properties is a *PropertyStmt.
type PropertyListStmt struct {
	Position      *position.Position
	Modifiers     []*Identifier
	Type          Node
	Properties    []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// ReturnStmt is a `return $Expr` statement.
type ReturnStmt struct {
	Token        *token.Token
	Position     *position.Position
	ReturnTkn    *token.Token
	Expr         Node
	SemiColonTkn *token.Token
}

// StaticStmt is a `static $Vars...` statement.
// Every element in $Vars is a *StaticVarStmt.
type StaticStmt struct {
	Position      *position.Position
	StaticTkn     *token.Token
	Vars          []Node
	SeparatorTkns []*token.Token
	SemiColonTkn  *token.Token
}

// StaticVarStmt is a `$Variable = $Expr`.
// It's a part of the *StaticStmt.
type StaticVarStmt struct {
	Position *position.Position
	Variable *SimpleVar
	EqualTkn *token.Token
	Expr     Node
}

// StmtList is a `{ $Stmts... }` statement.
// It's also known as "block statement".
type StmtList struct {
	Position             *position.Position
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Node
	CloseCurlyBracketTkn *token.Token
}

// SwitchStmt is a `switch ($Cond) $CaseList` statement.
// If $AltSyntax is true, the block will begin with `:` and end with `endswitch`.
type SwitchStmt struct {
	Position             *position.Position
	SwitchTkn            *token.Token
	OpenParenthesisTkn   *token.Token
	Cond                 Node
	CloseParenthesisTkn  *token.Token
	ColonTkn             *token.Token
	OpenCurlyBracketTkn  *token.Token
	CaseSeparatorTkn     *token.Token
	CaseList             []Node // *CaseListStmt
	CloseCurlyBracketTkn *token.Token
	EndSwitchTkn         *token.Token
	SemiColonTkn         *token.Token
	AltSyntax            bool
}

// ThrowStmt is a `throw $Expr` statement.
type ThrowStmt struct {
	Position     *position.Position
	ThrowTkn     *token.Token
	Expr         Node
	SemiColonTkn *token.Token
}

// TraitStmt is a trait declaration.
type TraitStmt struct {
	Position             *position.Position
	PhpDocComment        string
	TraitTkn             *token.Token
	TraitName            *Identifier
	OpenCurlyBracketTkn  *token.Token
	Stmts                []Node
	CloseCurlyBracketTkn *token.Token
}

// TraitAdaptationListStmt is a block inside a *TraitUseStmt.
type TraitAdaptationListStmt struct {
	Token       *token.Token
	Position    *position.Position
	Adaptations []Node
}

type TraitMethodRefStmt struct {
	Token    *token.Token
	Position *position.Position
	Trait    Node
	Method   *Identifier
}

type TraitUseStmt struct {
	Token               *token.Token
	Position            *position.Position
	Traits              []Node
	TraitAdaptationList Node
}

type TraitUseAliasStmt struct {
	Token    *token.Token
	Position *position.Position
	Ref      Node
	Modifier Node
	Alias    *Identifier
}

type TraitUsePrecedenceStmt struct {
	Token     *token.Token
	Position  *position.Position
	Ref       Node
	Insteadof []Node
}

// TryStmt is a `try { $Stmts... } $Catches` statement.
// $Finally only presents if `finally {...}` block exists.
type TryStmt struct {
	Token    *token.Token
	Position *position.Position
	Stmts    []Node
	Catches  []Node
	Finally  Node
}

// UnsetStmt is a `unset($Vars...)` statement.
type UnsetStmt struct {
	Token    *token.Token
	Position *position.Position
	Vars     []Node
}

type UseStmt struct {
	Token    *token.Token
	Position *position.Position
	UseType  *Identifier
	Use      *Name
	Alias    *Identifier
}

type UseListStmt struct {
	Token    *token.Token
	Position *position.Position
	UseType  *Identifier
	Uses     []Node
}

// WhileStmt is a `while ($Cond) $Stmt` statement.
// If $AltSyntax is true, the block will begin with `:` and end with `endwhile`.
type WhileStmt struct {
	Token     *token.Token
	Position  *position.Position
	Cond      Node
	Stmt      Node
	AltSyntax bool
}
