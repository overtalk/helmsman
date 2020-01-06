package env_test

import (
	"os"
	"testing"

	. "config/source/env"
)

type ConfigDemo struct {
	Token string `env:"GITHUB_TOKEN,required"`
	Owner string `env:"GITHUB_OWNER,required"`
	Repo  string `env:"GITHUB_REPO,required"`
	Ref   string `env:"GITHUB_REF" envDefault:"master"`
	Path  string `env:"GITHUB_PATH,required"`
}

func TestEnvConfig_ParseConfig(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "test_token")
	os.Setenv("GITHUB_OWNER", "test_owner")
	os.Setenv("GITHUB_REPO", "test_repo")
	os.Setenv("GITHUB_REF", "test_ref")
	os.Setenv("GITHUB_PATH", "test_path")

	e := &EnvConfig{}
	c := &ConfigDemo{}
	if err := e.ParseConfig(nil, c); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", c)
}
