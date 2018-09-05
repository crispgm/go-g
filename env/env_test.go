package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	env := NewEnv("RUNTIME")
	assert.Equal(t, "RUNTIME", env.EnvName)

}

func TestNewDefEnv(t *testing.T) {
	env := NewDefaultEnv()
	assert.Equal(t, DefaultEnv, env.EnvName)
}

func TestIsFunc(t *testing.T) {
	env := NewDefaultEnv()
	assert.False(t, env.IsStaging())
	assert.False(t, env.IsDevelopment())
	assert.False(t, env.IsProduction())
}
