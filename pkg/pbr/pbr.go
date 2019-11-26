package pbr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type OS struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Package struct {
	Environment string `json:"environment"`
	Name        string `json:"name"`
	NetworkType string `json:"networkType"`
	Description string `json:"description"`
	Protocol    string `json:"protocol"`
	Subtype     string `json:"subtype"`
}

type Version struct {
	OS          OS       `json:"os"`
	Package     *Package `json:"package,omitempty"`
	RegistryURL string   `json:"registryUrl"`
	Version     string   `json:"version"`
}

type CLIVersion struct {
	OS          OS     `json:"os"`
	RegistryURL string `json:"registryUrl"`
	Version     string `json:"version"`
}

// response is a non-discernible response format from secrets api
type response struct {
	Data  json.RawMessage   `json:"data,omitempty"`
	Error map[string]string `json:"error,omitempty"`
}

type Client struct {
	host       string
	httpClient *http.Client
}

func New(host string) *Client {
	return &Client{
		host: host,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// newRequest creates a new safe http.Request.
func newRequest(method, host, path, query string, body interface{}) (*http.Request, error) {
	var buf bytes.Buffer

	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%s/v1/%s?%s", host, path, query)

	// Create HTTP request
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// do makes a request to the Safe API
func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for http status no content
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Decode into standard response from API
	sr := &response{}
	if err := json.NewDecoder(resp.Body).Decode(sr); err != nil {
		return err
	}

	// Check status for error code
	if resp.StatusCode >= 400 {
		return errors.New(sr.Error["message"])
	}

	if v != nil {
		if err := json.Unmarshal(sr.Data, v); err != nil {
			return err
		}
	}

	return nil
}
