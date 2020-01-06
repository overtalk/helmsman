package config

type Parser interface {
	Parse(data []byte, cfg interface{}) error
}
