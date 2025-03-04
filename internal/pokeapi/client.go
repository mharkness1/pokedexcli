package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mharkness1/pokedexcli/internal/pokecache"
)

// API main URL
const baseURL = "https://pokeapi.co/api/v2"

// Client struct with pointer to default client struct and a baseURL - alter with more commands.
type Client struct {
	baseURL    string
	cache      pokecache.Cache
	httpClient *http.Client
}

// Creates a new client to make request, returns pointer to default client with timeout at 30s and baseURL as above.
func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(20 * time.Minute),
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

	if val, ok := c.cache.Get(pageURL); ok {
		var locations *LocationAreas
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return nil, err
		}
		return locations, nil
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

	c.cache.Add(pageURL, body)

	return locations, nil
}

func (c *Client) GetExploreResults(AreaName string) (*ExploreResults, error) {
	if AreaName == "" {
		fmt.Println("invalid area name")
		return nil, nil
	}
	pageURL := c.baseURL + "/location-area/" + AreaName

	if val, ok := c.cache.Get(pageURL); ok {
		var exploreResults *ExploreResults
		err := json.Unmarshal(val, &exploreResults)
		if err != nil {
			return nil, err
		}
		return exploreResults, nil
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

	var exploreResults *ExploreResults
	if err := json.Unmarshal(body, &exploreResults); err != nil {
		return nil, fmt.Errorf("error in json.unmarshal: %v", err)
	}

	c.cache.Add(pageURL, body)

	return exploreResults, nil
}
