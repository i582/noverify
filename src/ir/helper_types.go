package ir

import (
	"github.com/VKCOM/noverify/src/phpdoc"
)

// Helper types are not real nodes, they're usually used
// to express some structure that is common between several nodes.
//
// In other words, structs defined in this file do not implement the Node interface.

// Class is a common shape between the ClassStmt and AnonClassExpr.
// It doesn't include positions info.
type Class struct {
	Extends    *ClassExtendsStmt
	Implements *ClassImplementsStmt
	Stmts      []Node

	Doc
}

type Doc struct {
	PhpDocComment string
	PhpDoc        []phpdoc.CommentPart
}
