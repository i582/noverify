package php7_test

import (
	"io/ioutil"
	"testing"

	"github.com/VKCOM/noverify/src/php/parser/internal/php7"
	"github.com/VKCOM/noverify/src/php/parser/internal/scanner"
	"github.com/VKCOM/noverify/src/php/parser/pkg/cfg"
	"github.com/VKCOM/noverify/src/php/parser/pkg/version"
)

func BenchmarkPhp7(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")

	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := cfg.Config{
			Version: &version.Version{
				Major: 7,
				Minor: 4,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php7parser := php7.NewParser(lexer, config)
		php7parser.Parse()
	}
}
