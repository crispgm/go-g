package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicDelim(t *testing.T) {
	tok := NewTokenizer("new js monokai", DefaultDelimiters, " ")
	assert.Equal(t, []string{"new", "js", "monokai"}, tok.Parse())
}

func TestBasicDelimWithOtherChars(t *testing.T) {
	tok := NewTokenizer("new js\t monokai\t  go", DefaultDelimiters, " ")
	assert.Equal(t, []string{"new", "js", "monokai", "go"}, tok.Parse())
}

func TestCustomDelim(t *testing.T) {
	tok := NewTokenizer("newsep  go sep fff", DefaultDelimiters, "sep")
	assert.Equal(t, []string{"new", "go", "fff"}, tok.Parse())
}

func TestCustomDelimWithOtherChars(t *testing.T) {
	tok := NewTokenizer("newsep  go  w w sep ffw", "w ", "sep")
	assert.Equal(t, []string{"ne", "go", "ff"}, tok.Parse())
}
