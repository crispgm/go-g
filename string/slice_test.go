package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	var slice = []string{
		"hello",
		"world",
	}
	assert.True(t, ExistsInSlice("hello", slice))
	assert.False(t, ExistsInSlice("olleh", slice))

	var emptySlice []string
	assert.False(t, ExistsInSlice("olleh", emptySlice))
}

func TestDedupSlice(t *testing.T) {
	var slice = []string{
		"hello",
		"world",
	}
	assert.Equal(t, slice, DedupSlice(slice))

	var sliceDup = []string{
		"hello",
		"world",
		"hello",
	}
	assert.Equal(t, slice, DedupSlice(sliceDup))

	var emptySlice []string
	assert.Equal(t, emptySlice, DedupSlice(emptySlice))
}
