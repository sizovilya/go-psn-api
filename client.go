package psn

import (
	"fmt"
	"net/http"
)

type psn struct {
	http           *http.Client
	lang           string
	region         string
	npsso          string
	accessToken    string
	refreshToken   string
	accessExpired  int32
	refreshExpired int32
}

// Creates new psn api
func NewPsnApi(lang, region, npsso string) (*psn, error) {
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
		http:   &http.Client{},
		lang:   lang,
		region: region,
		npsso:  npsso,
	}, nil
}

// Setter for lang
func (p *psn) SetLang(lang string) error {
	if !isContain(languages, lang) {
		return fmt.Errorf("unsupported lang %s", lang)
	}
	p.lang = lang
	return nil
}

// Getter for lang
func (p *psn) GetLang() string {
	return p.lang
}

// Setter for region
func (p *psn) SetRegion(region string) error {
	if !isContain(regions, region) {
		return fmt.Errorf("cunsupported region %s", region)
	}
	p.region = region
	return nil
}

// Getter for region
func (p *psn) GetRegion() string {
	return p.region
}

// Setter for npsso
func (p *psn) SetNPSSO(npsso string) error {
	if npsso == "" {
		return fmt.Errorf("npsso is empty")
	}
	p.npsso = npsso
	return nil
}

// Getter for npsso
func (p *psn) GetNPSSO() string {
	return p.npsso
}

// Setter for access token
func (p *psn) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token is empty")
	}
	p.accessToken = accessToken
	return nil
}

// Getter for access token
func (p *psn) GetAccessToken() (string, int32) {
	return p.accessToken, p.accessExpired
}

// Getter for refresh token
func (p *psn) SetRefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}
	p.refreshToken = refreshToken
	return nil
}

// Getter for refresh token
func (p *psn) GetRefreshToken() (string, int32) {
	return p.refreshToken, p.refreshExpired
}
