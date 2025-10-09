package psn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type headers map[string]string

func (c *Client) get(ctx context.Context, url string, headers headers, value interface{}) error {
	return c.do(ctx, "GET", url, nil, headers, value)
}

func (c *Client) post(ctx context.Context, url string, formData url.Values, headers headers, value interface{}) error {
	return c.do(ctx, "POST", url, strings.NewReader(formData.Encode()), headers, value)
}

func (c *Client) do(ctx context.Context, method, url string, body io.Reader, headers headers, value interface{}) (err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %w", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}