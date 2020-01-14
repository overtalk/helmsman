package parser

import (
	"github.com/BurntSushi/toml"
)

type TomlParser struct{}

func (t *TomlParser) Parse(data []byte, cfg interface{}) error {
	return toml.Unmarshal(data, cfg)
}
