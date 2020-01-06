package aliyun_oss

import (
	"config"
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
Config demo
{
	accessKey : "xxx"
	secretKey : "xxx"
	endpoint  : "endpoint=oss-cn-shanghai.aliyuncs.com"
	bucket    : "qinhan"
	dir       : "test/"
}
=> 对于的取文件的格式
=> qinhan@oss-cn-shanghai.aliyuncs.com/test/xxx
*/

type Config struct {
	AccessKey string `env:"ALIYUN_AK"`
	SecretKey string `env:"ALIYUN_SK"`
	Endpoint  string `env:"ALIYUN_ENDPOINT"`
	Bucket    string `env:"ALIYUN_BUCKET"`
	Path      string `env:"ALIYUN_PATH"`
}

type Client struct {
	cfg    *Config
	client *oss.Client
	bucket *oss.Bucket
}

func NewClient(cfg *Config) (*Client, error) {
	cli, err := oss.New(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey)
	if err != nil {
		return nil, err
	}

	bucket, err := cli.Bucket(cfg.Bucket)
	if err != nil {
		return nil, err
	}
	return &Client{
		cfg:    cfg,
		client: cli,
		bucket: bucket,
	}, nil
}

func (f *Client) Bucket() (*oss.Bucket, error) {
	if f.bucket != nil {
		return nil, errors.New("oss bucket is nil")
	}
	return f.bucket, nil
}

func (f *Client) Client() (*oss.Client, error) {
	if f.bucket != nil {
		return nil, errors.New("oss Client is nil")
	}
	return f.client, nil
}

func (f *Client) ParseConfig(parser config.Parser, cfg interface{}) error {
	data, err := f.GetObject(f.cfg.Path)
	if err != nil {
		return err
	}

	return parser.Parse(data, cfg)
}
