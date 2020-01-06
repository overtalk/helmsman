package aliyun_oss

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func (f *Client) GetObject(path string) ([]byte, error) {
	body, err := f.bucket.GetObject(path)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file %s is empty", path)
	}

	return data, nil
}

func (f *Client) PutObject(objectKey, data string) error {
	return f.bucket.PutObject(objectKey, strings.NewReader(data))
}
