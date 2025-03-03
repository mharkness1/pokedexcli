package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: (30 * time.Second),
		},
		baseURL: baseURL,
	}
}

func (c *Client) GetLocationAreas(pageURL string) (*LocationAreas, error) {
	if pageURL == "" {
		pageURL = c.baseURL + "/location-area"
	}

	res, err := http.Get(pageURL)

	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("http error status code: %v", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var locations *LocationAreas
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("error in json.unmarshal: %v", err)
	}

	return locations, nil
}
