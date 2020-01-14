package consul

import (
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/mitchellh/consulstructure"
)

type ConsulConfig struct {
	Consul     *consulAPI.Config
	Prefix     string
	UpdateChan chan interface{}
	ErrChan    chan error
	decoder    *consulstructure.Decoder
}

func (c *ConsulConfig) ParseConfig(cfg interface{}) {
	c.decoder = &consulstructure.Decoder{
		Consul:   consulAPI.DefaultConfig(),
		Target:   cfg,
		Prefix:   c.Prefix,
		UpdateCh: c.UpdateChan,
		ErrCh:    c.ErrChan,
	}

	c.decoder.Run()
}

func (c *ConsulConfig) Close() {
	c.decoder.Close()
}
