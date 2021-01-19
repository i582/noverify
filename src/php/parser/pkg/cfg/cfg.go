package cfg

import (
	"github.com/VKCOM/noverify/src/php/parser/pkg/errors"
	"github.com/VKCOM/noverify/src/php/parser/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
