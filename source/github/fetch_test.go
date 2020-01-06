package github_test

import (
	"testing"

	"github.com/caarlos0/env"

	. "config/source/github"
)

func TestCatcher_Fetch(t *testing.T) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Error(err)
	}
	fetcher := NewClient(cfg)

	data, err := fetcher.Fetch("docker/docker.md")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(data))
}
