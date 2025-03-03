package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func newClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: (30 * time.Second),
		},
		baseURL: baseURL,
	}
}

func (c *Client) getLocationAreas(pageURL string) (*LocationAreas, error) {
	return nil, nil
}
