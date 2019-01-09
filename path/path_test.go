package path

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoPath(t *testing.T) {
	assert.Equal(t, os.Getenv("GOPATH"), GetGoPath())
}

func TestGetRepoPath(t *testing.T) {
	repoPath, err := GetRepoPath("a")
	assert.NoError(t, err)
	assert.Equal(t, os.Getenv("GOPATH")+"/src/a", repoPath)
}
