package env

import (
	"github.com/caarlos0/env"
)

type EnvConfig struct{}

func (e *EnvConfig) ParseConfig(cfg interface{}) error {
	return env.Parse(cfg)
}
