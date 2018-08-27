package env

import "testing"

func TestNewEnv(t *testing.T) {
	env := NewEnv("RUNTIME")
	if env.EnvName != "RUNTIME" {
		t.Error("NewEnv failed")
	}
}

func TestNewDefEnv(t *testing.T) {
	env := NewDefaultEnv()
	if env.EnvName != DefaultEnv {
		t.Error("NewDefaultEnv failed")
	}
}
