package parser

import (
	"encoding/json"
)

type JsonParser struct{}

// Parse parse config from []byte
// cfg is required to be a ptr of struct
func (p *JsonParser) Parse(data []byte, cfg interface{}) error {
	return json.Unmarshal(data, cfg)
}
