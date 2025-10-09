package psn

import (
	"fmt"
	"net/http"
)

// Client represents the PSN API client.
type Client struct {
	http         *http.Client
	lang         string
	region       string
	npsso        string
	accessToken  string
	refreshToken string
	accessExp    int32
	refreshExp   int32
}

// Options holds the configuration for the PSN API client.
type Options struct {
	Lang   string
	Region string
	Npsso  string
}

// NewClient creates a new PSN API client.
func NewClient(opts *Options) (*Client, error) {
	if !isContain(languages, opts.Lang) {
		return nil, fmt.Errorf("unsupported language: %s", opts.Lang)
	}
	if !isContain(regions, opts.Region) {
		return nil, fmt.Errorf("unsupported region: %s", opts.Region)
	}
	return &Client{
		http:   &http.Client{},
		lang:   opts.Lang,
		region: opts.Region,
		npsso:  opts.Npsso,
	}, nil
}

// Lang returns the client's language.
func (c *Client) Lang() string {
	return c.lang
}

// Region returns the client's region.
func (c *Client) Region() string {
	return c.region
}

// Npsso returns the client's NPSSO code.
func (c *Client) Npsso() string {
	return c.npsso
}

// AccessToken returns the access token and its expiration time.
func (c *Client) AccessToken() (string, int32) {
	return c.accessToken, c.accessExp
}

// RefreshToken returns the refresh token and its expiration time.
func (c *Client) RefreshToken() (string, int32) {
	return c.refreshToken, c.refreshExp
}