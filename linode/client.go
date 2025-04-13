package linode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

// Client represents a Linode API client
type Client struct {
	client *linodego.Client
}

// NewClient creates a new Linode client with the provided API token
func NewClient(apiToken string) *Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiToken})
	oauth2Transport := &oauth2.Transport{
		Source: tokenSource,
	}

	httpClient := &http.Client{
		Transport: oauth2Transport,
	}

	client := linodego.NewClient(httpClient)
	return &Client{
		client: &client,
	}
}

// ListInstances returns a list of all Linode instances
func (c *Client) ListInstances(ctx context.Context) ([]linodego.Instance, error) {
	instances, err := c.client.ListInstances(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %w", err)
	}
	return instances, nil
}
