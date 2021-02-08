// Code generated by the `ir/codegen` package. DO NOT EDIT.
package irutil

import (
	"fmt"
	"github.com/VKCOM/noverify/src/ir"
)

func NodeEqual(x, y ir.Node) bool {
	if x == nil || y == nil {
		return x == y
	}
	switch x := x.(type) {
	case *ir.AnonClassExpr:
		y, ok := y.(*ir.AnonClassExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Args, y.Args) {
			return false
		}
		if !classEqual(x.Class, y.Class) {
			return false
		}
		return true
	case *ir.Argument:
		y, ok := y.(*ir.Argument)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if x.Variadic != y.Variadic {
			return false
		}
		if x.IsReference != y.IsReference {
			return false
		}
		return true
	case *ir.ArrayDimFetchExpr:
		y, ok := y.(*ir.ArrayDimFetchExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Dim, y.Dim) {
			return false
		}
		if x.CurlyBrace != y.CurlyBrace {
			return false
		}
		return true
	case *ir.ArrayExpr:
		y, ok := y.(*ir.ArrayExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Items) != len(y.Items) {
			return false
		}
		for i := range x.Items {
			if !NodeEqual(x.Items[i], y.Items[i]) {
				return false
			}
		}
		if x.ShortSyntax != y.ShortSyntax {
			return false
		}
		return true
	case *ir.ArrayItemExpr:
		y, ok := y.(*ir.ArrayItemExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Key, y.Key) {
			return false
		}
		if !NodeEqual(x.Val, y.Val) {
			return false
		}
		if x.Unpack != y.Unpack {
			return false
		}
		return true
	case *ir.ArrowFunctionExpr:
		y, ok := y.(*ir.ArrowFunctionExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Params, y.Params) {
			return false
		}
		if !NodeEqual(x.ReturnType, y.ReturnType) {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if x.ReturnsRef != y.ReturnsRef {
			return false
		}
		if x.Static != y.Static {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.Assign:
		y, ok := y.(*ir.Assign)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignBitwiseAnd:
		y, ok := y.(*ir.AssignBitwiseAnd)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignBitwiseOr:
		y, ok := y.(*ir.AssignBitwiseOr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignBitwiseXor:
		y, ok := y.(*ir.AssignBitwiseXor)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignCoalesce:
		y, ok := y.(*ir.AssignCoalesce)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignConcat:
		y, ok := y.(*ir.AssignConcat)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignDiv:
		y, ok := y.(*ir.AssignDiv)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignMinus:
		y, ok := y.(*ir.AssignMinus)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignMod:
		y, ok := y.(*ir.AssignMod)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignMul:
		y, ok := y.(*ir.AssignMul)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignPlus:
		y, ok := y.(*ir.AssignPlus)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignPow:
		y, ok := y.(*ir.AssignPow)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignReference:
		y, ok := y.(*ir.AssignReference)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignShiftLeft:
		y, ok := y.(*ir.AssignShiftLeft)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.AssignShiftRight:
		y, ok := y.(*ir.AssignShiftRight)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expression, y.Expression) {
			return false
		}
		return true
	case *ir.BadString:
		y, ok := y.(*ir.BadString)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		if x.DoubleQuotes != y.DoubleQuotes {
			return false
		}
		if x.Error != y.Error {
			return false
		}
		return true
	case *ir.BitwiseAndExpr:
		y, ok := y.(*ir.BitwiseAndExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.BitwiseNotExpr:
		y, ok := y.(*ir.BitwiseNotExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.BitwiseOrExpr:
		y, ok := y.(*ir.BitwiseOrExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.BitwiseXorExpr:
		y, ok := y.(*ir.BitwiseXorExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.BooleanAndExpr:
		y, ok := y.(*ir.BooleanAndExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.BooleanNotExpr:
		y, ok := y.(*ir.BooleanNotExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.BooleanOrExpr:
		y, ok := y.(*ir.BooleanOrExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.BreakStmt:
		y, ok := y.(*ir.BreakStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.CaseStmt:
		y, ok := y.(*ir.CaseStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.CatchStmt:
		y, ok := y.(*ir.CatchStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Types, y.Types) {
			return false
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.ClassConstFetchExpr:
		y, ok := y.(*ir.ClassConstFetchExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Class, y.Class) {
			return false
		}
		if !NodeEqual(x.ConstantName, y.ConstantName) {
			return false
		}
		return true
	case *ir.ClassConstListStmt:
		y, ok := y.(*ir.ClassConstListStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Modifiers) != len(y.Modifiers) {
			return false
		}
		for i := range x.Modifiers {
			if !NodeEqual(x.Modifiers[i], y.Modifiers[i]) {
				return false
			}
		}
		if !NodeSliceEqual(x.Consts, y.Consts) {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.ClassExtendsStmt:
		y, ok := y.(*ir.ClassExtendsStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.ClassName, y.ClassName) {
			return false
		}
		return true
	case *ir.ClassImplementsStmt:
		y, ok := y.(*ir.ClassImplementsStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.InterfaceNames, y.InterfaceNames) {
			return false
		}
		return true
	case *ir.ClassMethodStmt:
		y, ok := y.(*ir.ClassMethodStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Modifiers) != len(y.Modifiers) {
			return false
		}
		for i := range x.Modifiers {
			if !NodeEqual(x.Modifiers[i], y.Modifiers[i]) {
				return false
			}
		}
		if !NodeEqual(x.MethodName, y.MethodName) {
			return false
		}
		if !NodeSliceEqual(x.Params, y.Params) {
			return false
		}
		if !NodeEqual(x.ReturnType, y.ReturnType) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.ReturnsRef != y.ReturnsRef {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.ClassStmt:
		y, ok := y.(*ir.ClassStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Modifiers) != len(y.Modifiers) {
			return false
		}
		for i := range x.Modifiers {
			if !NodeEqual(x.Modifiers[i], y.Modifiers[i]) {
				return false
			}
		}
		if !NodeEqual(x.ClassName, y.ClassName) {
			return false
		}
		if !classEqual(x.Class, y.Class) {
			return false
		}
		return true
	case *ir.CloneExpr:
		y, ok := y.(*ir.CloneExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.ClosureExpr:
		y, ok := y.(*ir.ClosureExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Params, y.Params) {
			return false
		}
		if !NodeEqual(x.ClosureUse, y.ClosureUse) {
			return false
		}
		if !NodeEqual(x.ReturnType, y.ReturnType) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		if x.ReturnsRef != y.ReturnsRef {
			return false
		}
		if x.Static != y.Static {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.ClosureUsesExpr:
		y, ok := y.(*ir.ClosureUsesExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Uses, y.Uses) {
			return false
		}
		return true
	case *ir.CoalesceExpr:
		y, ok := y.(*ir.CoalesceExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.ConcatExpr:
		y, ok := y.(*ir.ConcatExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.ConstFetchExpr:
		y, ok := y.(*ir.ConstFetchExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Constant, y.Constant) {
			return false
		}
		return true
	case *ir.ConstListStmt:
		y, ok := y.(*ir.ConstListStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Consts, y.Consts) {
			return false
		}
		return true
	case *ir.ConstantStmt:
		y, ok := y.(*ir.ConstantStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.ConstantName, y.ConstantName) {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.ContinueStmt:
		y, ok := y.(*ir.ContinueStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.DeclareStmt:
		y, ok := y.(*ir.DeclareStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Consts, y.Consts) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.Alt != y.Alt {
			return false
		}
		return true
	case *ir.DefaultStmt:
		y, ok := y.(*ir.DefaultStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.DivExpr:
		y, ok := y.(*ir.DivExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.Dnumber:
		y, ok := y.(*ir.Dnumber)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.DoStmt:
		y, ok := y.(*ir.DoStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		return true
	case *ir.EchoStmt:
		y, ok := y.(*ir.EchoStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Exprs, y.Exprs) {
			return false
		}
		return true
	case *ir.ElseIfStmt:
		y, ok := y.(*ir.ElseIfStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		if x.Merged != y.Merged {
			return false
		}
		return true
	case *ir.ElseStmt:
		y, ok := y.(*ir.ElseStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.EmptyExpr:
		y, ok := y.(*ir.EmptyExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.Encapsed:
		y, ok := y.(*ir.Encapsed)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Parts, y.Parts) {
			return false
		}
		return true
	case *ir.EncapsedStringPart:
		y, ok := y.(*ir.EncapsedStringPart)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.EqualExpr:
		y, ok := y.(*ir.EqualExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.ErrorSuppressExpr:
		y, ok := y.(*ir.ErrorSuppressExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.EvalExpr:
		y, ok := y.(*ir.EvalExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.ExitExpr:
		y, ok := y.(*ir.ExitExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if x.Die != y.Die {
			return false
		}
		return true
	case *ir.ExpressionStmt:
		y, ok := y.(*ir.ExpressionStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.FinallyStmt:
		y, ok := y.(*ir.FinallyStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.ForStmt:
		y, ok := y.(*ir.ForStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Init, y.Init) {
			return false
		}
		if !NodeSliceEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeSliceEqual(x.Loop, y.Loop) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.ForeachStmt:
		y, ok := y.(*ir.ForeachStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if !NodeEqual(x.Key, y.Key) {
			return false
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.FunctionCallExpr:
		y, ok := y.(*ir.FunctionCallExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Function, y.Function) {
			return false
		}
		if !NodeSliceEqual(x.Args, y.Args) {
			return false
		}
		return true
	case *ir.FunctionStmt:
		y, ok := y.(*ir.FunctionStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.FunctionName, y.FunctionName) {
			return false
		}
		if !NodeSliceEqual(x.Params, y.Params) {
			return false
		}
		if !NodeEqual(x.ReturnType, y.ReturnType) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		if x.ReturnsRef != y.ReturnsRef {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.GlobalStmt:
		y, ok := y.(*ir.GlobalStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Vars, y.Vars) {
			return false
		}
		return true
	case *ir.GotoStmt:
		y, ok := y.(*ir.GotoStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Label, y.Label) {
			return false
		}
		return true
	case *ir.GreaterExpr:
		y, ok := y.(*ir.GreaterExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.GreaterOrEqualExpr:
		y, ok := y.(*ir.GreaterOrEqualExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.GroupUseStmt:
		y, ok := y.(*ir.GroupUseStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.UseType, y.UseType) {
			return false
		}
		if !NodeEqual(x.Prefix, y.Prefix) {
			return false
		}
		if !NodeSliceEqual(x.UseList, y.UseList) {
			return false
		}
		return true
	case *ir.HaltCompilerStmt:
		y, ok := y.(*ir.HaltCompilerStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		return true
	case *ir.Heredoc:
		y, ok := y.(*ir.Heredoc)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Label != y.Label {
			return false
		}
		if !NodeSliceEqual(x.Parts, y.Parts) {
			return false
		}
		return true
	case *ir.IdenticalExpr:
		y, ok := y.(*ir.IdenticalExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.Identifier:
		y, ok := y.(*ir.Identifier)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.IfStmt:
		y, ok := y.(*ir.IfStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if !NodeSliceEqual(x.ElseIf, y.ElseIf) {
			return false
		}
		if !NodeEqual(x.Else, y.Else) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.ImportExpr:
		y, ok := y.(*ir.ImportExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Func != y.Func {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.InlineHTMLStmt:
		y, ok := y.(*ir.InlineHTMLStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.InstanceOfExpr:
		y, ok := y.(*ir.InstanceOfExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if !NodeEqual(x.Class, y.Class) {
			return false
		}
		return true
	case *ir.InterfaceExtendsStmt:
		y, ok := y.(*ir.InterfaceExtendsStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.InterfaceNames, y.InterfaceNames) {
			return false
		}
		return true
	case *ir.InterfaceStmt:
		y, ok := y.(*ir.InterfaceStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		if !NodeEqual(x.InterfaceName, y.InterfaceName) {
			return false
		}
		if !NodeEqual(x.Extends, y.Extends) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.IssetExpr:
		y, ok := y.(*ir.IssetExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Variables, y.Variables) {
			return false
		}
		return true
	case *ir.LabelStmt:
		y, ok := y.(*ir.LabelStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.LabelName, y.LabelName) {
			return false
		}
		return true
	case *ir.ListExpr:
		y, ok := y.(*ir.ListExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Items) != len(y.Items) {
			return false
		}
		for i := range x.Items {
			if !NodeEqual(x.Items[i], y.Items[i]) {
				return false
			}
		}
		if x.ShortSyntax != y.ShortSyntax {
			return false
		}
		return true
	case *ir.Lnumber:
		y, ok := y.(*ir.Lnumber)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.LogicalAndExpr:
		y, ok := y.(*ir.LogicalAndExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.LogicalOrExpr:
		y, ok := y.(*ir.LogicalOrExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.LogicalXorExpr:
		y, ok := y.(*ir.LogicalXorExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.MagicConstant:
		y, ok := y.(*ir.MagicConstant)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.MethodCallExpr:
		y, ok := y.(*ir.MethodCallExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Method, y.Method) {
			return false
		}
		if !NodeSliceEqual(x.Args, y.Args) {
			return false
		}
		return true
	case *ir.MinusExpr:
		y, ok := y.(*ir.MinusExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.ModExpr:
		y, ok := y.(*ir.ModExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.MulExpr:
		y, ok := y.(*ir.MulExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.Name:
		y, ok := y.(*ir.Name)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		return true
	case *ir.NamespaceStmt:
		y, ok := y.(*ir.NamespaceStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.NamespaceName, y.NamespaceName) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.NewExpr:
		y, ok := y.(*ir.NewExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Class, y.Class) {
			return false
		}
		if !NodeSliceEqual(x.Args, y.Args) {
			return false
		}
		return true
	case *ir.NopStmt:
		y, ok := y.(*ir.NopStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		return true
	case *ir.NotEqualExpr:
		y, ok := y.(*ir.NotEqualExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.NotIdenticalExpr:
		y, ok := y.(*ir.NotIdenticalExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.Nullable:
		y, ok := y.(*ir.Nullable)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.Parameter:
		y, ok := y.(*ir.Parameter)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.VariableType, y.VariableType) {
			return false
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.DefaultValue, y.DefaultValue) {
			return false
		}
		if x.ByRef != y.ByRef {
			return false
		}
		if x.Variadic != y.Variadic {
			return false
		}
		return true
	case *ir.ParenExpr:
		y, ok := y.(*ir.ParenExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.PlusExpr:
		y, ok := y.(*ir.PlusExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.PostDecExpr:
		y, ok := y.(*ir.PostDecExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		return true
	case *ir.PostIncExpr:
		y, ok := y.(*ir.PostIncExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		return true
	case *ir.PowExpr:
		y, ok := y.(*ir.PowExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.PreDecExpr:
		y, ok := y.(*ir.PreDecExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		return true
	case *ir.PreIncExpr:
		y, ok := y.(*ir.PreIncExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		return true
	case *ir.PrintExpr:
		y, ok := y.(*ir.PrintExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.PropertyFetchExpr:
		y, ok := y.(*ir.PropertyFetchExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Property, y.Property) {
			return false
		}
		return true
	case *ir.PropertyListStmt:
		y, ok := y.(*ir.PropertyListStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if len(x.Modifiers) != len(y.Modifiers) {
			return false
		}
		for i := range x.Modifiers {
			if !NodeEqual(x.Modifiers[i], y.Modifiers[i]) {
				return false
			}
		}
		if !NodeEqual(x.Type, y.Type) {
			return false
		}
		if !NodeSliceEqual(x.Properties, y.Properties) {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.PropertyStmt:
		y, ok := y.(*ir.PropertyStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.ReferenceExpr:
		y, ok := y.(*ir.ReferenceExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		return true
	case *ir.ReturnStmt:
		y, ok := y.(*ir.ReturnStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.Root:
		y, ok := y.(*ir.Root)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.ShellExecExpr:
		y, ok := y.(*ir.ShellExecExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Parts, y.Parts) {
			return false
		}
		return true
	case *ir.ShiftLeftExpr:
		y, ok := y.(*ir.ShiftLeftExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.ShiftRightExpr:
		y, ok := y.(*ir.ShiftRightExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.SimpleVar:
		y, ok := y.(*ir.SimpleVar)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Name != y.Name {
			return false
		}
		return true
	case *ir.SmallerExpr:
		y, ok := y.(*ir.SmallerExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.SmallerOrEqualExpr:
		y, ok := y.(*ir.SmallerOrEqualExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.SpaceshipExpr:
		y, ok := y.(*ir.SpaceshipExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Left, y.Left) {
			return false
		}
		if !NodeEqual(x.Right, y.Right) {
			return false
		}
		return true
	case *ir.StaticCallExpr:
		y, ok := y.(*ir.StaticCallExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Class, y.Class) {
			return false
		}
		if !NodeEqual(x.Call, y.Call) {
			return false
		}
		if !NodeSliceEqual(x.Args, y.Args) {
			return false
		}
		return true
	case *ir.StaticPropertyFetchExpr:
		y, ok := y.(*ir.StaticPropertyFetchExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Class, y.Class) {
			return false
		}
		if !NodeEqual(x.Property, y.Property) {
			return false
		}
		return true
	case *ir.StaticStmt:
		y, ok := y.(*ir.StaticStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Vars, y.Vars) {
			return false
		}
		return true
	case *ir.StaticVarStmt:
		y, ok := y.(*ir.StaticVarStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Variable, y.Variable) {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.StmtList:
		y, ok := y.(*ir.StmtList)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		return true
	case *ir.String:
		y, ok := y.(*ir.String)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Value != y.Value {
			return false
		}
		if x.DoubleQuotes != y.DoubleQuotes {
			return false
		}
		return true
	case *ir.SwitchStmt:
		y, ok := y.(*ir.SwitchStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeSliceEqual(x.Cases, y.Cases) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.TernaryExpr:
		y, ok := y.(*ir.TernaryExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Condition, y.Condition) {
			return false
		}
		if !NodeEqual(x.IfTrue, y.IfTrue) {
			return false
		}
		if !NodeEqual(x.IfFalse, y.IfFalse) {
			return false
		}
		return true
	case *ir.ThrowStmt:
		y, ok := y.(*ir.ThrowStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.TraitAdaptationListStmt:
		y, ok := y.(*ir.TraitAdaptationListStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Adaptations, y.Adaptations) {
			return false
		}
		return true
	case *ir.TraitMethodRefStmt:
		y, ok := y.(*ir.TraitMethodRefStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Trait, y.Trait) {
			return false
		}
		if !NodeEqual(x.Method, y.Method) {
			return false
		}
		return true
	case *ir.TraitStmt:
		y, ok := y.(*ir.TraitStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.TraitName, y.TraitName) {
			return false
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		if x.PhpDocComment != y.PhpDocComment {
			return false
		}
		return true
	case *ir.TraitUseAliasStmt:
		y, ok := y.(*ir.TraitUseAliasStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Ref, y.Ref) {
			return false
		}
		if !NodeEqual(x.Modifier, y.Modifier) {
			return false
		}
		if !NodeEqual(x.Alias, y.Alias) {
			return false
		}
		return true
	case *ir.TraitUsePrecedenceStmt:
		y, ok := y.(*ir.TraitUsePrecedenceStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Ref, y.Ref) {
			return false
		}
		if !NodeSliceEqual(x.Insteadof, y.Insteadof) {
			return false
		}
		return true
	case *ir.TraitUseStmt:
		y, ok := y.(*ir.TraitUseStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Traits, y.Traits) {
			return false
		}
		if !NodeEqual(x.TraitAdaptationList, y.TraitAdaptationList) {
			return false
		}
		return true
	case *ir.TryStmt:
		y, ok := y.(*ir.TryStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Stmts, y.Stmts) {
			return false
		}
		if !NodeSliceEqual(x.Catches, y.Catches) {
			return false
		}
		if !NodeEqual(x.Finally, y.Finally) {
			return false
		}
		return true
	case *ir.TypeCastExpr:
		y, ok := y.(*ir.TypeCastExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if x.Type != y.Type {
			return false
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.UnaryMinusExpr:
		y, ok := y.(*ir.UnaryMinusExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.UnaryPlusExpr:
		y, ok := y.(*ir.UnaryPlusExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.UnsetCastExpr:
		y, ok := y.(*ir.UnsetCastExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.UnsetStmt:
		y, ok := y.(*ir.UnsetStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeSliceEqual(x.Vars, y.Vars) {
			return false
		}
		return true
	case *ir.UseListStmt:
		y, ok := y.(*ir.UseListStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.UseType, y.UseType) {
			return false
		}
		if !NodeSliceEqual(x.Uses, y.Uses) {
			return false
		}
		return true
	case *ir.UseStmt:
		y, ok := y.(*ir.UseStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.UseType, y.UseType) {
			return false
		}
		if !NodeEqual(x.Use, y.Use) {
			return false
		}
		if !NodeEqual(x.Alias, y.Alias) {
			return false
		}
		return true
	case *ir.Var:
		y, ok := y.(*ir.Var)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	case *ir.WhileStmt:
		y, ok := y.(*ir.WhileStmt)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Cond, y.Cond) {
			return false
		}
		if !NodeEqual(x.Stmt, y.Stmt) {
			return false
		}
		if x.AltSyntax != y.AltSyntax {
			return false
		}
		return true
	case *ir.YieldExpr:
		y, ok := y.(*ir.YieldExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Key, y.Key) {
			return false
		}
		if !NodeEqual(x.Value, y.Value) {
			return false
		}
		return true
	case *ir.YieldFromExpr:
		y, ok := y.(*ir.YieldFromExpr)
		if !ok || x == nil || y == nil {
			return x == y
		}
		if !NodeEqual(x.Expr, y.Expr) {
			return false
		}
		return true
	default:
		panic(fmt.Sprintf(`unhandled type %T`, x))
	}
}
