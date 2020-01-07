package parser_test

import (
	"testing"

	. "helmsman/parser/yml"
)

const str = `LogLevel: debug

ProtoDir: /Users/qinhan/go/src/oss/testdata/xml-parse/proto

LogDir: /Users/qinhan/data/log

TCP:
  Addr: 127.0.0.1
  Port: 9999

UDP:
  Addr: 127.0.0.1
  Port: 9998

HTTP:
  Addr: 127.0.0.1
  Port: 9997
`

type ServerInfo struct {
	Addr string `yaml:"Addr"`
	Port int    `yaml:"Port"`
}

type ConfigDemo struct {
	UDP       *ServerInfo `yaml:"UDP"`
	TCP       *ServerInfo `yaml:"TCP"`
	HTTP      *ServerInfo `yaml:"HTTP"`
	ProtoDir  string      `yaml:"ProtoDir"`
	PluginDir string      `yaml:"PluginDir"`
	LogLevel  string      `yaml:"LogLevel"`
	LogDir    string      `yaml:"LogDir"`
}

func TestYmlParser_Parse(t *testing.T) {
	c := &ConfigDemo{}
	j := &YmlParser{}
	if err := j.Parse([]byte(str), c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c)
}
