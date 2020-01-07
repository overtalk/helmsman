package github

import (
	"helmsman"
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Config struct {
	Token string `env:"GITHUB_TOKEN,required"`
	Owner string `env:"GITHUB_OWNER,required"`
	Repo  string `env:"GITHUB_REPO,required"`
	Ref   string `env:"GITHUB_REF" envDefault:"master"`
	Path  string `env:"GITHUB_PATH,required"`
}

type Client struct {
	cfg *Config
	git *github.Client
}

func NewClient(c *Config) *Client {
	cli := github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.Token},
			)),
	)
	return &Client{
		git: cli,
		cfg: c,
	}
}

func (c *Client) Client() (*github.Client, error) {
	if c.git != nil {
		return nil, errors.New("github Client is nil")
	}
	return c.git, nil
}

func (c *Client) ParseConfig(t config.ParserType, cfg interface{}) error {
	data, err := c.Fetch(c.cfg.Path)
	if err != nil {
		return err
	}

	return config.GetParser(t).Parse(data, cfg)
}
