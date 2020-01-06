package github

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/google/go-github/github"
)

func (c *Client) Fetch(path string) ([]byte, error) {
	fmt.Printf("%v", c.git.BaseURL)
	fileContent, _, _, err := c.git.Repositories.GetContents(context.Background(), c.cfg.Owner, c.cfg.Repo, path, &github.RepositoryContentGetOptions{Ref: c.cfg.Ref})
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(*fileContent.Content)
}
