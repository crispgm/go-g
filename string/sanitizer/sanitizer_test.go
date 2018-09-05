package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSpace(t *testing.T) {
	w := RemoveDelimiters("Hello World")
	assert.Equal(t, "HelloWorld", w)
	w = RemoveDelimiters("宝马     3")
	assert.Equal(t, "宝马3", w)
	w = RemoveDelimiters("宝马 3 系")
	assert.Equal(t, "宝马3系", w)
}

func TestRemovePunctuation(t *testing.T) {
	w := RemoveDelimiters("奥迪·A3")
	assert.Equal(t, "奥迪A3", w)
	w = RemoveDelimiters("奥迪-A3好不好？")
	assert.Equal(t, "奥迪A3好不好", w)
}

func TestRemoveUnicodePunctuation(t *testing.T) {
	w := RemoveDelimiters("奥迪「A3」")
	assert.Equal(t, "奥迪A3", w)
	w = RemoveDelimiters("奥迪A3很≈不错")
	assert.Equal(t, "奥迪A3很不错", w)
}
