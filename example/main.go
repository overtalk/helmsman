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
