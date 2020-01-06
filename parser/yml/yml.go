package parser

import (
	"gopkg.in/yaml.v2"
)

type YmlParser struct{}

// Parse parse config from []byte
// cfg is required to be a ptr of struct
func (p *YmlParser) Parse(data []byte, cfg interface{}) error {
	return yaml.Unmarshal(data, cfg)
}
