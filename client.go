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
	if npsso == "" {
		return nil, fmt.Errorf("can't create psnapi: npsso is empty")
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

func (p *psn) SetLang(lang string) error {
	if !isContain(languages, lang) {
		return fmt.Errorf("unsupported lang %s", lang)
	}
	p.lang = lang
	return nil
}

func (p *psn) GetLang() string {
	return p.lang
}

func (p *psn) SetRegion(region string) error {
	if !isContain(regions, region) {
		return fmt.Errorf("cunsupported region %s", region)
	}
	p.region = region
	return nil
}

func (p *psn) GetRegion() string {
	return p.region
}

func (p *psn) SetNPSSO(npsso string) error {
	if npsso == "" {
		return fmt.Errorf("npsso is empty")
	}
	p.npsso = npsso
	return nil
}

func (p *psn) GetNPSSO() string {
	return p.npsso
}

func (p *psn) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("accessToken is empty")
	}
	p.accessToken = accessToken
	return nil
}

func (p *psn) GetAccessToken() string {
	return p.accessToken
}

func (p *psn) SetRefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("refreshToken is empty")
	}
	p.refreshToken = refreshToken
	return nil
}

func (p *psn) GetRefreshToken() string {
	return p.refreshToken
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
	p.accessToken = tokens.AccessToken
	p.refreshToken = tokens.RefreshToken
	return nil
}