package parser_test

import (
	"testing"

	. "helmsman/parser/json"
)

const str = `{
    "A": "this is the test of a ",
    "B": "this is the test of bbb"
}`

type ConfigDemo struct {
	A string `json:"A"`
	B string `json:"B"`
}

func TestJsonParser_Parse(t *testing.T) {
	c := &ConfigDemo{}
	j := &JsonParser{}
	if err := j.Parse([]byte(str), c); err != nil {
		t.Error(err)
		return
	}

	t.Log(c.A)
	t.Log(c.B)
}
