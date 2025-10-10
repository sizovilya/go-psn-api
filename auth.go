package psn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	authHost             = "https://ca.account.sony.com/api/"
	scope                = "psn:mobile.v2.core psn:clientapp"
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

	fmt.Println("code:", code)

	tokens, err := c.exchangeCodeForTokens(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to exchange authorization code for tokens: %w", err)
	}

	fmt.Println("tokens:", tokens)

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

func (c *Client) getAuthorizationCode(ctx context.Context, npsso string) (_ string, err error) {
	getValues := url.Values{
		"access_type":           {"offline"},
		"app_context":           {"inapp_ios"},
		"auth_ver":              {"v3"},
		"cid":                   {"60351282-8C5F-4D5E-9033-E48FEA973E11"},
		"client_id":             {"09515159-7237-4370-9b40-3806e67c0891"},
		"darkmode":              {"true"},
		"device_base_font_size": {"10"},
		"device_profile":        {"mobile"},
		"duid":                  {"0000000d0004008088347AA0C79542D3B656EBB51CE3EBE1"},
		"elements_visibility":   {"no_aclink"},
		"extraQueryParams":      {`{ PlatformPrivacyWs1 = minimal; }`},
		"no_captcha":            {"true"},
		"redirect_uri":          {"com.scee.psxandroid.scecompcall://redirect"},
		"response_type":         {"code"},
		"scope":                 {"psn:mobile.v2.core psn:clientapp"},
		"service_entity":        {"urn:service-entity:psn"},
		"service_logo":          {"ps"},
		"smcid":                 {"psapp:settings-entrance"},
		"support_scheme":        {"sneiprls"},
		"token_format":          {"jwt"},
		"ui":                    {"pr"},
	}

	uri, _ := url.Parse(fmt.Sprintf("%s/authz/v3/oauth/authorize", authHost))
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
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close response body: %w", cerr)
		}
	}()

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
	postValues := url.Values{
		"smcid":                 {"psapp%3Asettings-entrance"},
		"access_type":           {"offline"},
		"code":                  {code},
		"service_logo":          {"ps"},
		"ui":                    {"pr"},
		"elements_visibility":   {"no_aclink"},
		"redirect_uri":          {"com.scee.psxandroid.scecompcall://redirect"},
		"support_scheme":        {"sneiprls"},
		"grant_type":            {"authorization_code"},
		"darkmode":              {"true"},
		"device_base_font_size": {"10"},
		"device_profile":        {"mobile"},
		"app_context":           {"inapp_ios"},
		"extraQueryParams":      {`{ PlatformPrivacyWs1 = minimal; }`},
		"token_format":          {"jwt"},
	}

	headers := headers{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic MDk1MTUxNTktNzIzNy00MzcwLTliNDAtMzgwNmU2N2MwODkxOnVjUGprYTV0bnRCMktxc1A=",
		"Cookie":        fmt.Sprintf("npsso=%s", c.npsso),
	}

	var t tokens
	err := c.post(ctx, fmt.Sprintf("%s/authz/v3/oauth/token", authHost), postValues, headers, &t)
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
