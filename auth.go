package psn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	authHost             = "https://ca.account.sony.com/api/authz/v3/oauth"
	redirectURI          = "com.playstation.PlayStationApp://redirect"
	clientID             = "ac8d161a-d966-4728-b0ea-ffec22f69edc"
	scope                = "psn:mobile.v1 psn:clientapp"
	tokenFormat          = "jwt"
	basicAuthCredentials = "Basic YWM4ZDE2MWEtZDk2Ni00NzI4LWIwZWEtZmZlYzIyZjY5ZWRjOkRFaXhFcVhYQ2RYZHdqMHY="
)

type tokens struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	AccessExpires  int32  `json:"expires_in"`
	RefreshExpires int32  `json:"refresh_token_expires_in"`
}

// AuthWithNPSSO authenticates the client using an NPSSO code.
func (c *Client) AuthWithNPSSO(ctx context.Context, npsso string) error {
	if npsso == "" {
		return fmt.Errorf("npsso code is empty")
	}
	c.npsso = npsso

	code, err := c.getAuthorizationCode(ctx, npsso)
	if err != nil {
		return fmt.Errorf("failed to get authorization code: %w", err)
	}

	tokens, err := c.exchangeCodeForTokens(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to exchange authorization code for tokens: %w", err)
	}

	c.setTokens(tokens)
	return nil
}

// AuthWithRefreshToken authenticates the client using a refresh token.
func (c *Client) AuthWithRefreshToken(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}

	postValues := url.Values{}
	postValues.Add("scope", scope)
	postValues.Add("refresh_token", refreshToken)
	postValues.Add("grant_type", "refresh_token")
	postValues.Add("token_format", tokenFormat)

	headers := headers{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": basicAuthCredentials,
	}

	var t tokens
	err := c.post(ctx, fmt.Sprintf("%s/token", authHost), postValues, headers, &t)
	if err != nil {
		return fmt.Errorf("failed to refresh tokens: %w", err)
	}

	c.setTokens(&t)
	return nil
}

func (c *Client) getAuthorizationCode(ctx context.Context, npsso string) (string, error) {
	getValues := url.Values{}
	getValues.Add("access_type", "offline")
	getValues.Add("client_id", clientID)
	getValues.Add("response_type", "code")
	getValues.Add("scope", scope)
	getValues.Add("redirect_uri", redirectURI)
	getValues.Add("app_context", "inapp_ios")
	getValues.Add("auth_ver", "v3")
	getValues.Add("cid", "60351282-8C5F-4D5E-9033-E48FEA973E11")
	getValues.Add("darkmode", "true")
	getValues.Add("device_base_font_size", "10")
	getValues.Add("device_profile", "mobile")
	getValues.Add("duid", "0000000d0004008088347AA0C79542D3B656EBB51CE3EBE1")
	getValues.Add("elements_visibility", "no_aclink")
	getValues.Add("extraQueryParams", `{PlatformPrivacyWs1=minimal;}`)
	getValues.Add("no_captcha", "true")
	getValues.Add("service_entity", "urn:service-entity:psn")
	getValues.Add("service_logo", "ps")
	getValues.Add("smcid", "psapp:settings-entrance")
	getValues.Add("support_scheme", "sneiprls")
	getValues.Add("token_format", tokenFormat)
	getValues.Add("ui", "pr")

	uri, _ := url.Parse(fmt.Sprintf("%s/authorize", authHost))
	uri.RawQuery = getValues.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", uri.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Cookie", fmt.Sprintf("npsso=%s", npsso))

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		return "", fmt.Errorf("expected redirect, got status %d", resp.StatusCode)
	}

	location, err := resp.Location()
	if err != nil {
		return "", fmt.Errorf("failed to get redirect location: %w", err)
	}

	if errorDesc := location.Query().Get("error_description"); errorDesc != "" {
		return "", fmt.Errorf("authentication failed: %s", errorDesc)
	}

	code := location.Query().Get("code")
	if code == "" {
		return "", fmt.Errorf("authorization code not found in redirect URL")
	}

	return code, nil
}

func (c *Client) exchangeCodeForTokens(ctx context.Context, code string) (*tokens, error) {
	postValues := url.Values{}
	postValues.Add("code", code)
	postValues.Add("redirect_uri", redirectURI)
	postValues.Add("grant_type", "authorization_code")
	postValues.Add("token_format", tokenFormat)
	postValues.Add("smcid", "psapp%3Asettings-entrance")
	postValues.Add("access_type", "offline")
	postValues.Add("service_logo", "ps")
	postValues.Add("ui", "pr")
	postValues.Add("elements_visibility", "no_aclink")
	postValues.Add("support_scheme", "sneiprls")
	postValues.Add("darkmode", "true")
	postValues.Add("device_base_font_size", "10")
	postValues.Add("device_profile", "mobile")
	postValues.Add("app_context", "inapp_ios")
	postValues.Add("extraQueryParams", `{PlatformPrivacyWs1=minimal;}`)

	headers := headers{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": basicAuthCredentials,
		"Cookie":        fmt.Sprintf("npsso=%s", c.npsso),
	}

	var t tokens
	err := c.post(ctx, fmt.Sprintf("%s/token", authHost), postValues, headers, &t)
	if err != nil {
		return nil, fmt.Errorf("token exchange request failed: %w", err)
	}

	return &t, nil
}

func (c *Client) setTokens(t *tokens) {
	c.accessToken = t.AccessToken
	c.refreshToken = t.RefreshToken
	c.accessExp = t.AccessExpires
	c.refreshExp = t.RefreshExpires
}