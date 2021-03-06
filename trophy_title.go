package psn

import (
	"context"
	"fmt"
	"time"
)

const trophyTitleApi = "-tpy.np.community.playstation.net/trophy/v1/trophyTitles?"

type TrophyTitleResponse struct {
	TotalResults int `json:"totalResults"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	TrophyTitles []struct {
		NpCommunicationID       string `json:"npCommunicationId"`
		TrophyTitleName         string `json:"trophyTitleName"`
		TrophyTitleDetail       string `json:"trophyTitleDetail"`
		TrophyTitleIconURL      string `json:"trophyTitleIconUrl"`
		TrophyTitleSmallIconURL string `json:"trophyTitleSmallIconUrl"`
		TrophyTitlePlatfrom     string `json:"trophyTitlePlatfrom"` // typo in Sony's response
		HasTrophyGroups         bool   `json:"hasTrophyGroups"`
		DefinedTrophies         struct {
			Bronze   int `json:"bronze"`
			Silver   int `json:"silver"`
			Gold     int `json:"gold"`
			Platinum int `json:"platinum"`
		} `json:"definedTrophies"`
		ComparedUser struct {
			OnlineID       string `json:"onlineId"`
			Progress       int    `json:"progress"`
			EarnedTrophies struct {
				Bronze   int `json:"bronze"`
				Silver   int `json:"silver"`
				Gold     int `json:"gold"`
				Platinum int `json:"platinum"`
			} `json:"earnedTrophies"`
			LastUpdateDate time.Time `json:"lastUpdateDate"`
		} `json:"comparedUser"`
		FromUser struct {
			OnlineID       string `json:"onlineId"`
			Progress       int    `json:"progress"`
			EarnedTrophies struct {
				Bronze   int `json:"bronze"`
				Silver   int `json:"silver"`
				Gold     int `json:"gold"`
				Platinum int `json:"platinum"`
			} `json:"earnedTrophies"`
			HiddenFlag     bool      `json:"hiddenFlag"`
			LastUpdateDate time.Time `json:"lastUpdateDate"`
		} `json:"fromUser,omitempty"`
	} `json:"trophyTitles"`
}

// Method retrieves user's trophy titles
func (p *psn) GetTrophyTitles(ctx context.Context, username string, limit, offset int32) (*TrophyTitleResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	response := TrophyTitleResponse{}
	err := p.get(
		ctx,
		fmt.Sprintf(
			"https://%s%sfields=@default,trophyTitleSmallIconUrl&platform=PS3,PS4,PSVITA&limit=%d&offset=%d&comparedUser=%s&npLanguage=%s",
			p.region,
			trophyTitleApi,
			limit,
			offset,
			username,
			p.lang,
		),
		h,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &response, nil
}
