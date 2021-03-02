package psn

import (
	"fmt"
	"net/http"
	"net/url"
	// "encoding/json"
	// "bytes"
)

const (
	authUrl = "https://ca.account.sony.com/api/"
)

type tokens struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	AccessExpires  int32  `json:"expires_in"`
	RefreshExpires int32  `json:"refresh_token_expires_in"`
}

// Method makes auth request to Sony's server and retrieves tokens
func (p *psn) AuthWithNPSSO() error {
	tokens, err := p.authRequest()
	if err != nil {
		return fmt.Errorf("can't do auth request: %w", err)
	}
	p.accessToken = tokens.AccessToken
	return nil
}

// Method makes auth request to Sony's server and retrieves tokens
// TODO: write logic
func (p *psn) AuthWithRefreshToken() error {
	tokens, err := p.authRequest()
	if err != nil {
		return fmt.Errorf("can't do auth request: %w", err)
	}
	p.accessToken = tokens.AccessToken
	return nil
}

func (p *psn) authRequest() (tokens *tokens, err error) {
	getValues := url.Values{}
	getValues.Add("access_type", "offline")
	getValues.Add("app_context", "inapp_ios")
	getValues.Add("auth_ver", "v3")
	getValues.Add("cid", "60351282-8C5F-4D5E-9033-E48FEA973E11")
	getValues.Add("client_id", "ac8d161a-d966-4728-b0ea-ffec22f69edc")
	getValues.Add("darkmode", "true")
	getValues.Add("device_base_font_size", "10")
	getValues.Add("device_profile", "mobile")
	getValues.Add("duid", "0000000d0004008088347AA0C79542D3B656EBB51CE3EBE1")
	getValues.Add("elements_visibility", "no_aclink")
	getValues.Add("extraQueryParams", `{PlatformPrivacyWs1=minimal;}`)
	getValues.Add("no_captcha", "true")
	getValues.Add("redirect_uri", "com.playstation.PlayStationApp://redirect")
	getValues.Add("response_type", "code")
	getValues.Add("scope", "psn:mobile.v1 psn:clientapp")
	getValues.Add("service_entity", "urn:service-entity:psn")
	getValues.Add("service_logo", "ps")
	getValues.Add("smcid", "psapp:settings-entrance")
	getValues.Add("support_scheme", "sneiprls")
	getValues.Add("token_format", "jwt")
	getValues.Add("ui", "pr")

	var getHeaders = headers{}
	getHeaders["Cookie"] = fmt.Sprintf("npsso=%s", p.npsso)

	uri, _ := url.Parse(fmt.Sprintf("%sauthz/v3/oauth/authorize", authUrl))
	uri.RawQuery = getValues.Encode()

	var i int
	var code = ""
	nextUrl := uri.String()
	for i < 2 {
		req, err := http.NewRequest(
			"GET",
			nextUrl,
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("can't create new GET request %w: ", err)
		}

		for k, v := range getHeaders {
			req.Header.Add(k, v)
		}

		// create new httpclient with ability to check redirects
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("can't execute GET request %w: ", err)
		}

		defer func() {
			err = resp.Body.Close()
		}()

		if resp.StatusCode == http.StatusFound {
			nextUrl = resp.Header.Get("Location")
			parsed, _ := url.ParseQuery(nextUrl)
			if len(parsed["com.playstation.PlayStationApp://redirect/?code"]) > 0 {
				code = parsed["com.playstation.PlayStationApp://redirect/?code"][0]
				break
			}
			i += 1
		}
	}

	if code == "" {
		return nil, fmt.Errorf("code doesn't retrieved from redirect")
	}

	postValues := url.Values{}
	postValues.Add("smcid", "psapp%3Asettings-entrance")
	postValues.Add("access_type", "offline")
	postValues.Add("code", code)
	postValues.Add("service_logo", "ps")
	postValues.Add("ui", "pr")
	postValues.Add("elements_visibility", "no_aclink")
	postValues.Add("redirect_uri", "com.playstation.PlayStationApp://redirect")
	postValues.Add("support_scheme", "sneiprls")
	postValues.Add("grant_type", "authorization_code")
	postValues.Add("darkmode", "true")
	postValues.Add("device_base_font_size", "10")
	postValues.Add("device_profile", "mobile")
	postValues.Add("app_context", "inapp_ios")
	postValues.Add("extraQueryParams", `{PlatformPrivacyWs1=minimal;}`)
	postValues.Add("token_format", "jwt")

	var postHeaders = headers{}
	postHeaders["Content-Type"] = "application/x-www-form-urlencoded"
	postHeaders["Cookie"] = fmt.Sprintf("npsso=%s", p.npsso)
	postHeaders["Authorization"] = "Basic YWM4ZDE2MWEtZDk2Ni00NzI4LWIwZWEtZmZlYzIyZjY5ZWRjOkRFaXhFcVhYQ2RYZHdqMHY="

	err = p.post(postValues, fmt.Sprintf("%sauthz/v3/oauth/token", authUrl), postHeaders, &tokens)
	if err != nil {
		return nil, fmt.Errorf("can't create new POST request %w: ", err)
	}

	return tokens, nil
}
