package config

import (
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/env"
)

// Config is the top-level configuration struct
type Config struct {
	Name string
}

// New will recover the environment settings and parse them into a struct
func New() Config {
	return Config{
		Name: env.String("MODULE_NAME"),
	}
}
