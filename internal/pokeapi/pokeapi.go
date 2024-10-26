package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
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
	return &locationsResponse, nil
}
