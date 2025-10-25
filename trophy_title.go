package psn

import (
	"context"
	"fmt"
	"time"
)

const trophyTitleApi = "m.np.playstation.com/api/trophy/v1/users/"

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
// accountId: Use "me" for the authenticating account or a numeric accountId to query another user's titles.
func (c *Client) GetTrophyTitles(ctx context.Context, accountId string, limit, offset int) (*TrophyTitleResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	h["Accept-Language"] = c.lang

	var response TrophyTitleResponse
	err := c.get(
		ctx,
		fmt.Sprintf(
			"https://%s%s/trophyTitles?limit=%d&offset=%d",
			trophyTitleApi,
			accountId,
			limit,
			offset,
		),
		h,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get trophy titles: %w", err)
	}
	return &response, nil
}
