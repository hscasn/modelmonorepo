package config

import (
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/envprops"
)

// Config is the top-level configuration struct
type Config struct {
	Name string
	Port int
}

// New will recover the environment settings and parse them into a struct
func New() Config {
	return Config{
		Name: envprops.String("MODULE_NAME"),
		Port: int(envprops.Int("PORT")),
	}
}
