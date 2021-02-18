package psn

import (
	"fmt"
	"net/url"
)

const (
	authApi      = "https://auth.api.sonyentertainmentnetwork.com/2.0/"
	scope        = "openid:age openid:content_ctrl kamaji:get_privacy_settings kamaji:get_account_hash openid:user_id openid:ctry_code openid:lang"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	Expires int32 `json:"expires_in"`
}

func (p *psn) authRequest() (tokens *tokens, err error) {
	form := url.Values{}
	form.Add("client_id", p.clientId)
	form.Add("client_secret", p.clientSecret)
	form.Add("scope", scope)
	form.Add("grant_type", "sso_cookie")

	var h = headers{}
	h["Content-Type"] = "application/x-www-form-urlencoded"
	h["Cookie"] = fmt.Sprintf("npsso=%s", p.npsso)

	err = p.post(form, fmt.Sprintf("%soauth/token", authApi), h, &tokens)
	if err != nil {
		return nil, fmt.Errorf("can't do post request: %w", err)
	}
	return tokens, nil
}