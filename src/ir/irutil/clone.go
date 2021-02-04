// Code generated by the `ir/codegen` package. DO NOT EDIT.
package irutil

import (
	"fmt"
	"github.com/VKCOM/noverify/src/ir"
)

func NodeClone(x ir.Node) ir.Node {
	if x == nil {
		return nil
	}
	switch x := x.(type) {
	case *ir.AnonClassExpr:
		clone := *x
		clone.Args = NodeSliceClone(x.Args)
		clone.Class = classClone(x.Class)
		return &clone
	case *ir.Argument:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ArrayDimFetchExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Dim != nil {
			clone.Dim = NodeClone(x.Dim)
		}
		return &clone
	case *ir.ArrayExpr:
		clone := *x
		{
			sliceClone := make([]*ir.ArrayItemExpr, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*ir.ArrayItemExpr)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *ir.ArrayItemExpr:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key)
		}
		if x.Val != nil {
			clone.Val = NodeClone(x.Val)
		}
		return &clone
	case *ir.ArrowFunctionExpr:
		clone := *x
		clone.Params = NodeSliceClone(x.Params)
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.Assign:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignBitwiseAnd:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignBitwiseOr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignBitwiseXor:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignCoalesce:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignConcat:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignDiv:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignMinus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignMod:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignMul:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignPlus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignPow:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignReference:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignShiftLeft:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.AssignShiftRight:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression)
		}
		return &clone
	case *ir.BadString:
		clone := *x
		return &clone
	case *ir.BitwiseAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.BitwiseNotExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.BitwiseOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.BitwiseXorExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.BooleanAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.BooleanNotExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.BooleanOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.BreakStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.CaseStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.CatchStmt:
		clone := *x
		clone.Types = NodeSliceClone(x.Types)
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.ClassConstFetchExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class)
		}
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*ir.Identifier)
		}
		return &clone
	case *ir.ClassConstListStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		clone.Consts = NodeSliceClone(x.Consts)
		return &clone
	case *ir.ClassExtendsStmt:
		clone := *x
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(*ir.Name)
		}
		return &clone
	case *ir.ClassImplementsStmt:
		clone := *x
		clone.InterfaceNames = NodeSliceClone(x.InterfaceNames)
		return &clone
	case *ir.ClassMethodStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.MethodName != nil {
			clone.MethodName = NodeClone(x.MethodName).(*ir.Identifier)
		}
		clone.Params = NodeSliceClone(x.Params)
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.ClassStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(*ir.Identifier)
		}
		clone.Class = classClone(x.Class)
		return &clone
	case *ir.CloneExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ClosureExpr:
		clone := *x
		clone.Params = NodeSliceClone(x.Params)
		if x.ClosureUse != nil {
			clone.ClosureUse = NodeClone(x.ClosureUse).(*ir.ClosureUseExpr)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.ClosureUseExpr:
		clone := *x
		clone.Uses = NodeSliceClone(x.Uses)
		return &clone
	case *ir.CoalesceExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.ConcatExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.ConstFetchExpr:
		clone := *x
		if x.Constant != nil {
			clone.Constant = NodeClone(x.Constant).(*ir.Name)
		}
		return &clone
	case *ir.ConstListStmt:
		clone := *x
		clone.Consts = NodeSliceClone(x.Consts)
		return &clone
	case *ir.ConstantStmt:
		clone := *x
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*ir.Identifier)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ContinueStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.DeclareStmt:
		clone := *x
		clone.Consts = NodeSliceClone(x.Consts)
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.DefaultStmt:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.DivExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.Dnumber:
		clone := *x
		return &clone
	case *ir.DoStmt:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		return &clone
	case *ir.EchoStmt:
		clone := *x
		clone.Exprs = NodeSliceClone(x.Exprs)
		return &clone
	case *ir.ElseIfStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.ElseStmt:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.EmptyExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.Encapsed:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *ir.EncapsedStringPart:
		clone := *x
		return &clone
	case *ir.EqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.ErrorSuppressExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.EvalExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ExitExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ExpressionStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.FinallyStmt:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.ForStmt:
		clone := *x
		clone.Init = NodeSliceClone(x.Init)
		clone.Cond = NodeSliceClone(x.Cond)
		clone.Loop = NodeSliceClone(x.Loop)
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.ForeachStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		if x.Key != nil {
			clone.Key = NodeClone(x.Key)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.FunctionCallExpr:
		clone := *x
		if x.Function != nil {
			clone.Function = NodeClone(x.Function)
		}
		clone.Args = NodeSliceClone(x.Args)
		return &clone
	case *ir.FunctionStmt:
		clone := *x
		if x.FunctionName != nil {
			clone.FunctionName = NodeClone(x.FunctionName).(*ir.Identifier)
		}
		clone.Params = NodeSliceClone(x.Params)
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.GlobalStmt:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *ir.GotoStmt:
		clone := *x
		if x.Label != nil {
			clone.Label = NodeClone(x.Label).(*ir.Identifier)
		}
		return &clone
	case *ir.GreaterExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.GreaterOrEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.GroupUseStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(*ir.Identifier)
		}
		if x.Prefix != nil {
			clone.Prefix = NodeClone(x.Prefix).(*ir.Name)
		}
		clone.UseList = NodeSliceClone(x.UseList)
		return &clone
	case *ir.HaltCompilerStmt:
		clone := *x
		return &clone
	case *ir.Heredoc:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *ir.IdenticalExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.Identifier:
		clone := *x
		return &clone
	case *ir.IfStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		clone.ElseIf = NodeSliceClone(x.ElseIf)
		if x.Else != nil {
			clone.Else = NodeClone(x.Else)
		}
		return &clone
	case *ir.ImportExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.InlineHTMLStmt:
		clone := *x
		return &clone
	case *ir.InstanceOfExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		if x.Class != nil {
			clone.Class = NodeClone(x.Class)
		}
		return &clone
	case *ir.InterfaceExtendsStmt:
		clone := *x
		clone.InterfaceNames = NodeSliceClone(x.InterfaceNames)
		return &clone
	case *ir.InterfaceStmt:
		clone := *x
		if x.InterfaceName != nil {
			clone.InterfaceName = NodeClone(x.InterfaceName).(*ir.Identifier)
		}
		if x.Extends != nil {
			clone.Extends = NodeClone(x.Extends).(*ir.InterfaceExtendsStmt)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.IssetExpr:
		clone := *x
		clone.Variables = NodeSliceClone(x.Variables)
		return &clone
	case *ir.LabelStmt:
		clone := *x
		if x.LabelName != nil {
			clone.LabelName = NodeClone(x.LabelName).(*ir.Identifier)
		}
		return &clone
	case *ir.ListExpr:
		clone := *x
		{
			sliceClone := make([]*ir.ArrayItemExpr, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*ir.ArrayItemExpr)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *ir.Lnumber:
		clone := *x
		return &clone
	case *ir.LogicalAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.LogicalOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.LogicalXorExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.MagicConstant:
		clone := *x
		return &clone
	case *ir.MethodCallExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method)
		}
		clone.Args = NodeSliceClone(x.Args)
		return &clone
	case *ir.MinusExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.ModExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.MulExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.Name:
		clone := *x
		return &clone
	case *ir.NamespaceStmt:
		clone := *x
		if x.NamespaceName != nil {
			clone.NamespaceName = NodeClone(x.NamespaceName).(*ir.Name)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.NewExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class)
		}
		clone.Args = NodeSliceClone(x.Args)
		return &clone
	case *ir.NopStmt:
		clone := *x
		return &clone
	case *ir.NotEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.NotIdenticalExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.Nullable:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.Parameter:
		clone := *x
		if x.VariableType != nil {
			clone.VariableType = NodeClone(x.VariableType)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.DefaultValue != nil {
			clone.DefaultValue = NodeClone(x.DefaultValue)
		}
		return &clone
	case *ir.ParenExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.PlusExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.PostDecExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		return &clone
	case *ir.PostIncExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		return &clone
	case *ir.PowExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.PreDecExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		return &clone
	case *ir.PreIncExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		return &clone
	case *ir.PrintExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.PropertyFetchExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property)
		}
		return &clone
	case *ir.PropertyListStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.Type != nil {
			clone.Type = NodeClone(x.Type)
		}
		clone.Properties = NodeSliceClone(x.Properties)
		return &clone
	case *ir.PropertyStmt:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.ReferenceExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable)
		}
		return &clone
	case *ir.ReturnStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.Root:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.ShellExecExpr:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *ir.ShiftLeftExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.ShiftRightExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.SimpleVar:
		clone := *x
		return &clone
	case *ir.SmallerExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.SmallerOrEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.SpaceshipExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right)
		}
		return &clone
	case *ir.StaticCallExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class)
		}
		if x.Call != nil {
			clone.Call = NodeClone(x.Call)
		}
		clone.Args = NodeSliceClone(x.Args)
		return &clone
	case *ir.StaticPropertyFetchExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property)
		}
		return &clone
	case *ir.StaticStmt:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *ir.StaticVarStmt:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.StmtList:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.String:
		clone := *x
		return &clone
	case *ir.SwitchStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		if x.Cases != nil {
			clone.Cases = NodeSliceClone(x.Cases)
		}
		return &clone
	case *ir.TernaryExpr:
		clone := *x
		if x.Condition != nil {
			clone.Condition = NodeClone(x.Condition)
		}
		if x.IfTrue != nil {
			clone.IfTrue = NodeClone(x.IfTrue)
		}
		if x.IfFalse != nil {
			clone.IfFalse = NodeClone(x.IfFalse)
		}
		return &clone
	case *ir.ThrowStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.TraitAdaptationListStmt:
		clone := *x
		clone.Adaptations = NodeSliceClone(x.Adaptations)
		return &clone
	case *ir.TraitMethodRefStmt:
		clone := *x
		if x.Trait != nil {
			clone.Trait = NodeClone(x.Trait)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method).(*ir.Identifier)
		}
		return &clone
	case *ir.TraitStmt:
		clone := *x
		if x.TraitName != nil {
			clone.TraitName = NodeClone(x.TraitName).(*ir.Identifier)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *ir.TraitUseAliasStmt:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref)
		}
		if x.Modifier != nil {
			clone.Modifier = NodeClone(x.Modifier)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*ir.Identifier)
		}
		return &clone
	case *ir.TraitUsePrecedenceStmt:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref)
		}
		clone.Insteadof = NodeSliceClone(x.Insteadof)
		return &clone
	case *ir.TraitUseStmt:
		clone := *x
		clone.Traits = NodeSliceClone(x.Traits)
		if x.TraitAdaptationList != nil {
			clone.TraitAdaptationList = NodeClone(x.TraitAdaptationList)
		}
		return &clone
	case *ir.TryStmt:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		clone.Catches = NodeSliceClone(x.Catches)
		if x.Finally != nil {
			clone.Finally = NodeClone(x.Finally)
		}
		return &clone
	case *ir.TypeCastExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.UnaryMinusExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.UnaryPlusExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.UnsetCastExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.UnsetStmt:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *ir.UseListStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(*ir.Identifier)
		}
		clone.Uses = NodeSliceClone(x.Uses)
		return &clone
	case *ir.UseStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(*ir.Identifier)
		}
		if x.Use != nil {
			clone.Use = NodeClone(x.Use).(*ir.Name)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*ir.Identifier)
		}
		return &clone
	case *ir.Var:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	case *ir.WhileStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt)
		}
		return &clone
	case *ir.YieldExpr:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key)
		}
		if x.Value != nil {
			clone.Value = NodeClone(x.Value)
		}
		return &clone
	case *ir.YieldFromExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr)
		}
		return &clone
	}
	panic(fmt.Sprintf(`unhandled type %T`, x))
}
