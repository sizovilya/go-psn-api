package psn

import (
	"fmt"
	"net/http"
)

type psn struct {
	http         *http.Client
	lang         string
	region       string
	npsso        string
	clientId     string
	clientSecret string
	accessToken  string
}

// Creates new psn api
func NewPsnApi(lang, region, npsso, clientId, clientSecret string) (*psn, error) {
	if !isContain(languages, lang) {
		return nil, fmt.Errorf("can't create psnapi: unsupported lang %s", lang)
	}
	if !isContain(regions, region) {
		return nil, fmt.Errorf("can't create psnapi: unsupported region %s", region)
	}
	if npsso == "" {
		return nil, fmt.Errorf("can't create psnapi: npsso is empty")
	}
	if clientId == "" {
		return nil, fmt.Errorf("can't create psnapi: clientId is empty")
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("can't create psnapi: clientSecret is empty")
	}
	return &psn{
		http:         &http.Client{},
		lang:         lang,
		region:       region,
		npsso:        npsso,
		clientId:     clientId,
		clientSecret: clientSecret,
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

// Setter for client id
func (p *psn) SetClientId(clientId string) error {
	if clientId == "" {
		return fmt.Errorf("clientId is empty")
	}
	p.clientId = clientId
	return nil
}

// Getter for client id
func (p *psn) GetClientId() string {
	return p.clientId
}

// Setter for secret
func (p *psn) SetClientSecret(clientSecret string) error {
	if clientSecret == "" {
		return fmt.Errorf("clientSecret is empty")
	}
	p.clientSecret = clientSecret
	return nil
}

// Getter for secret
func (p *psn) GetClientSecret() string {
	return p.clientSecret
}

// Method makes auth request to Sony's server and retrieves access token
func (p *psn) Auth() error {
	tokens, err := p.authRequest()
	if err != nil {
		return fmt.Errorf("can't do auth request: %w", err)
	}
	p.accessToken = tokens.AccessToken
	return nil
}
