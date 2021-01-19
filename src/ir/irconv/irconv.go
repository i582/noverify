package irconv

import (
	"bytes"
	"fmt"

	"github.com/VKCOM/noverify/src/ir"
	"github.com/VKCOM/noverify/src/ir/irutil"
	"github.com/VKCOM/noverify/src/php/parser/pkg/ast"
	"github.com/VKCOM/noverify/src/phpdoc"
)

func ConvertNode(n ast.Vertex) ir.Node {
	c := NewConverter(phpdoc.NewTypeParser())
	return c.ConvertNode(n)
}

type Converter struct {
	namespace string

	phpdocTypeParser *phpdoc.TypeParser
}

// NewConverter returns a new AST->IR converter.
//
// If typeParser is nil, it will not eagerly try to parse phpdoc
// strings into phpdoc.CommentPart.
//
// It's intended to be re-used inside a signle thread context.
func NewConverter(typeParser *phpdoc.TypeParser) *Converter {
	return &Converter{
		phpdocTypeParser: typeParser,
	}
}

func (c *Converter) ConvertRoot(n *ast.Root) *ir.Root {
	return c.ConvertNode(n).(*ir.Root)
}

func (c *Converter) ConvertNode(n ast.Vertex) ir.Node {
	c.reset()
	return c.convNode(n)
}

func (c *Converter) reset() {
	c.namespace = ""
}

func (c *Converter) convNodeSlice(xs []ast.Vertex) []ir.Node {
	out := make([]ir.Node, len(xs))
	for i, x := range xs {
		out[i] = c.convNode(x)
	}
	return out
}

func (c *Converter) convNode(n ast.Vertex) ir.Node {
	if n == nil {
		return nil
	}
	switch n := n.(type) {
	case *ast.ExprAssign:
		if n == nil {
			return (*ir.Assign)(nil)
		}
		out := &ir.Assign{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignBitwiseAnd:
		if n == nil {
			return (*ir.AssignBitwiseAnd)(nil)
		}
		out := &ir.AssignBitwiseAnd{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignBitwiseOr:
		if n == nil {
			return (*ir.AssignBitwiseOr)(nil)
		}
		out := &ir.AssignBitwiseOr{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignBitwiseXor:
		if n == nil {
			return (*ir.AssignBitwiseXor)(nil)
		}
		out := &ir.AssignBitwiseXor{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignCoalesce:
		if n == nil {
			return (*ir.AssignCoalesce)(nil)
		}
		out := &ir.AssignCoalesce{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignConcat:
		if n == nil {
			return (*ir.AssignConcat)(nil)
		}
		out := &ir.AssignConcat{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignDiv:
		if n == nil {
			return (*ir.AssignDiv)(nil)
		}
		out := &ir.AssignDiv{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignMinus:
		if n == nil {
			return (*ir.AssignMinus)(nil)
		}
		out := &ir.AssignMinus{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignMod:
		if n == nil {
			return (*ir.AssignMod)(nil)
		}
		out := &ir.AssignMod{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignMul:
		if n == nil {
			return (*ir.AssignMul)(nil)
		}
		out := &ir.AssignMul{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignPlus:
		if n == nil {
			return (*ir.AssignPlus)(nil)
		}
		out := &ir.AssignPlus{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignPow:
		if n == nil {
			return (*ir.AssignPow)(nil)
		}
		out := &ir.AssignPow{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignReference:
		if n == nil {
			return (*ir.AssignReference)(nil)
		}
		out := &ir.AssignReference{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignShiftLeft:
		if n == nil {
			return (*ir.AssignShiftLeft)(nil)
		}
		out := &ir.AssignShiftLeft{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprAssignShiftRight:
		if n == nil {
			return (*ir.AssignShiftRight)(nil)
		}
		out := &ir.AssignShiftRight{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Expression = c.convNode(n.Expr)
		return out

	case *ast.ExprBinaryBitwiseAnd:
		if n == nil {
			return (*ir.BitwiseAndExpr)(nil)
		}
		out := &ir.BitwiseAndExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryBitwiseOr:
		if n == nil {
			return (*ir.BitwiseOrExpr)(nil)
		}
		out := &ir.BitwiseOrExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryBitwiseXor:
		if n == nil {
			return (*ir.BitwiseXorExpr)(nil)
		}
		out := &ir.BitwiseXorExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryBooleanAnd:
		if n == nil {
			return (*ir.BooleanAndExpr)(nil)
		}
		out := &ir.BooleanAndExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryBooleanOr:
		if n == nil {
			return (*ir.BooleanOrExpr)(nil)
		}
		out := &ir.BooleanOrExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryCoalesce:
		if n == nil {
			return (*ir.CoalesceExpr)(nil)
		}
		out := &ir.CoalesceExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryConcat:
		if n == nil {
			return (*ir.ConcatExpr)(nil)
		}
		out := &ir.ConcatExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryDiv:
		if n == nil {
			return (*ir.DivExpr)(nil)
		}
		out := &ir.DivExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryEqual:
		if n == nil {
			return (*ir.EqualExpr)(nil)
		}
		out := &ir.EqualExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryGreater:
		if n == nil {
			return (*ir.GreaterExpr)(nil)
		}
		out := &ir.GreaterExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryGreaterOrEqual:
		if n == nil {
			return (*ir.GreaterOrEqualExpr)(nil)
		}
		out := &ir.GreaterOrEqualExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryIdentical:
		if n == nil {
			return (*ir.IdenticalExpr)(nil)
		}
		out := &ir.IdenticalExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryLogicalAnd:
		if n == nil {
			return (*ir.LogicalAndExpr)(nil)
		}
		out := &ir.LogicalAndExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryLogicalOr:
		if n == nil {
			return (*ir.LogicalOrExpr)(nil)
		}
		out := &ir.LogicalOrExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryLogicalXor:
		if n == nil {
			return (*ir.LogicalXorExpr)(nil)
		}
		out := &ir.LogicalXorExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryMinus:
		if n == nil {
			return (*ir.MinusExpr)(nil)
		}
		out := &ir.MinusExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryMod:
		if n == nil {
			return (*ir.ModExpr)(nil)
		}
		out := &ir.ModExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryMul:
		if n == nil {
			return (*ir.MulExpr)(nil)
		}
		out := &ir.MulExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryNotEqual:
		if n == nil {
			return (*ir.NotEqualExpr)(nil)
		}
		out := &ir.NotEqualExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryNotIdentical:
		if n == nil {
			return (*ir.NotIdenticalExpr)(nil)
		}
		out := &ir.NotIdenticalExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryPlus:
		if n == nil {
			return (*ir.PlusExpr)(nil)
		}
		out := &ir.PlusExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryPow:
		if n == nil {
			return (*ir.PowExpr)(nil)
		}
		out := &ir.PowExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryShiftLeft:
		if n == nil {
			return (*ir.ShiftLeftExpr)(nil)
		}
		out := &ir.ShiftLeftExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinaryShiftRight:
		if n == nil {
			return (*ir.ShiftRightExpr)(nil)
		}
		out := &ir.ShiftRightExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinarySmaller:
		if n == nil {
			return (*ir.SmallerExpr)(nil)
		}
		out := &ir.SmallerExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinarySmallerOrEqual:
		if n == nil {
			return (*ir.SmallerOrEqualExpr)(nil)
		}
		out := &ir.SmallerOrEqualExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprBinarySpaceship:
		if n == nil {
			return (*ir.SpaceshipExpr)(nil)
		}
		out := &ir.SpaceshipExpr{}
		out.OpTkn = n.OpTkn
		out.Position = n.Position
		out.Left = c.convNode(n.Left)
		out.Right = c.convNode(n.Right)
		return out

	case *ast.ExprCastArray:
		return c.convCastExpr(n, n.Expr, "array")
	case *ast.ExprCastBool:
		return c.convCastExpr(n, n.Expr, "bool")
	case *ast.ExprCastInt:
		return c.convCastExpr(n, n.Expr, "int")
	case *ast.ExprCastDouble:
		return c.convCastExpr(n, n.Expr, "float")
	case *ast.ExprCastObject:
		return c.convCastExpr(n, n.Expr, "object")
	case *ast.ExprCastString:
		return c.convCastExpr(n, n.Expr, "string")

	case *ast.ExprCastUnset:
		// We dont convert (unset)$x into CastExpr deliberately.
		if n == nil {
			return (*ir.UnsetCastExpr)(nil)
		}
		out := &ir.UnsetCastExpr{}
		out.CastTkn = n.CastTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprArray:
		if n == nil {
			return (*ir.ArrayExpr)(nil)
		}
		out := &ir.ArrayExpr{}
		out.ArrayTkn = n.ArrayTkn
		out.OpenBracketTkn = n.OpenBracketTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseBracketTkn = n.CloseBracketTkn

		out.Position = n.Position
		{
			slice := make([]*ir.ArrayItemExpr, len(n.Items))
			for i := range n.Items {
				slice[i] = c.convNode(n.Items[i]).(*ir.ArrayItemExpr)
			}
			out.Items = slice
		}
		out.ShortSyntax = bytes.Contains(n.ArrayTkn.Value, []byte("..."))
		return out

	case *ast.ExprArrayDimFetch:
		if n == nil {
			return (*ir.ArrayDimFetchExpr)(nil)
		}
		out := &ir.ArrayDimFetchExpr{}
		out.OpenBracketTkn = n.OpenBracketTkn
		out.CloseBracketTkn = n.CloseBracketTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Dim = c.convNode(n.Dim)
		return out

	case *ast.ExprArrayItem:
		if n == nil {
			return (*ir.ArrayItemExpr)(nil)
		}
		out := &ir.ArrayItemExpr{}
		out.EllipsisTkn = n.EllipsisTkn
		out.DoubleArrowTkn = n.DoubleArrowTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.Position = n.Position
		out.Key = c.convNode(n.Key)
		out.Val = c.convNode(n.Val)
		out.Unpack = bytes.Contains(n.EllipsisTkn.Value, []byte("..."))
		return out

	case *ast.ExprArrowFunction:
		if n == nil {
			return (*ir.ArrowFunctionExpr)(nil)
		}
		out := &ir.ArrowFunctionExpr{}

		out.StaticTkn = n.StaticTkn
		out.FnTkn = n.FnTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn

		out.Position = n.Position

		out.ReturnsRef = bytes.ContainsRune(n.AmpersandTkn.Value, '&')
		out.Static = bytes.Contains(n.StaticTkn.Value, []byte("static"))

		// out.PhpDocComment = n.PhpDocComment
		// out.PhpDoc = c.parsePHPDoc(n.PhpDocComment)

		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn

		out.Params = c.convNodeSlice(n.Params)
		out.ReturnType = c.convNode(n.ReturnType)

		out.DoubleArrowTkn = n.DoubleArrowTkn

		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprBitwiseNot:
		if n == nil {
			return (*ir.BitwiseNotExpr)(nil)
		}
		out := &ir.BitwiseNotExpr{}
		out.TildaTkn = n.TildaTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprBooleanNot:
		if n == nil {
			return (*ir.BooleanNotExpr)(nil)
		}
		out := &ir.BooleanNotExpr{}
		out.ExclamationTkn = n.ExclamationTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprClassConstFetch:
		if n == nil {
			return (*ir.ClassConstFetchExpr)(nil)
		}
		out := &ir.ClassConstFetchExpr{}
		out.DoubleColonTkn = n.DoubleColonTkn
		out.Position = n.Position
		out.Class = c.convNode(n.Class)
		out.ConstantName = c.convNode(n.Const).(*ir.Identifier)
		return out

	case *ast.ExprClone:
		if n == nil {
			return (*ir.CloneExpr)(nil)
		}
		out := &ir.CloneExpr{}
		out.CloneTkn = n.CloneTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprClosure:
		if n == nil {
			return (*ir.ClosureExpr)(nil)
		}
		out := &ir.ClosureExpr{}

		out.StaticTkn = n.StaticTkn
		out.FunctionTkn = n.FunctionTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn

		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.UseTkn = n.UseTkn
		out.UseOpenParenthesisTkn = n.UseOpenParenthesisTkn

		out.UseSeparatorTkns = n.UseSeparatorTkns
		out.UseCloseParenthesisTkn = n.UseCloseParenthesisTkn
		out.ColonTkn = n.ColonTkn

		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn

		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn

		out.Position = n.Position

		out.ReturnsRef = bytes.ContainsRune(n.AmpersandTkn.Value, '&')
		out.Static = bytes.Contains(n.StaticTkn.Value, []byte("static"))

		// out.PhpDocComment = n.PhpDocComment
		// out.PhpDoc = c.parsePHPDoc(n.PhpDocComment)

		out.Params = c.convNodeSlice(n.Params)

		// out.ClosureUse = c.convNodeSlice(n.Uses).(*ir.ClosureUseExpr)
		out.ClosureUse = &ir.ClosureUseExpr{
			Token:    nil,
			Position: nil,
			Uses:     c.convNodeSlice(n.Uses),
		}

		out.ReturnType = c.convNode(n.ReturnType)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	// case *ast.ExprClosureUse:
	// 	if n == nil {
	// 		return (*ir.ClosureUseExpr)(nil)
	// 	}
	// 	out := &ir.ClosureUseExpr{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Uses = c.convNodeSlice(n.Uses)
	// 	return out

	case *ast.ExprConstFetch:
		if n == nil {
			return (*ir.ConstFetchExpr)(nil)
		}
		out := &ir.ConstFetchExpr{}
		out.Position = n.Position
		out.Constant = c.convNode(n.Const).(*ir.Name)
		return out

	case *ast.ExprEmpty:
		if n == nil {
			return (*ir.EmptyExpr)(nil)
		}
		out := &ir.EmptyExpr{}
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn

		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprErrorSuppress:
		if n == nil {
			return (*ir.ErrorSuppressExpr)(nil)
		}
		out := &ir.ErrorSuppressExpr{}
		out.AtTkn = n.AtTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprEval:
		if n == nil {
			return (*ir.EvalExpr)(nil)
		}
		out := &ir.EvalExpr{}
		out.EvalTkn = n.EvalTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprExit:
		if n == nil {
			return (*ir.ExitExpr)(nil)
		}
		out := &ir.ExitExpr{}
		out.ExitTkn = n.ExitTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.Die = bytes.Contains(n.ExitTkn.Value, []byte("die"))
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprFunctionCall:
		if n == nil {
			return (*ir.FunctionCallExpr)(nil)
		}
		out := &ir.FunctionCallExpr{}
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.Position = n.Position
		out.Function = c.convNode(n.Function)
		out.Args = c.convNodeSlice(n.Args)
		return out

	case *ast.ExprInstanceOf:
		if n == nil {
			return (*ir.InstanceOfExpr)(nil)
		}
		out := &ir.InstanceOfExpr{}
		out.InstanceOfTkn = n.InstanceOfTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		out.Class = c.convNode(n.Class)
		return out

	case *ast.ExprIsset:
		if n == nil {
			return (*ir.IssetExpr)(nil)
		}
		out := &ir.IssetExpr{}
		out.IssetTkn = n.IssetTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.Position = n.Position
		out.Variables = c.convNodeSlice(n.Vars)
		return out

	case *ast.ExprList:
		if n == nil {
			return (*ir.ListExpr)(nil)
		}
		out := &ir.ListExpr{}

		out.ListTkn = n.ListTkn
		out.OpenBracketTkn = n.OpenBracketTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseBracketTkn = n.CloseBracketTkn

		out.Position = n.Position
		{
			slice := make([]*ir.ArrayItemExpr, len(n.Items))
			for i := range n.Items {
				slice[i] = c.convNode(n.Items[i]).(*ir.ArrayItemExpr)
			}
			out.Items = slice
		}
		out.ShortSyntax = !bytes.Contains(n.ListTkn.Value, []byte("list"))
		return out

	case *ast.ExprMethodCall:
		if n == nil {
			return (*ir.MethodCallExpr)(nil)
		}
		out := &ir.MethodCallExpr{}
		out.ObjectOperatorTkn = n.ObjectOperatorTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn

		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Method = c.convNode(n.Method)
		out.Args = c.convNodeSlice(n.Args)
		return out

	case *ast.ExprNew:
		if n == nil {
			return (*ir.NewExpr)(nil)
		}
		out := &ir.NewExpr{}

		out.NewTkn = n.NewTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn

		out.Position = n.Position
		out.Class = c.convNode(n.Class)

		if n.Args != nil {
			out.Args = c.convNodeSlice(n.Args)
		}
		return out

	case *ast.ExprBrackets:
		if n == nil {
			return (*ir.ParenExpr)(nil)
		}
		out := &ir.ParenExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprPostDec:
		if n == nil {
			return (*ir.PostDecExpr)(nil)
		}
		out := &ir.PostDecExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		return out

	case *ast.ExprPostInc:
		if n == nil {
			return (*ir.PostIncExpr)(nil)
		}
		out := &ir.PostIncExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		return out

	case *ast.ExprPreDec:
		if n == nil {
			return (*ir.PreDecExpr)(nil)
		}
		out := &ir.PreDecExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		return out

	case *ast.ExprPreInc:
		if n == nil {
			return (*ir.PreIncExpr)(nil)
		}
		out := &ir.PreIncExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		return out

	case *ast.ExprPrint:
		if n == nil {
			return (*ir.PrintExpr)(nil)
		}
		out := &ir.PrintExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprPropertyFetch:
		if n == nil {
			return (*ir.PropertyFetchExpr)(nil)
		}
		out := &ir.PropertyFetchExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variable = c.convNode(n.Var)
		out.Property = c.convNode(n.Property)
		return out

	// case *ast.ExprReference:
	// 	if n == nil {
	// 		return (*ir.ReferenceExpr)(nil)
	// 	}
	// 	out := &ir.ReferenceExpr{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Variable = c.convNode(n.Variable)
	// 	return out

	case *ast.ExprRequire:
		return c.convImportExpr(n, n.Expr, "require")
	case *ast.ExprRequireOnce:
		return c.convImportExpr(n, n.Expr, "require_once")
	case *ast.ExprInclude:
		return c.convImportExpr(n, n.Expr, "include")
	case *ast.ExprIncludeOnce:
		return c.convImportExpr(n, n.Expr, "include_once")

	case *ast.ExprShellExec:
		if n == nil {
			return (*ir.ShellExecExpr)(nil)
		}
		out := &ir.ShellExecExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Parts = c.convNodeSlice(n.Parts)
		return out

	case *ast.ExprStaticCall:
		if n == nil {
			return (*ir.StaticCallExpr)(nil)
		}
		out := &ir.StaticCallExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Class = c.convNode(n.Class)
		out.Call = c.convNode(n.Call)
		out.ArgsToken = n.ArgumentList.Token
		out.Args = c.convNodeSlice(n.ArgumentList.Arguments)
		return out

	case *ast.ExprStaticPropertyFetch:
		if n == nil {
			return (*ir.StaticPropertyFetchExpr)(nil)
		}
		out := &ir.StaticPropertyFetchExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Class = c.convNode(n.Class)
		out.Property = c.convNode(n.Prop)
		return out

	case *ast.ExprTernary:
		if n == nil {
			return (*ir.TernaryExpr)(nil)
		}
		out := &ir.TernaryExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Condition = c.convNode(n.Cond)
		out.IfTrue = c.convNode(n.IfTrue)
		out.IfFalse = c.convNode(n.IfFalse)
		return out

	case *ast.ExprUnaryMinus:
		if n == nil {
			return (*ir.UnaryMinusExpr)(nil)
		}
		out := &ir.UnaryMinusExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprUnaryPlus:
		if n == nil {
			return (*ir.UnaryPlusExpr)(nil)
		}
		out := &ir.UnaryPlusExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.ExprYield:
		if n == nil {
			return (*ir.YieldExpr)(nil)
		}
		out := &ir.YieldExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Key = c.convNode(n.Key)
		out.Value = c.convNode(n.Val)
		return out

	case *ast.ExprYieldFrom:
		if n == nil {
			return (*ir.YieldFromExpr)(nil)
		}
		out := &ir.YieldFromExpr{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.NameFullyQualified:
		return &ir.Name{
			Token:    n.Token,
			Position: n.Position,
			Value:    fullyQualifiedToString(n),
		}
	case *ast.Name:
		return &ir.Name{
			Token:    n.Token,
			Position: n.Position,
			Value:    namePartsToString(n.Parts),
		}
	case *ast.NameRelative:
		return c.convRelativeName(n)

	case *ast.Argument:
		if n == nil {
			return (*ir.Argument)(nil)
		}
		out := &ir.Argument{}
		out.Token = n.Token
		out.Position = n.Position
		out.Variadic = n.VariadicTkn != nil
		out.IsReference = n.IsReference
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.Identifier:
		if n == nil {
			return (*ir.Identifier)(nil)
		}
		out := &ir.Identifier{}
		out.Token = n.Token
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.Nullable:
		if n == nil {
			return (*ir.Nullable)(nil)
		}
		out := &ir.Nullable{}
		out.Token = n.Token
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.Parameter:
		if n == nil {
			return (*ir.Parameter)(nil)
		}
		out := &ir.Parameter{}
		out.Token = n.Token
		out.Position = n.Position
		out.ByRef = n.ByRef
		out.Variadic = n.Variadic
		out.VariableType = c.convNode(n.VariableType)
		out.Variable = c.convNode(n.Variable).(*ir.SimpleVar)
		out.DefaultValue = c.convNode(n.DefaultValue)
		return out

	case *ast.Root:
		if n == nil {
			return (*ir.Root)(nil)
		}
		out := &ir.Root{}
		out.EndTkn = n.EndTkn
		out.Position = n.Position
		{
			slice := make([]ir.Node, len(n.Stmts))
			for i := range n.Stmts {
				slice[i] = c.convNode(n.Stmts[i])
			}
			out.Stmts = slice
		}
		return out

		// TODO:
	case *ast.ExprVariable:
		if n == nil {
			return (*ir.SimpleVar)(nil)
		}
		out := &ir.SimpleVar{}
		out.Token = n.Token
		out.Position = n.Position
		out.Name = n.Name
		return out

	// case *ast.Var:
	// 	if n == nil {
	// 		return (*ir.Var)(nil)
	// 	}
	// 	out := &ir.Var{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Expr = c.convNode(n.Expr)
	// 	return out

	case *ast.ScalarDnumber:
		if n == nil {
			return (*ir.Dnumber)(nil)
		}
		out := &ir.Dnumber{}
		out.NumberTkn = n.NumberTkn
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.ScalarEncapsed:
		if n == nil {
			return (*ir.Encapsed)(nil)
		}
		out := &ir.Encapsed{}
		out.OpenQuoteTkn = n.OpenQuoteTkn
		out.CloseQuoteTkn = n.CloseQuoteTkn
		out.Position = n.Position
		out.Parts = c.convNodeSlice(n.Parts)
		return out

	case *ast.ScalarEncapsedStringPart:
		if n == nil {
			return (*ir.EncapsedStringPart)(nil)
		}
		out := &ir.EncapsedStringPart{}
		out.EncapsedStrTkn = n.EncapsedStrTkn
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.ScalarHeredoc:
		if n == nil {
			return (*ir.Heredoc)(nil)
		}
		out := &ir.Heredoc{}
		out.OpenHeredocTkn = n.OpenHeredocTkn
		out.CloseHeredocTkn = n.CloseHeredocTkn
		out.Position = n.Position
		// out.Label = n.Label
		out.Parts = c.convNodeSlice(n.Parts)
		return out

	case *ast.ScalarLnumber:
		if n == nil {
			return (*ir.Lnumber)(nil)
		}
		out := &ir.Lnumber{}
		out.NumberTkn = n.NumberTkn
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.ScalarMagicConstant:
		if n == nil {
			return (*ir.MagicConstant)(nil)
		}
		out := &ir.MagicConstant{}
		out.MagicConstTkn = n.MagicConstTkn
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.ScalarString:
		return convString(n)

	case *ast.StmtBreak:
		if n == nil {
			return (*ir.BreakStmt)(nil)
		}
		out := &ir.BreakStmt{}
		out.BreakTkn = n.BreakTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtCase:
		if n == nil {
			return (*ir.CaseStmt)(nil)
		}
		out := &ir.CaseStmt{}
		out.CaseTkn = n.CaseTkn
		out.CaseSeparatorTkn = n.CaseSeparatorTkn
		out.Position = n.Position
		out.Cond = c.convNode(n.Cond)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	// case *ast.StmtCaseList:
	// 	if n == nil {
	// 		return (*ir.CaseListStmt)(nil)
	// 	}
	// 	out := &ir.CaseListStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Cases = c.convNodeSlice(n.Cases)
	// 	return out

	case *ast.StmtCatch:
		if n == nil {
			return (*ir.CatchStmt)(nil)
		}
		out := &ir.CatchStmt{}

		out.CatchTkn = n.CatchTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn

		out.Position = n.Position
		out.Types = c.convNodeSlice(n.Types)
		out.Variable = c.convNode(n.Var).(*ir.SimpleVar)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtClass:
		if n == nil {
			return (*ir.ClassStmt)(nil)
		}
		return c.convClass(n)

	case *ast.StmtClassConstList:
		if n == nil {
			return (*ir.ClassConstListStmt)(nil)
		}
		out := &ir.ClassConstListStmt{}
		out.ConstTkn = n.ConstTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position

		{
			slice := make([]*ir.Identifier, len(n.Modifiers))
			for i := range n.Modifiers {
				slice[i] = c.convNode(n.Modifiers[i]).(*ir.Identifier)
			}
			out.Modifiers = slice
		}
		out.Consts = c.convNodeSlice(n.Consts)
		return out

	// case *ast.StmtClassExtends:
	// 	if n == nil {
	// 		return (*ir.ClassExtendsStmt)(nil)
	// 	}
	// 	out := &ir.ClassExtendsStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.ClassName = c.convNode(n.ClassName).(*ir.Name)
	// 	return out
	//
	// case *ast.StmtClassImplements:
	// 	if n == nil {
	// 		return (*ir.ClassImplementsStmt)(nil)
	// 	}
	// 	out := &ir.ClassImplementsStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.InterfaceNames = c.convNodeSlice(n.InterfaceNames)
	// 	return out

	case *ast.StmtClassMethod:
		if n == nil {
			return (*ir.ClassMethodStmt)(nil)
		}
		out := &ir.ClassMethodStmt{}
		out.FunctionTkn = n.FunctionTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn

		out.Position = n.Position
		out.ReturnsRef = n.AmpersandTkn != nil
		// out.PhpDocComment = n.PhpDocComment
		// out.PhpDoc = c.parsePHPDoc(n.PhpDocComment)
		out.MethodName = c.convNode(n.Name).(*ir.Identifier)
		{
			slice := make([]*ir.Identifier, len(n.Modifiers))
			for i := range n.Modifiers {
				slice[i] = c.convNode(n.Modifiers[i]).(*ir.Identifier)
			}
			out.Modifiers = slice
		}
		out.Params = c.convNodeSlice(n.Params)
		out.ReturnType = c.convNode(n.ReturnType)
		out.Stmt = c.convNode(n.Stmt)
		return out

	case *ast.StmtConstList:
		if n == nil {
			return (*ir.ConstListStmt)(nil)
		}
		out := &ir.ConstListStmt{}

		out.ConstTkn = n.ConstTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Consts = c.convNodeSlice(n.Consts)
		return out

	case *ast.StmtConstant:
		if n == nil {
			return (*ir.ConstantStmt)(nil)
		}
		out := &ir.ConstantStmt{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		// out.PhpDocComment = n.PhpDocComment
		out.ConstantName = c.convNode(n.Name).(*ir.Identifier)
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtContinue:
		if n == nil {
			return (*ir.ContinueStmt)(nil)
		}
		out := &ir.ContinueStmt{}
		out.ContinueTkn = n.ContinueTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtDeclare:
		if n == nil {
			return (*ir.DeclareStmt)(nil)
		}
		out := &ir.DeclareStmt{}
		out.DeclareTkn = n.DeclareTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.EndDeclareTkn = n.EndDeclareTkn

		out.Position = n.Position
		out.Consts = c.convNodeSlice(n.Consts)
		out.Stmt = c.convNode(n.Stmt)
		out.Alt = n.EndDeclareTkn != nil // TODO:
		return out

	case *ast.StmtDefault:
		if n == nil {
			return (*ir.DefaultStmt)(nil)
		}
		out := &ir.DefaultStmt{}
		out.DefaultTkn = n.DefaultTkn
		out.CaseSeparatorTkn = n.CaseSeparatorTkn
		out.Position = n.Position
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtDo:
		if n == nil {
			return (*ir.DoStmt)(nil)
		}
		out := &ir.DoStmt{}
		out.DoTkn = n.DoTkn
		out.WhileTkn = n.WhileTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Stmt = c.convNode(n.Stmt)
		out.Cond = c.convNode(n.Cond)
		return out

	case *ast.StmtEcho:
		if n == nil {
			return (*ir.EchoStmt)(nil)
		}
		out := &ir.EchoStmt{}
		out.EchoTkn = n.EchoTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Exprs = c.convNodeSlice(n.Exprs)
		return out

	case *ast.StmtElse:
		if n == nil {
			return (*ir.ElseStmt)(nil)
		}
		out := &ir.ElseStmt{}
		out.ElseTkn = n.ElseTkn
		out.ColonTkn = n.ColonTkn
		out.Position = n.Position
		out.Stmt = c.convNode(n.Stmt)
		out.AltSyntax = n.ColonTkn != nil
		return out

	case *ast.StmtElseIf:
		if n == nil {
			return (*ir.ElseIfStmt)(nil)
		}
		out := &ir.ElseIfStmt{}
		out.ElseIfTkn = n.ElseIfTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn

		out.Position = n.Position
		out.Cond = c.convNode(n.Cond)
		out.Stmt = c.convNode(n.Stmt)
		out.AltSyntax = bytes.ContainsRune(n.CloseParenthesisTkn.Value, ':')
		out.Merged = bytes.Contains(n.ElseIfTkn.Value, []byte("elseif"))
		return out

	case *ast.StmtExpression:
		if n == nil {
			return (*ir.ExpressionStmt)(nil)
		}
		out := &ir.ExpressionStmt{}
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtFinally:
		if n == nil {
			return (*ir.FinallyStmt)(nil)
		}
		out := &ir.FinallyStmt{}
		out.FinallyTkn = n.FinallyTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn

		out.Position = n.Position
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtFor:
		if n == nil {
			return (*ir.ForStmt)(nil)
		}
		out := &ir.ForStmt{}
		out.ForTkn = n.ForTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.InitSeparatorTkns = n.InitSeparatorTkns
		out.InitSemiColonTkn = n.InitSemiColonTkn
		out.CondSeparatorTkns = n.CondSeparatorTkns
		out.CondSemiColonTkn = n.CondSemiColonTkn
		out.LoopSeparatorTkns = n.LoopSeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.EndForTkn = n.EndForTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Init = c.convNodeSlice(n.Init)
		out.Cond = c.convNodeSlice(n.Cond)
		out.Loop = c.convNodeSlice(n.Loop)
		out.Stmt = c.convNode(n.Stmt)
		out.AltSyntax = bytes.Contains(n.EndForTkn.Value, []byte("endfor"))
		return out

	case *ast.StmtForeach:
		if n == nil {
			return (*ir.ForeachStmt)(nil)
		}
		out := &ir.ForeachStmt{}
		out.ForeachTkn = n.ForeachTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.AsTkn = n.AsTkn
		out.DoubleArrowTkn = n.DoubleArrowTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.EndForeachTkn = n.EndForeachTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		out.Key = c.convNode(n.Key)
		out.Variable = c.convNode(n.Var)
		out.Stmt = c.convNode(n.Stmt)
		out.AltSyntax = bytes.Contains(n.EndForeachTkn.Value, []byte("endforeach"))
		return out

	case *ast.StmtFunction:
		if n == nil {
			return (*ir.FunctionStmt)(nil)
		}
		out := &ir.FunctionStmt{}
		out.FunctionTkn = n.FunctionTkn
		out.AmpersandTkn = n.AmpersandTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn

		out.Position = n.Position
		out.ReturnsRef = n.AmpersandTkn != nil
		// out.PhpDocComment = n.PhpDocComment
		// out.PhpDoc = c.parsePHPDoc(n.PhpDocComment)

		out.FunctionName = c.convNode(n.Name).(*ir.Identifier)
		out.Params = c.convNodeSlice(n.Params)
		out.ReturnType = c.convNode(n.ReturnType)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtGlobal:
		if n == nil {
			return (*ir.GlobalStmt)(nil)
		}
		out := &ir.GlobalStmt{}
		out.GlobalTkn = n.GlobalTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Vars = c.convNodeSlice(n.Vars)
		return out

	case *ast.StmtGoto:
		if n == nil {
			return (*ir.GotoStmt)(nil)
		}
		out := &ir.GotoStmt{}
		out.GotoTkn = n.GotoTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Label = c.convNode(n.Label).(*ir.Identifier)
		return out

	case *ast.StmtGroupUseList:
		if n == nil {
			return (*ir.GroupUseStmt)(nil)
		}
		out := &ir.GroupUseStmt{}
		out.UseTkn = n.UseTkn
		out.LeadingNsSeparatorTkn = n.LeadingNsSeparatorTkn
		out.NsSeparatorTkn = n.NsSeparatorTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		useType := c.convNode(n.Type)
		if useType != nil {
			out.UseType = useType.(*ir.Identifier)
		}
		out.Prefix = c.convNode(n.Prefix).(*ir.Name)
		out.UseList = c.convNodeSlice(n.Uses)
		return out

	case *ast.StmtHaltCompiler:
		if n == nil {
			return (*ir.HaltCompilerStmt)(nil)
		}
		out := &ir.HaltCompilerStmt{}
		out.HaltCompilerTkn = n.HaltCompilerTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		return out

	case *ast.StmtIf:
		if n == nil {
			return (*ir.IfStmt)(nil)
		}
		out := &ir.IfStmt{}
		out.IfTkn = n.IfTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.EndIfTkn = n.EndIfTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Cond = c.convNode(n.Cond)
		out.Stmt = c.convNode(n.Stmt)
		out.ElseIf = c.convNodeSlice(n.ElseIf)
		out.Else = c.convNode(n.Else)
		out.AltSyntax = n.ColonTkn != nil
		return out

	case *ast.StmtInlineHtml:
		if n == nil {
			return (*ir.InlineHtmlStmt)(nil)
		}
		out := &ir.InlineHtmlStmt{}
		out.InlineHtmlTkn = n.InlineHtmlTkn
		out.Position = n.Position
		out.Value = string(n.Value)
		return out

	case *ast.StmtInterface:
		if n == nil {
			return (*ir.InterfaceStmt)(nil)
		}
		out := &ir.InterfaceStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.PhpDocComment = n.PhpDocComment
		out.InterfaceName = c.convNode(n.InterfaceName).(*ir.Identifier)
		out.Extends = c.convNode(n.Extends).(*ir.InterfaceExtendsStmt)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	// case *ast.StmtInterfaceExtends:
	// 	if n == nil {
	// 		return (*ir.InterfaceExtendsStmt)(nil)
	// 	}
	// 	out := &ir.InterfaceExtendsStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.InterfaceNames = c.convNodeSlice(n.InterfaceNames)
	// 	return out

	case *ast.StmtLabel:
		if n == nil {
			return (*ir.LabelStmt)(nil)
		}
		out := &ir.LabelStmt{}
		out.ColonTkn = n.ColonTkn
		out.Position = n.Position
		out.LabelName = c.convNode(n.Name).(*ir.Identifier)
		return out

	case *ast.StmtNamespace:
		if n == nil {
			return (*ir.NamespaceStmt)(nil)
		}
		out := &ir.NamespaceStmt{}
		out.Token = n.Token
		out.Position = n.Position
		if n.NamespaceName != nil {
			out.NamespaceName = c.convNode(n.NamespaceName).(*ir.Name)
			c.namespace = out.NamespaceName.Value
		}
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtNop:
		if n == nil {
			return (*ir.NopStmt)(nil)
		}
		out := &ir.NopStmt{}
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		return out

	case *ast.StmtProperty:
		if n == nil {
			return (*ir.PropertyStmt)(nil)
		}
		out := &ir.PropertyStmt{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position

		// out.PhpDocComment = n.PhpDocComment
		// out.PhpDoc = c.parsePHPDoc(n.PhpDocComment)

		out.Variable = c.convNode(n.Var).(*ir.SimpleVar)
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtPropertyList:
		if n == nil {
			return (*ir.PropertyListStmt)(nil)
		}
		out := &ir.PropertyListStmt{}
		out.SeparatorTkns = n.SeparatorTkns
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		{
			slice := make([]*ir.Identifier, len(n.Modifiers))
			for i := range n.Modifiers {
				slice[i] = c.convNode(n.Modifiers[i]).(*ir.Identifier)
			}
			out.Modifiers = slice
		}
		out.Type = c.convNode(n.Type)
		out.Properties = c.convNodeSlice(n.Props)
		return out

	case *ast.StmtReturn:
		if n == nil {
			return (*ir.ReturnStmt)(nil)
		}
		out := &ir.ReturnStmt{}
		out.ReturnTkn = n.ReturnTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtStatic:
		if n == nil {
			return (*ir.StaticStmt)(nil)
		}
		out := &ir.StaticStmt{}
		out.StaticTkn = n.StaticTkn
		out.SeparatorTkns = n.SeparatorTkns
		out.Position = n.Position
		out.Vars = c.convNodeSlice(n.Vars)
		return out

	case *ast.StmtStaticVar:
		if n == nil {
			return (*ir.StaticVarStmt)(nil)
		}
		out := &ir.StaticVarStmt{}
		out.EqualTkn = n.EqualTkn
		out.Position = n.Position
		out.Variable = c.convNode(n.Var).(*ir.SimpleVar)
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtStmtList:
		if n == nil {
			return (*ir.StmtList)(nil)
		}
		out := &ir.StmtList{}
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn
		out.Position = n.Position
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	case *ast.StmtSwitch:
		if n == nil {
			return (*ir.SwitchStmt)(nil)
		}
		out := &ir.SwitchStmt{}
		out.SwitchTkn = n.SwitchTkn
		out.OpenParenthesisTkn = n.OpenParenthesisTkn
		out.CloseParenthesisTkn = n.CloseParenthesisTkn
		out.ColonTkn = n.ColonTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CaseSeparatorTkn = n.CaseSeparatorTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn
		out.EndSwitchTkn = n.EndSwitchTkn
		out.SemiColonTkn = n.SemiColonTkn

		out.Position = n.Position
		out.Cond = c.convNode(n.Cond)
		out.CaseList = c.convNodeSlice(n.Cases) // .(*ir.CaseListStmt)
		out.AltSyntax = n.ColonTkn != nil
		return out

	case *ast.StmtThrow:
		if n == nil {
			return (*ir.ThrowStmt)(nil)
		}
		out := &ir.ThrowStmt{}
		out.ThrowTkn = n.ThrowTkn
		out.SemiColonTkn = n.SemiColonTkn
		out.Position = n.Position
		out.Expr = c.convNode(n.Expr)
		return out

	case *ast.StmtTrait:
		if n == nil {
			return (*ir.TraitStmt)(nil)
		}
		out := &ir.TraitStmt{}
		out.TraitTkn = n.TraitTkn
		out.OpenCurlyBracketTkn = n.OpenCurlyBracketTkn
		out.CloseCurlyBracketTkn = n.CloseCurlyBracketTkn

		out.Position = n.Position
		// out.PhpDocComment = n.PhpDocComment
		out.TraitName = c.convNode(n.Name).(*ir.Identifier)
		out.Stmts = c.convNodeSlice(n.Stmts)
		return out

	// case *ast.StmtTraitAdaptationList:
	// 	if n == nil {
	// 		return (*ir.TraitAdaptationListStmt)(nil)
	// 	}
	// 	out := &ir.TraitAdaptationListStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Adaptations = c.convNodeSlice(n.Adaptations)
	// 	return out
	//
	// case *ast.StmtTraitMethodRef:
	// 	if n == nil {
	// 		return (*ir.TraitMethodRefStmt)(nil)
	// 	}
	// 	out := &ir.TraitMethodRefStmt{}
	// 	out.Token = n.Token
	// 	out.Position = n.Position
	// 	out.Trait = c.convNode(n.Trait)
	// 	out.Method = c.convNode(n.Method).(*ir.Identifier)
	// 	return out

	case *ast.StmtTraitUse:
		if n == nil {
			return (*ir.TraitUseStmt)(nil)
		}
		out := &ir.TraitUseStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Traits = c.convNodeSlice(n.Traits)
		out.TraitAdaptationList = c.convNode(n.TraitAdaptationList)
		return out

	case *ast.StmtTraitUseAlias:
		if n == nil {
			return (*ir.TraitUseAliasStmt)(nil)
		}
		out := &ir.TraitUseAliasStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Ref = c.convNode(n.Ref)
		out.Modifier = c.convNode(n.Modifier)
		out.Alias = c.convNode(n.Alias).(*ir.Identifier)
		return out

	case *ast.StmtTraitUsePrecedence:
		if n == nil {
			return (*ir.TraitUsePrecedenceStmt)(nil)
		}
		out := &ir.TraitUsePrecedenceStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Ref = c.convNode(n.Ref)
		out.Insteadof = c.convNodeSlice(n.Insteadof)
		return out

	case *ast.StmtTry:
		if n == nil {
			return (*ir.TryStmt)(nil)
		}
		out := &ir.TryStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Stmts = c.convNodeSlice(n.Stmts)
		out.Catches = c.convNodeSlice(n.Catches)
		out.Finally = c.convNode(n.Finally)
		return out

	case *ast.StmtUnset:
		if n == nil {
			return (*ir.UnsetStmt)(nil)
		}
		out := &ir.UnsetStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Vars = c.convNodeSlice(n.Vars)
		return out

	case *ast.StmtUse:
		if n == nil {
			return (*ir.UseStmt)(nil)
		}
		out := &ir.UseStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.UseType = c.convNode(n.UseType).(*ir.Identifier)
		out.Use = c.convNode(n.Use).(*ir.Name)
		out.Alias = c.convNode(n.Alias).(*ir.Identifier)
		return out

	case *ast.StmtUseList:
		if n == nil {
			return (*ir.UseListStmt)(nil)
		}
		out := &ir.UseListStmt{}
		out.Token = n.Token
		out.Position = n.Position
		useType := c.convNode(n.UseType)
		if useType != nil {
			out.UseType = useType.(*ir.Identifier)
		}
		out.Uses = c.convNodeSlice(n.Uses)
		return out

	case *ast.StmtWhile:
		if n == nil {
			return (*ir.WhileStmt)(nil)
		}
		out := &ir.WhileStmt{}
		out.Token = n.Token
		out.Position = n.Position
		out.Cond = c.convNode(n.Cond)
		out.Stmt = c.convNode(n.Stmt)
		out.AltSyntax = n.AltSyntax
		return out
	}

	panic(fmt.Sprintf("unhandled type %T", n))
}

func (c *Converter) convRelativeName(n *ast.NameRelative) *ir.Name {
	value := namePartsToString(n.Parts)
	if c.namespace != "" {
		value = `\` + c.namespace + `\` + value
	}
	return &ir.Name{
		Token:    n.Token,
		Position: n.Position,
		Value:    value,
	}
}

func (c *Converter) convImportExpr(n, e ast.Vertex, fn string) *ir.ImportExpr {
	return &ir.ImportExpr{
		Token:    *n.GetToken(),
		Position: n.GetPosition(),
		Func:     fn,
		Expr:     c.convNode(e),
	}
}

func (c *Converter) convCastExpr(n, e ast.Vertex, typ string) *ir.TypeCastExpr {
	return &ir.TypeCastExpr{
		Token:    *n.GetToken(),
		Position: n.GetPosition(),
		Type:     typ,
		Expr:     c.convNode(e),
	}
}

func (c *Converter) convClass(n *ast.StmtClass) ir.Node {
	class := ir.Class{
		PhpDocComment: n.PhpDocComment,
		PhpDoc:        c.parsePHPDoc(n.PhpDocComment),
		Extends:       c.convNode(n.Extends).(*ir.ClassExtendsStmt),
		Implements:    c.convNode(n.Implements).(*ir.ClassImplementsStmt),
		Stmts:         c.convNodeSlice(n.Stmts),
	}

	if n.ClassName == nil {
		// Anonymous class expression.
		out := &ir.AnonClassExpr{
			Token:    n.Token,
			Position: n.Position,
			Class:    class,
		}
		if n.ArgumentList != nil {
			out.ArgsToken = n.ArgumentList.Token
			out.Args = c.convNodeSlice(n.ArgumentList.Arguments)
		}
		return out
	}

	out := &ir.ClassStmt{
		Token:     n.Token,
		Position:  n.Position,
		Class:     class,
		ClassName: c.convNode(n.ClassName).(*ir.Identifier),
	}
	if n.Modifiers != nil {
		slice := make([]*ir.Identifier, len(n.Modifiers))
		for i := range n.Modifiers {
			slice[i] = c.convNode(n.Modifiers[i]).(*ir.Identifier)
		}
		out.Modifiers = slice
	}
	return out
}

func (c *Converter) parsePHPDoc(doc string) []phpdoc.CommentPart {
	if c.phpdocTypeParser != nil {
		return phpdoc.Parse(c.phpdocTypeParser, doc)
	}
	return nil
}

func convString(n *ast.ScalarString) ir.Node {
	out := &ir.String{
		Token:    n.Token,
		Position: n.Position,
	}

	// We can't use n.Value[0] as quote char directly as when
	// we parse string parts like $_SERVER[HTTP_HOST] we get
	// HTTP_HOST as a value with no quotes.
	var quote byte
	if n.Value[0] == '"' {
		quote = '"'
	} else {
		quote = '\''
	}

	out.DoubleQuotes = n.Value[0] == '"'
	unquoted := irutil.Unquote(n.Value)
	s, err := interpretString(unquoted, quote)
	if err != nil {
		return &ir.BadString{
			Token:        n.Token,
			Position:     n.Position,
			Value:        unquoted,
			Error:        err.Error(),
			DoubleQuotes: out.DoubleQuotes,
		}
	}
	out.Value = s

	return out
}
