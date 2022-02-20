package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	gh "github.com/google/go-github/v32/github"
)

type GitHubClientInterface interface {
	FetchLatestTag(string, string) (string, error)
}

type Client struct {
	Client *gh.Client
}

func NewClient() *Client {
	c := Client{
		Client: github.NewClient(nil),
	}
	return &c
}

func (c *Client) FetchLatestTag(owner, repo string) (string, error) {
	rel, _, err := c.Client.Repositories.ListReleases(context.Background(), owner, repo, nil)
	if err != nil {
		return "", err
	}
	var latestTag string
	for _, rel := range rel {
		r := *rel
		latestTag = r.GetTagName()
		return latestTag, nil
	}
	return "", nil
}
