package aliyun_oss_test

import (
	"github.com/caarlos0/env"

	. "config/source/aliyun-oss"
)

func getClient() (*Client, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return NewClient(cfg)
}
