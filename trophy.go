package psn

import "fmt"

// const trophyTitleApi = "-tpy.np.community.playstation.net/trophy/v1/trophyTitles?"

// Method retrieves user's trophy titles
func (p *psn) GetTrophyTitles(name string) (profile *Profile, err error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)
	h["accept-language"] = "ru-RU"
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	userResponse := &ProfileResponse{}
	//err = p.get(
	//	fmt.Sprintf(
	//		"https://%s-tpy.np.community.playstation.net/trophy/v1/trophyTitles?fields=@default&npLanguage=%s&iconSize=m&platform=PS3,PSVITA,PS4&offset=0&limit=20&comparedUser=%s",
	//		//"https://%s%sfields=@default,trophyTitleSmallIconUrl&platform=PS3,PS4,PSVITA&limit=12&offset=0&comparedUser=%s&npLanguage=ru",
	//		p.region,
	//		p.lang,
	//		//trophyTitleApi,
	//		name,
	//	),
	//	h,
	//	&userResponse,
	//)
	// check scope
	err = p.get(
		"https://ru-tpy.np.community.playstation.net/trophy/v1/trophyTitles?fields=%40default%2CtrophyTitleSmallIconUrl&platform=PS3%2CPS4%2CPSVITA&limit=12&offset=0&npLanguage=ru",
		h,
		&userResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &userResponse.Profile, nil
}
