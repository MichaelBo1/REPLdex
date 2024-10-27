package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/MichaelBo1/repldex/internal/pokecache"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, purgeInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(purgeInterval),
	}
}

type ShallowLocationAreas struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func (c *Client) ListLocations(pageURL *string) (*ShallowLocationAreas, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Try get cached response
	if val, ok := c.cache.Get(url); ok {
		// fmt.Printf("Using cache for URL: %s", url)
		locationsResponse := ShallowLocationAreas{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return nil, err
		}

		return &locationsResponse, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	locationsResponse := ShallowLocationAreas{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)
	return &locationsResponse, nil
}
