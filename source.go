package config

type Source interface {
	ParseConfig(t ParserType, cfg interface{}) error
}
