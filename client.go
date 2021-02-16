package psn

import (
	"fmt"
	"net/http"
)

type psn struct {
	http *http.Client
	lang string
	region string
	npsso string
	refreshToken string
	accessToken string
}

func NewPsnApi (lang, region, npsso, refreshToken, accessToken string) (*psn, error) {
	if !isContain(languages, lang) {
		return nil, fmt.Errorf("can't create psnapi: unsupported lang %s", lang)
	}
	if !isContain(regions, region) {
		return nil, fmt.Errorf("can't create psnapi: unsupported region %s", region)
	}
	return &psn{
		http:         &http.Client{},
		lang:         lang,
		region:       region,
		npsso:        npsso,
		refreshToken: refreshToken,
		accessToken:  accessToken,
	}, nil
}

func (p *psn) Auth() error {
	tokens, err :=p.authRequest()
	if err != nil {
		return fmt.Errorf("can't do auth request %w: ", err)
	}
	p.accessToken = tokens.AccessToken
	p.refreshToken = tokens.RefreshToken
	return nil
}

func (p *psn) RefreshTokens() error {
	tokens, err :=p.refreshTokens()
	if err != nil {
		return fmt.Errorf("can't refresh tokens %w: ", err)
	}
	fmt.Println(tokens)
	return nil
}