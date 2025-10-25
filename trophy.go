package psn

import (
	"context"
	"fmt"
)

const trophiesApi = "m.np.playstation.com/api/trophy/v1/users/"

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
// accountId: Use "me" for the authenticating account or a numeric accountId to query another user's trophies.
// trophyGroupId: Use "all" to get all trophies, "default" for base game, or specific group IDs like "001", "002", etc.
// Note: For PS5/PC games, no npServiceName is needed. For PS3/PS4/Vita games, use GetTrophiesWithServiceName with npServiceName="trophy".
func (c *Client) GetTrophies(ctx context.Context, trophyTitleId, trophyGroupId, accountId string) (*TrophiesResponse, error) {
	return c.GetTrophiesWithServiceName(ctx, trophyTitleId, trophyGroupId, accountId, "")
}

// GetTrophiesWithServiceName retrieves a user's trophies for a specific title and group with an explicit npServiceName.
// accountId: Use "me" for the authenticating account or a numeric accountId to query another user's trophies.
// trophyGroupId: Use "all" to get all trophies, "default" for base game, or specific group IDs like "001", "002", etc.
// npServiceName: Use "trophy" for PS3/PS4/Vita, "trophy2" for PS5/PC, or "" (empty) to omit the parameter.
func (c *Client) GetTrophiesWithServiceName(ctx context.Context, trophyTitleId, trophyGroupId, accountId, npServiceName string) (*TrophiesResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	h["Accept-Language"] = c.lang

	var trophiesResponse TrophiesResponse
	url := fmt.Sprintf(
		"https://%s%s/npCommunicationIds/%s/trophyGroups/%s/trophies",
		trophiesApi,
		accountId,
		trophyTitleId,
		trophyGroupId,
	)

	if npServiceName != "" {
		url += "?npServiceName=" + npServiceName
	}

	err := c.get(ctx, url, h, &trophiesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to get trophies: %w", err)
	}
	return &trophiesResponse, nil
}
