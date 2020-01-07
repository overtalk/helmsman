package source_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"helmsman"
	. "helmsman/source/file"
)

const str = `{
    "A": "this is the test of a ",
    "B": "this is the test of bbb"
}`

type ConfigDemo struct {
	A string `json:"A"`
	B string `json:"B"`
}

func TestFileConfig_ParseConfig(t *testing.T) {
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

	fc := NewFileConfig(tmpFile.Name())
	c := &ConfigDemo{}
	if err := fc.ParseConfig(config.JsonFormat, c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c.A)
	t.Log(c.B)
}
