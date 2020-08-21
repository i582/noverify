package irconv

import (
	"bytes"
	"testing"
)

func BenchmarkInterpretString(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "NoEscapes", input: `Hello, World!`},
		{name: "SingleEscape", input: `Hell?o, World!\n`},
		{name: "HarderInput", input: `\u{1f575} Hell\151, \x57orld! 💔`},
	}

	for _, test := range tests {
		b.Run(test.name+"Q1", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = interpretString(test.input, '\'')
			}
		})
		b.Run(test.name+"Q2", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = interpretString(test.input, '"')
			}
		})
	}
}

func TestInterpretString(t *testing.T) {
	tests := []struct {
		raw         string
		singleQuote []byte
		doubleQuote []byte
	}{
		// Empty strings evaluate to the same value.
		// (Basically check that we can handle this corner case.)
		{
			raw:         ``,
			singleQuote: []byte{},
			doubleQuote: []byte{},
		},

		// Simple strings evaluate to the same value.
		{
			raw:         `HTTP_HOST`,
			singleQuote: []byte("HTTP_HOST"),
			doubleQuote: []byte("HTTP_HOST"),
		},

		// Both literals treat \\ as literal \.
		{
			raw:         `\\`,
			singleQuote: []byte{'\\'},
			doubleQuote: []byte{'\\'},
		},

		// Single quote escape works in '-literals.
		{
			raw:         `\'`,
			singleQuote: []byte{'\''},
			doubleQuote: []byte{'\\', '\''},
		},

		// Double quote escape works in "-literals.
		{
			raw:         `\"`,
			singleQuote: []byte{'\\', '"'},
			doubleQuote: []byte{'"'},
		},

		// \n is interpreted only in "-literals.
		{
			raw:         `\n`,
			singleQuote: []byte{'\\', 'n'},
			doubleQuote: []byte{'\n'},
		},
		{
			raw:         `\r`,
			singleQuote: []byte{'\\', 'r'},
			doubleQuote: []byte{'\r'},
		},
		{
			raw:         `\t`,
			singleQuote: []byte{'\\', 't'},
			doubleQuote: []byte{'\t'},
		},
		{
			raw:         `\f`,
			singleQuote: []byte{'\\', 'f'},
			doubleQuote: []byte{'\f'},
		},
		{
			raw:         `\v`,
			singleQuote: []byte{'\\', 'v'},
			doubleQuote: []byte{'\v'},
		},
		{
			raw:         `\e`,
			singleQuote: []byte{'\\', 'e'},
			doubleQuote: []byte{0x1B},
		},

		// \$ is a way to suppress string interpolation in "-literals.
		// It's ignored in '-literals.
		{
			raw:         `\$`,
			singleQuote: []byte{'\\', '$'},
			doubleQuote: []byte{'$'},
		},

		// In "-literals, valid hex escapes evaluate to a single byte.
		{
			raw:         `\x1`,
			singleQuote: []byte(`\x1`),
			doubleQuote: []byte{0x1},
		},
		{
			raw:         `\x12`,
			singleQuote: []byte(`\x12`),
			doubleQuote: []byte{0x12},
		},
		{
			raw:         `\xFf`,
			singleQuote: []byte(`\xFf`),
			doubleQuote: []byte{0xff},
		},

		// In "-literals, invalid hex escapes are inserted literally.
		{
			raw:         `\xgg`,
			singleQuote: []byte(`\xgg`),
			doubleQuote: []byte(`\xgg`),
		},
		{
			raw:         `\xz`,
			singleQuote: []byte(`\xz`),
			doubleQuote: []byte(`\xz`),
		},
		{
			raw:         `\x`,
			singleQuote: []byte(`\x`),
			doubleQuote: []byte(`\x`),
		},

		// Octal literals can have up to 3 digits.
		{
			raw:         `\1`,
			singleQuote: []byte(`\1`),
			doubleQuote: []byte{1},
		},
		{
			raw:         `\12`,
			singleQuote: []byte(`\12`),
			doubleQuote: []byte{012},
		},
		{
			raw:         `\123`,
			singleQuote: []byte(`\123`),
			doubleQuote: []byte{0123},
		},
		{
			raw:         `\1234`,
			singleQuote: []byte(`\1234`),
			doubleQuote: []byte{0123, '4'},
		},
		{
			raw:         `\12\12`,
			singleQuote: []byte(`\12\12`),
			doubleQuote: []byte{012, 012},
		},
		{
			raw:         `\129`,
			singleQuote: []byte(`\129`),
			doubleQuote: []byte{012, '9'},
		},

		// Octal literals overflow to a byte (e.g. "\400" === "\000").
		{
			raw:         `\400`,
			singleQuote: []byte(`\400`),
			doubleQuote: []byte{0},
		},

		// Invalid octal sequences are interpreted literally.
		{
			raw:         `\8`,
			singleQuote: []byte(`\8`),
			doubleQuote: []byte(`\8`),
		},

		// Check that we can handle unicode correctly.
		{
			raw:         `💔`,
			singleQuote: []byte(`💔`),
			doubleQuote: []byte(`💔`),
		},

		// Valid u-sequences.
		{
			raw:         `\u{1f575}`,
			singleQuote: []byte(`\u{1f575}`),
			doubleQuote: []byte(`🕵`),
		},
		{
			raw:         `\u{1F575}`,
			singleQuote: []byte(`\u{1F575}`),
			doubleQuote: []byte(`🕵`),
		},
		{
			raw:         `\u{a}`,
			singleQuote: []byte(`\u{a}`),
			doubleQuote: []byte{10},
		},
		{
			raw:         `\u{aA}`,
			singleQuote: []byte(`\u{aA}`),
			doubleQuote: []byte(`ª`),
		},
		{
			raw:         `\u{aA}x\u{aA}`,
			singleQuote: []byte(`\u{aA}x\u{aA}`),
			doubleQuote: []byte(`ªxª`),
		},
		{
			raw:         `\u{aA1}`,
			singleQuote: []byte(`\u{aA1}`),
			doubleQuote: []byte(`ડ`),
		},
		{
			raw:         `\u{aA19}`,
			singleQuote: []byte(`\u{aA19}`),
			doubleQuote: []byte(`ꨙ`),
		},

		// Mixed tests.
		{
			raw:         `\nHello, \$username!`,
			singleQuote: []byte(`\nHello, \$username!`),
			doubleQuote: []byte("\nHello, $username!"),
		},
		{
			raw:         `~(?i)(?m)\bfoo|bar|^baz~`,
			singleQuote: []byte(`~(?i)(?m)\bfoo|bar|^baz~`),
			doubleQuote: []byte(`~(?i)(?m)\bfoo|bar|^baz~`),
		},

		// Invalid u-sequences.
		{
			raw:         `\u{zzzz}`,
			singleQuote: []byte(`\u{zzzz}`),
			doubleQuote: []byte("FAIL"),
		},
		{
			raw:         `\u{zz`,
			singleQuote: []byte(`\u{zz`),
			doubleQuote: []byte("FAIL"),
		},
		{
			raw:         `\u0`,
			singleQuote: []byte(`\u0`),
			doubleQuote: []byte(`\u0`),
		},
		{
			raw:         `\u0123`,
			singleQuote: []byte(`\u0123`),
			doubleQuote: []byte(`\u0123`),
		},

		// Unterminated escape sequence.
		{
			raw:         `a\`,
			singleQuote: []byte("FAIL"),
			doubleQuote: []byte("FAIL"),
		},
	}

	runTest := func(t *testing.T, raw string, quote byte, want []byte) {
		t.Helper()
		have, ok := interpretString(raw, quote)

		if bytes.Equal(want, []byte("FAIL")) {
			if ok {
				t.Errorf("expected %s interpretation to fail, got success", raw)
			}
			return
		}

		if !ok {
			t.Errorf("failed to interpret %s (quote=%c)", raw, quote)
			return
		}
		haveBytes := []byte(have)
		if !bytes.Equal(haveBytes, want) {
			t.Errorf("%s bytes mismatch (quote=%c):\nhave: %[3]v\nwant: %[4]v\nhave-string: %[3]s\nwant-string: %[4]s",
				raw, quote, haveBytes, want)
		}
	}

	for _, test := range tests {
		runTest(t, test.raw, '\'', test.singleQuote)
		runTest(t, test.raw, '"', test.doubleQuote)
	}
}