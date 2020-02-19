package config

import (
	jsonParser "github.com/overtalk/helmsman/parser/json"
	ymlParser "github.com/overtalk/helmsman/parser/yml"
)

type ParserType string

const (
	JsonFormat = "JSON"
	YamlFormat = "YAML"
)

func GetParser(t ParserType) Parser {
	switch t {
	case YamlFormat:
		return &ymlParser.YmlParser{}
	default:
		return &jsonParser.JsonParser{}
	}
}

type Parser interface {
	Parse(data []byte, cfg interface{}) error
}
