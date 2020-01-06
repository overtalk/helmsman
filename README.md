# Golang Configuration Hub

## Why use?
- [x] multi-format
    - [x] Json
    - [x] Yaml
    - [x] Envionment vars
    - [x] Xml
    - [ ] ini
- [x] multi-source
    - [x] Local File
    - [x] Github Repo
    - [x] Gitlab Repo
    - [x] Aliyun OSS
    - [ ] Etcd
    - [ ] Consul
- [x] data format & data source separation
- [ ] watching changes

## Install
```bash
go get github.com/qinhan-shu/config
```


## How to use?
### file source & json format
```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qinhan-shu/config"
	fileSource "github.com/qinhan-shu/config/source/file"
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

### change to github config source
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

## Tools
- [ ] Config Struct Builder tool