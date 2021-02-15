package psn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AuthApi      = "https://auth.api.sonyentertainmentnetwork.com/2.0/"
	UsersApi     = "-prof.np.community.playstation.net/userProfile/v1/users/"
	ClientId     = "7c01ce37-cb6b-4938-9c1b-9e36fd5477fa"
	ClientSecret = "GNumO5QMsagNcO2q"
	Duid         = "00000007000801a8000000000000008241fdf6ab09ba863a20202020476f6f676c653a416e64726f696420534400000000000000000000000000000000"
	Scope        = "kamaji:get_players_met+kamaji:get_account_hash+kamaji:activity_feed_submit_feed_story+kamaji:activity_feed_internal_feed_submit_story+kamaji:activity_feed_get_news_feed+kamaji:communities+kamaji:game_list+kamaji:ugc:distributor+oauth:manage_device_usercodes+psn:sceapp+user:account.profile.get+user:account.attributes.validate+user:account.settings.privacy.get+kamaji:activity_feed_set_feed_privacy+kamaji:satchel+kamaji:satchel_delete+user:account.profile.update+kamaji:url_preview"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (p *psn) authRequest() (tokens *tokens, err error) {
	requestBody, err := json.Marshal(map[string]string{
		"client_id": ClientId,
		"client_secret": ClientSecret,
		"scope": Scope,
		"grant_type": "sso_cookie",
	})
	if err != nil {
		return nil, fmt.Errorf("Can't marhsal params %w: ", err)
	}

	req, err := http.NewRequest(
		"POST",
		AuthApi + "oauth/token",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("Can't create new request %w: ", err)
	}

	resp, err := p.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Can't execute request %w: ", err)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(tokens)
	if err != nil {
		return nil, fmt.Errorf("Can't decode request %w: ", err)
	}

	return tokens, nil
}