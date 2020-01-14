简体中文 | [English](./README.md)

<h1 align="center"> Helmsman </h1>
<p align="center"> Helmsman 是一个即用的 Go项目 配置的解决方案 </p>

--- 

[![Go Report Card](https://goreportcard.com/badge/github.com/overtalk/helmsman)](https://goreportcard.com/report/github.com/overtalk/helmsman)
[![Coverage Status](https://coveralls.io/repos/github/overtalk/helmsman/badge.svg?branch=master)](https://coveralls.io/github/overtalk/helmsman?branch=master)
[![Build Status](https://travis-ci.org/overtalk/helmsman.svg?branch=master)](https://travis-ci.org/overtalk/helmsman)
![GitHub stars](https://img.shields.io/github/stars/overtalk/helmsman.svg?style=flat-square&label=Stars&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues-raw/overtalk/helmsman.svg?style=flat-square)

## 为什么使用?
- [x] 多格式
    - [x] Json
    - [x] Yaml
    - [x] Envionment vars
    - [x] Xml
    - [ ] ini
- [x] 多数据源
    - [x] 本地文件
    - [x] Github 仓库
    - [x] Gitlab 仓库
    - [x] 阿里云 OSS
    - [x] Consul
    - [ ] Etcd
- [x] `数据格式` 与 `数据源`分离
- [ ] 配置变更通知

## 安装
```bash
go get github.com/overtalk/helmsman
```

## 如何使用?
### 本地文件 & json
```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/overtalk/helmsman"
	fileSource "github.com/overtalk/helmsman/source/file"
)

const str = `{
    "A": "this is the test of a ",
    "B": "this is the test of bbb"
}`

type ConfigDemo struct {
	A string `json:"A"`
	B string `json:"B"`
}

func GetConfig(path string) (*ConfigDemo, error) {
	c := &ConfigDemo{}
	fc := fileSource.NewFileConfig(path)
	if err := fc.ParseConfig(config.JsonFormat, c); err != nil {
		return nil, err
	}

	return c, nil
}

func main() {
	// write config data to tmp file
	tmpFile, err := ioutil.TempFile("", "temp_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tmpFile.Name()) // clean up

	// write data to the temp file
	if err := ioutil.WriteFile(tmpFile.Name(), []byte(str), 0777); err != nil {
		fmt.Println(err)
		return
	}

	c, err := GetConfig(tmpFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c.A)
	fmt.Println(c.B)
}

```

### `github 仓库` 数据源
```go
func GetConfigFromGithub() (*ConfigDemo, error) {
	c := &ConfigDemo{}

	cfg := &githubSource.Config{
		Token: "your github token",
		Owner: "your github name",
		Repo:  "your github repo",
		Ref:   "your repo branch",
		Path:  "your config file path in the repo",
	}
	githubClient := githubSource.NewClient(cfg)
	if err := githubClient.ParseConfig(config.JsonFormat, c); err != nil {
		return nil, err
	}

	return c, nil
}
```

## 配置 `struct` 生成工具
- [x] [json 转 golang struct 工具](https://mholt.github.io/json-to-go/)
- [x] [yaml 转 golang struct 工具](https://mengzhuo.github.io/yaml-to-go/)
- [x] [xml 转 golang struct 工具](https://www.onlinetool.io/xmltogo)