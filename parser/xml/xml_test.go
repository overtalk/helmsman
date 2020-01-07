package parser_test

import (
	"encoding/xml"
	"testing"

	. "helmsman/parser/xml"
)

const str = `<?xml version="1.0" encoding="utf-8" standalone="yes"?>

<!-- XX建议使用项目代号 -->
<meta name="SAUSAGE_SHOOT_OSS" tagsetversion="1" version="1">

    <include file="sausageshoot.xml" />

    <!-- 所有的日志struct 必须定义在下面的union中，否则无法转为pb并进行转换为对应语言的操作 -->
    <enumgroup name="OSS_LOG_TYPE" desc="LOG类型">
        <enum name="OSS_LOG_INVALID"                value="0"     desc="无效日志"/>
        <enum name="OSS_LOG_PLAYER_LOGIN"           value="1"     desc="玩家登陆"/>
        <enum name="OSS_LOG_PLAYER_LOGOUT"          value="2"     desc="玩家登出"/>
        <enum name="OSS_LOG_PLAYER_PASS_MISSION"    value="3"     desc="玩家通关"/>
        <enum name="OSS_LOG_PLAYER_LEVEL_UP"        value="4"     desc="玩家升级"/>
        <enum name="OSS_LOG_GAME_SERVER_STATE"      value="5"     desc="服务器状态"/>
    </enumgroup>
</meta>
`

type ConfigDemo struct {
	XMLName       xml.Name   `xml:"meta"`
	ProjectName   string     `xml:"name,attr"`
	Tagsetversion int        `xml:"tagsetversion,attr"`
	Version       int        `xml:"version,attr"`
	Includes      []*Include `xml:"include"`
	Enums         []*Enum    `xml:"enumgroup"`
}

type Enum struct {
	XMLName xml.Name    `xml:"enumgroup"`
	Name    string      `xml:"name,attr"`
	Desc    string      `xml:"desc,attr"`
	Items   []*EnumItem `xml:"enum"`
}

type EnumItem struct {
	XMLName xml.Name `xml:"enum"`
	Name    string   `xml:"name,attr"`
	Value   int      `xml:"value,attr"`
	Desc    string   `xml:"desc,attr"`
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	File    string   `xml:"file,attr"`
}

func TestXmlParser_Parse(t *testing.T) {
	c := &ConfigDemo{}
	j := &XmlParser{}
	if err := j.Parse([]byte(str), c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c.ProjectName)
	t.Log(c.Tagsetversion)
	t.Log(c.Version)

	for k, v := range c.Enums {
		t.Log(k, v)
	}
}
