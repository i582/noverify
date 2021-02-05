package irutil

import (
	"github.com/z7zmey/php-parser/pkg/token"

	"github.com/VKCOM/noverify/src/ir"
)

func handleToken(t *token.Token, cb func(*token.Token) bool) bool {
	if !cb(t) {
		return false
	}

	needReturn := true
	for _, ff := range t.FreeFloating {
		needReturn = needReturn && handleToken(ff, cb)
	}

	return needReturn
}

func DeepIterateTokens(n ir.Node, cb func(*token.Token) bool) {
	n.IterateTokens(func(t *token.Token) bool {
		return handleToken(t, cb)
	})
}

func FindPhpDoc(n ir.Node) (string, bool) {
	var doc string

	Inspect(n, func(n ir.Node) bool {
		DeepIterateTokens(n, func(t *token.Token) bool {
			if t.ID == token.T_DOC_COMMENT {
				doc = string(t.Value)
				return false
			}

			return doc == ""
		})

		return doc == ""
	})

	return doc, doc != ""
}
