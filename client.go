package psn

import (
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

func NewPsnApi (lang, region, npsso, refreshToken, accessToken string) *psn {
	return &psn{
		http:         &http.Client{},
		lang:         lang,
		region:       region,
		npsso:        npsso,
		refreshToken: refreshToken,
		accessToken:  accessToken,
	}
}

func (p *psn) Auth() error {
	return nil
}