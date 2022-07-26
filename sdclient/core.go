// Package sdclient implements methods to interact with the Sysdig Monitoring API.
package sdclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client is the client for the Sysdig Monitoring API.
type Client struct {
	HTTPClient *http.Client
	Endpoint   string
	ApiKey     string
}

// New creates a new Sysdig Monitoring API client.
func New() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// WithEndpoint sets the endpoint of the Sysdig Monitoring API.
func (c *Client) WithEndpoint(endpoint string) *Client {
	c.Endpoint = endpoint
	return c
}

// WithAPIKey sets the API key of the Sysdig Monitoring API.
func (c *Client) WithAPIKey(apiKey string) *Client {
	c.ApiKey = apiKey
	return c
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	if req.Method == http.MethodDelete {
		return nil
	}

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
