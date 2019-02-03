// Package tokenizer is designed for parsing user tokens from string
// It's like a command line argument parser but different.
// Usecase:
//   new    javascript \t monokai\n// comment\nconsole.log('console.log');
//   parsed as:
//     token: [new, javascript, monokai]
//     data: // comment\nconsole.log('console.log');
package tokenizer

import (
	"strings"
)

// Tokenizer definition
type Tokenizer struct {
	input      string
	delimiters string
	tokenSep   string
	tokens     []string
}

// DefaultDelimiters for most cases
const DefaultDelimiters = "\t "

// NewTokenizer creates a tokenizer
func NewTokenizer(input string, delims string, tokenSep string) *Tokenizer {
	return &Tokenizer{
		input:      input,
		delimiters: delims,
		tokenSep:   tokenSep,
	}
}

// Parse the input
func (t *Tokenizer) Parse() []string {
	splits := strings.Split(t.input, t.tokenSep)
	if len(splits) == 0 {
		return nil
	}

	for _, item := range splits {
		item = strings.Trim(item, t.delimiters)
		if item == "" {
			continue
		}
		t.tokens = append(t.tokens, item)
	}

	return t.tokens
}
