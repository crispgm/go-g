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

func TestIsFunc(t *testing.T) {
	env := NewDefaultEnv()
	if env.IsStaging() != false {
		t.Error("IsStaging test failed")
	}
	if env.IsDevelopment() != false {
		t.Error("IsDevelopment test failed")
	}
	if env.IsProduction() != false {
		t.Error("IsProduction test failed")
	}
}
