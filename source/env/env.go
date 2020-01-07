package env

import (
	"helmsman"
	"github.com/caarlos0/env"
)

type EnvConfig struct{}

func (e *EnvConfig) ParseConfig(parser config.Parser, cfg interface{}) error {
	return env.Parse(cfg)
}
