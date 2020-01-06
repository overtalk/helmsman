package source_test

import (
	"testing"

	"github.com/caarlos0/env"

	. "config/source/gitlab"
)

func TestCatcher_Fetch(t *testing.T) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Error(err)
	}

	data, err := NewClient(cfg).Fetch("WeaponList.txt")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(data))
}
