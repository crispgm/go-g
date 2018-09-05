package mac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMac(t *testing.T) {
	addr, err := GetMacAddr()
	assert.NoError(t, err)
	t.Log(addr)
}
