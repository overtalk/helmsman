package parser_test

import (
	"testing"

	. "helmsman/parser/toml"
)

const str = `
[env]
debug       = true
host        = "127.0.0.1"
port        = 1080
ips         = ["127.0.0.1", "127.0.0.2", "127.0.0.3"]
ids         = [1, 2, 3]

[dbs.master]
host        = "127.0.0.1"
port        = 3306

[dbs.slave]
host        = "127.0.0.1"
port        = 3306

[[fruit]]
name = "apple"

[fruit.physical]
color = "red"
shape = "round"

[[fruit.variety]]
name = "red delicious"

[[fruite.variety]]

[[fruit.variety]]
name = "granny smith"

[[fruit]]
name = "banana"

[[fruit.variety]]
`

type ConfigDemo struct {
	EnvConfig *EnvConfig           `toml:"env"`
	DBConfigs map[string]*DBConfig `toml:"dbs"`
	Fruits    []*Fruit             `toml:"fruit"`
}

type EnvConfig struct {
	Debug bool     `toml:"debug"`
	Host  string   `toml:"host"`
	Port  int      `toml:"port"`
	IPs   []string `toml:"ips"`
	Ids   []int    `toml:"ids"`
}

type DBConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Fruit struct {
	Name      string     `toml:"name"`
	Physical  *Physical  `toml:"physical"`
	Varieties []*Variety `toml:"variety"`
}

type Physical struct {
	Color string `toml:"color"`
	Shape string `toml:"shape"`
}

type Variety struct {
	Name string `toml:"name"`
}

func TestTomlParser_Parse(t *testing.T) {
	c := &ConfigDemo{}
	j := &TomlParser{}
	if err := j.Parse([]byte(str), c); err != nil {
		t.Error(err)
		return
	}

	t.Log("---------------------")
	for k, v := range c.DBConfigs {
		t.Logf("%s - %v", k, v)
	}
	t.Log("---------------------")

	t.Logf("%v", c.EnvConfig)
	t.Log("---------------------")

	for _, v := range c.Fruits {
		t.Log(v.Name, v.Physical)
		for k1, v1 := range v.Varieties {
			t.Log(k1, v1)
		}
	}
}
