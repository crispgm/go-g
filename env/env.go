package env

import (
	"os"
)

// Env ...
type Env struct {
	EnvName string
}

// Const
const (
	DefaultEnv     = "GO_G_RUNTIME"
	ProductionEnv  = "prod"
	DevelopmentEnv = "dev"
	StaingEnv      = "stage"
)

// NewEnv creates Env
func NewEnv(envName string) *Env {
	return &Env{
		EnvName: envName,
	}
}

// NewDefaultEnv creates Env with default name
func NewDefaultEnv() *Env {
	return &Env{
		EnvName: DefaultEnv,
	}
}

// GetRuntimeEnv returns runtime env
func (env Env) GetRuntimeEnv() string {
	return os.Getenv(env.EnvName)
}

// IsProduction checks whether in production
func (env Env) IsProduction() bool {
	return env.GetRuntimeEnv() == ProductionEnv
}

// IsDevelopment checks whether in development
func (env Env) IsDevelopment() bool {
	return env.GetRuntimeEnv() == DevelopmentEnv
}

// IsStaging checks whether in staing
func (env Env) IsStaging() bool {
	return env.GetRuntimeEnv() == StaingEnv
}
