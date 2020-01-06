package config

type Source interface {
	ParseConfig(parser Parser, cfg interface{}) error
}
