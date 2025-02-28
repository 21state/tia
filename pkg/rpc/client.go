package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	userAgent = "tia-cli-explorer"
	timeout   = 30 * time.Second
)

// Client represents an RPC client for interacting with a Celestia node
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new RPC client
func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// GetBlock retrieves block information by height
// If height is 0, the latest block will be fetched
func (c *Client) GetBlock(ctx context.Context, height int64) (ResultBlock, error) {
	args := make(map[string]string)
	if height != 0 {
		args["height"] = strconv.FormatInt(height, 10)
	}

	var response Response[ResultBlock]
	if err := c.get(ctx, "block", args, &response); err != nil {
		return ResultBlock{}, err
	}

	if response.Error != nil {
		return ResultBlock{}, response.Error
	}

	return response.Result, nil
}

// get performs a GET request to the RPC endpoint
func (c *Client) get(ctx context.Context, path string, args map[string]string, output interface{}) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path, err = url.JoinPath(u.Path, path)
	if err != nil {
		return err
	}

	values := u.Query()
	for key, value := range args {
		values.Add(key, value)
	}
	u.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(output)
}
