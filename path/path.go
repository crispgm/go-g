package path

import (
	"errors"
	"fmt"
	"os"
)

var errNoGoPath = errors.New("GOPATH is not set")

// GetGoPath returns $GOPATH
func GetGoPath() string {
	return os.Getenv("GOPATH")
}

// GetRepoPath returns repo path
func GetRepoPath(repoName string) (string, error) {
	gopath := GetGoPath()
	if len(gopath) == 0 {
		return "", errNoGoPath
	}
	path := fmt.Sprintf("%s/src/%s", gopath, repoName)
	return path, nil
}
