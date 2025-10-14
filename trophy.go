package psn

import (
	"context"
	"fmt"
)

const trophiesApi = ""

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

// GetTrophies retrieves a user's trophies for a specific title and group.
func (c *Client) GetTrophies(ctx context.Context, trophyTitleId, trophyGroupId, username string) (*TrophiesResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	var trophiesResponse TrophiesResponse
	err := c.get(
		ctx,
		fmt.Sprintf(
			"https://m.np.playstation.net/api/trophy/v1/users/%s/trophyTitles/%s/trophyGroups/%s/trophies?fields=@default,trophyRare,trophyEarnedRate,trophySmallIconUrl&visibleType=1&npLanguage=%s",
			username,
			trophyTitleId,
			trophyGroupId,
			c.lang,
		),
		h,
		&trophiesResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get trophies: %w", err)
	}
	return &trophiesResponse, nil
}