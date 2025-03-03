package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// API main URL
const baseURL = "https://pokeapi.co/api/v2"

// Client struct with pointer to default client struct and a baseURL - alter with more commands.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Creates a new client to make request, returns pointer to default client with timeout at 30s and baseURL as above.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: (30 * time.Second),
		},
		baseURL: baseURL,
	}
}

// Actual get function creates URL resturns pointer to the unmarshalled JSON.
func (c *Client) GetLocationAreas(pageURL string) (*LocationAreas, error) {
	if pageURL == "" {
		pageURL = c.baseURL + "/location-area"
	}
	// GET action, checks for error in internal call.
	res, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	// Checks for error from API
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("http error status code: %v", res.StatusCode)
	}

	// Reads body of response.
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Creates struct for response, and unmarshalls the json into the struct.
	var locations *LocationAreas
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("error in json.unmarshal: %v", err)
	}

	return locations, nil
}
