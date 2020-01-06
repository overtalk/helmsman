package parser

import (
	"encoding/xml"
)

type XmlParser struct{}

// Parse parse config from []byte
// cfg is required to be a ptr of struct
func (p *XmlParser) Parse(data []byte, cfg interface{}) error {
	return xml.Unmarshal(data, cfg)
}
