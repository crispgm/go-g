package ip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalIP(t *testing.T) {
	assert.NotEmpty(t, GetLocalIP())
}
