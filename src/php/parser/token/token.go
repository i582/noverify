package token

import "github.com/VKCOM/noverify/src/php/parser/pkg/position"

//go:generate stringer -type=ID -output ./token_string.go
type ID int

const (
	T_INCLUDE ID = iota + 57346
	T_INCLUDE_ONCE
	T_EXIT
	T_IF
	T_LNUMBER
	T_DNUMBER
	T_STRING
	T_STRING_VARNAME
	T_VARIABLE
	T_NUM_STRING
	T_INLINE_HTML
	T_CHARACTER
	T_BAD_CHARACTER
	T_ENCAPSED_AND_WHITESPACE
	T_CONSTANT_ENCAPSED_STRING
	T_ECHO
	T_DO
	T_WHILE
	T_ENDWHILE
	T_FOR
	T_ENDFOR
	T_FOREACH
	T_ENDFOREACH
	T_DECLARE
	T_ENDDECLARE
	T_AS
	T_SWITCH
	T_ENDSWITCH
	T_CASE
	T_DEFAULT
	T_BREAK
	T_CONTINUE
	T_GOTO
	T_FUNCTION
	T_FN
	T_CONST
	T_RETURN
	T_TRY
	T_CATCH
	T_FINALLY
	T_THROW
	T_USE
	T_INSTEADOF
	T_GLOBAL
	T_VAR
	T_UNSET
	T_ISSET
	T_EMPTY
	T_HALT_COMPILER
	T_CLASS
	T_TRAIT
	T_INTERFACE
	T_EXTENDS
	T_IMPLEMENTS
	T_OBJECT_OPERATOR
	T_DOUBLE_ARROW
	T_LIST
	T_ARRAY
	T_CALLABLE
	T_CLASS_C
	T_TRAIT_C
	T_METHOD_C
	T_FUNC_C
	T_LINE
	T_FILE
	T_COMMENT
	T_DOC_COMMENT
	T_OPEN_TAG
	T_OPEN_TAG_WITH_ECHO
	T_CLOSE_TAG
	T_WHITESPACE
	T_START_HEREDOC
	T_END_HEREDOC
	T_DOLLAR_OPEN_CURLY_BRACES
	T_CURLY_OPEN
	T_PAAMAYIM_NEKUDOTAYIM
	T_NAMESPACE
	T_NS_C
	T_DIR
	T_NS_SEPARATOR
	T_ELLIPSIS
	T_EVAL
	T_REQUIRE
	T_REQUIRE_ONCE
	T_LOGICAL_OR
	T_LOGICAL_XOR
	T_LOGICAL_AND
	T_INSTANCEOF
	T_NEW
	T_CLONE
	T_ELSEIF
	T_ELSE
	T_ENDIF
	T_PRINT
	T_YIELD
	T_STATIC
	T_ABSTRACT
	T_FINAL
	T_PRIVATE
	T_PROTECTED
	T_PUBLIC
	T_INC
	T_DEC
	T_YIELD_FROM
	T_INT_CAST
	T_DOUBLE_CAST
	T_STRING_CAST
	T_ARRAY_CAST
	T_OBJECT_CAST
	T_BOOL_CAST
	T_UNSET_CAST
	T_COALESCE
	T_SPACESHIP
	T_NOELSE
	T_PLUS_EQUAL
	T_MINUS_EQUAL
	T_MUL_EQUAL
	T_POW_EQUAL
	T_DIV_EQUAL
	T_CONCAT_EQUAL
	T_MOD_EQUAL
	T_AND_EQUAL
	T_OR_EQUAL
	T_XOR_EQUAL
	T_SL_EQUAL
	T_SR_EQUAL
	T_COALESCE_EQUAL
	T_BOOLEAN_OR
	T_BOOLEAN_AND
	T_POW
	T_SL
	T_SR
	T_IS_IDENTICAL
	T_IS_NOT_IDENTICAL
	T_IS_EQUAL
	T_IS_NOT_EQUAL
	T_IS_SMALLER_OR_EQUAL
	T_IS_GREATER_OR_EQUAL
)

type Token struct {
	ID           ID
	Value        []byte
	Position     *position.Position
	FreeFloating []*Token
}

func (t *Token) GetPosition() *position.Position {
	return t.Position
}
