package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeg2Rad(t *testing.T) {
	assert.InEpsilon(t, 1.0471, Deg2Rad(60.0), 0.0001)
	assert.InEpsilon(t, 1.5707, Deg2Rad(90.0), 0.0001)
	assert.InEpsilon(t, 3.1415, Deg2Rad(180.0), 0.0001)
}

func TestRad2Deg(t *testing.T) {
	assert.InEpsilon(t, 60.0, Rad2Deg(1.0471), 0.0001)
	assert.InEpsilon(t, 90.0, Rad2Deg(1.5707), 0.0001)
	assert.InEpsilon(t, 180.0, Rad2Deg(3.1415), 0.0001)
}
