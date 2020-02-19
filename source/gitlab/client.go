package source

import (
	"github.com/overtalk/helmsman"
	"errors"

	"github.com/xanzy/go-gitlab"
)

type Config struct {
	Token string `env:"GITLAB_TOKEN,required"`
	Ref   string `env:"GITLAB_REF" envDefault:"master"`
	Pid   int    `env:"GITLAB_PID,required"`
	Url   string `env:"GITLAB_URL"`
	Path  string `env:"GITLAB_PATH,required"`
}

type Client struct {
	cfg *Config
	git *gitlab.Client
}

func NewClient(c *Config) *Client {
	git := gitlab.NewClient(nil, c.Token)
	// 一般默认为`https://gitlab.com/api/v4`,设置之后才进行修改
	if len(c.Url) != 0 {
		git.SetBaseURL(c.Url)
	}
	return &Client{
		cfg: c,
		git: git,
	}
}

func (c *Client) Client() (*gitlab.Client, error) {
	if c.git != nil {
		return nil, errors.New("gitlab Client is nil")
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
