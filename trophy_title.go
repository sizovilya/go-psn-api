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

// GetTrophyTitles retrieves a user's trophy titles.
func (c *Client) GetTrophyTitles(ctx context.Context, username string, limit, offset int) (*TrophyTitleResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	var response TrophyTitleResponse
	err := c.get(
		ctx,
		fmt.Sprintf(
			"https://%s%sfields=@default,trophyTitleSmallIconUrl&platform=PS3,PS4,PSVITA&limit=%d&offset=%d&comparedUser=%s&npLanguage=%s",
			c.region,
			trophyTitleApi,
			limit,
			offset,
			username,
			c.lang,
		),
		h,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get trophy titles: %w", err)
	}
	return &response, nil
}