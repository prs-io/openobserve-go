package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	ioutil "io"
	"net/http"
	"net/url"
)

// Client is the main structure for the OpenObserve client.
type Client struct {
	config     *Config
	httpClient *http.Client
}

// NewClient initializes a new OpenObserve client with the provided configuration.
func NewClient(config Config) *Client {
	client := &Client{
		config: &config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
	return client
}

// Do sends an HTTP request and returns an HTTP response.
// func (c *Client) Do(req *http.Request) (*http.Response, error) {
// 	creds := c.config.Username + ":" + c.config.Password
// 	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))
// 	req.Header.Set("Authorization", basicAuth)
// 	req.Header.Set("Content-Type", "application/json")
// 	return c.httpClient.Do(req)
// }

func (c *Client) Do(method, path string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}
	// %s/orgs/%s
	baseUrl, err := url.JoinPath(c.config.BaseURL, path)

	//fmt.Println("baseUrl", baseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to join URL path: %w", err)
	}
	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	creds := c.config.Username + ":" + c.config.Password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))
	req.Header.Set("Authorization", basicAuth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	errorStatusCodes := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
	}

	fmt.Printf("ResponseCli: %v\n", resp.StatusCode)

	if containsStatus(resp.StatusCode, errorStatusCodes) {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return resp, fmt.Errorf("request failed with status %s: %s", resp.Status, bodyBytes)
	}

	return resp, nil
}

func containsStatus(code int, codes []int) bool {
	for _, c := range codes {
		if code == c {
			return true
		}
	}
	return false
}

// func (c *Client) Decode(v interface{}, b []byte, contentType string) (err error) {

// }
