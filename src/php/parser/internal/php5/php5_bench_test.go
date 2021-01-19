package php5_test

import (
	"io/ioutil"
	"testing"

	"github.com/VKCOM/noverify/src/php/parser/internal/php5"
	"github.com/VKCOM/noverify/src/php/parser/internal/scanner"
	"github.com/VKCOM/noverify/src/php/parser/pkg/cfg"
	"github.com/VKCOM/noverify/src/php/parser/pkg/version"
)

func BenchmarkPhp5(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")
	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := cfg.Config{
			Version: &version.Version{
				Major: 5,
				Minor: 6,
			},
		}
		lexer := scanner.NewLexer(src, config)
		php5parser := php5.NewParser(lexer, config)
		php5parser.Parse()
	}
}
