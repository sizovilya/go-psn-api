package psn

import (
	"context"
	"fmt"
)

const trophiesApi = "-tpy.np.community.playstation.net/trophy/v1/trophyTitles/"

type TrophiesResponse struct {
	Trophies []struct {
		TrophyID           int    `json:"trophyId"`
		TrophyHidden       bool   `json:"trophyHidden"`
		TrophyType         string `json:"trophyType"`
		TrophyName         string `json:"trophyName"`
		TrophyDetail       string `json:"trophyDetail"`
		TrophyIconURL      string `json:"trophyIconUrl"`
		TrophySmallIconURL string `json:"trophySmallIconUrl"`
		TrophyRare         int    `json:"trophyRare"`
		TrophyEarnedRate   string `json:"trophyEarnedRate"`
		FromUser           struct {
			OnlineID string `json:"onlineId"`
			Earned   bool   `json:"earned"`
		} `json:"fromUser,omitempty"`
	} `json:"trophies"`
}

// Method retrieves user's trophies
func (p *psn) GetTrophies(ctx context.Context, trophyTitleId, trophyGroupId, username string) (*TrophiesResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	trophiesResponse := TrophiesResponse{}
	err := p.get(
		ctx,
		fmt.Sprintf(
			"https://%s%s%s/trophyGroups/%s/trophies?fields=@default,trophyRare,trophyEarnedRate,trophySmallIconUrl&visibleType=1&comparedUser=%s&npLanguage=%s",
			p.region,
			trophiesApi,
			trophyTitleId,
			trophyGroupId,
			username,
			p.lang,
		),
		h,
		&trophiesResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &trophiesResponse, nil
}
