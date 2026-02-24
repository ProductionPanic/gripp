package gripp

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	apiKey              string
	url                 string
	apiconnectorversion int
}

type Config struct {
	ApiKey              string
	Url                 string
	ApiConnectorVersion int
}

func NewClient(config Config) (*Client, error) {
	if config.Url == "" {
		config.Url = "https://api.gripp.com/public/api3.php"
	}

	if config.ApiConnectorVersion == 0 {
		config.ApiConnectorVersion = 3011
	}

	if config.ApiKey == "" {
		return nil, ErrMissingApiKey
	}

	return &Client{
		apiKey:              config.ApiKey,
		url:                 config.Url,
		apiconnectorversion: config.ApiConnectorVersion,
	}, nil
}

// helper functions for actually making the requests
func (c *Client) makeRequest(request RequestType) ([]Response, error) {
	data, err := json.Marshal(request)

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(out))

	var response []Response
	err = json.Unmarshal(out, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
