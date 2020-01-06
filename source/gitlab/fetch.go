package source

import (
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Fetch(path string) ([]byte, error) {
	data, _, err := c.git.RepositoryFiles.GetRawFile(c.cfg.Pid, path, &gitlab.GetRawFileOptions{Ref: &c.cfg.Ref})
	if err != nil {
		return nil, err
	}
	return data, nil
}
