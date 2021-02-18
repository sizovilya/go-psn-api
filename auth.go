package psn

import (
	"fmt"
	"net/url"
)

const (
	authApi      = "https://auth.api.sonyentertainmentnetwork.com/2.0/"
	// clientId     = "7c01ce37-cb6b-4938-9c1b-9e36fd5477fa"
	clientId     = "71a7beb8-f21a-47d9-a604-2e71bee24fe0"
	// clientSecret = "GNumO5QMsagNcO2q"
	clientSecret = "xSk2YI8qJqZfeLQv"
	duid         = "00000007000801a8000000000000008241fdf6ab09ba863a20202020476f6f676c653a416e64726f696420534400000000000000000000000000000000"
	// scope        = "kamaji:get_players_met+kamaji:get_account_hash+kamaji:activity_feed_submit_feed_story+kamaji:activity_feed_internal_feed_submit_story+kamaji:activity_feed_get_news_feed+kamaji:communities+kamaji:game_list+kamaji:ugc:distributor+oauth:manage_device_usercodes+psn:sceapp+user:account.profile.get+user:account.attributes.validate+user:account.settings.privacy.get+kamaji:activity_feed_set_feed_privacy+kamaji:satchel+kamaji:satchel_delete+user:account.profile.update+kamaji:url_preview"
	scope        = "openid:age openid:content_ctrl kamaji:get_privacy_settings kamaji:get_account_hash openid:user_id openid:ctry_code openid:lang"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires int32 `json:"expires_in"`
}

func (p *psn) authRequest() (tokens *tokens, err error) {
	form := url.Values{}
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
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

func (p *psn) refreshTokens() (tokens *tokens, err error) {
	if p.refreshToken == "" {
		return nil, fmt.Errorf("can't refresh tokens, refresh token is empty")
	}
	form := url.Values{}
	form.Add("app_context", "inapp_ios")
	form.Add("client_id", clientId)
	form.Add("client_secret", clientSecret)
	form.Add("refresh_token", p.refreshToken)
	form.Add("duid", duid)
	form.Add("scope", scope)
	form.Add("grant_type", "refresh_token")

	var h = headers{}
	h["Content-Type"] = "application/x-www-form-urlencoded"

	err = p.post(form, fmt.Sprintf("%soauth/token", authApi), h, &tokens)
	if err != nil {
		return nil, fmt.Errorf("can't do post request: %w", err)
	}

	return tokens, nil
}