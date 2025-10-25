package psn

import (
	"context"
	"fmt"
	"time"
)

const trophyGroupApi = "m.np.playstation.com/api/trophy/v1/users/"

type TrophyGroupResponse struct {
	TrophyTitleName     string `json:"trophyTitleName"`
	TrophyTitleDetail   string `json:"trophyTitleDetail"`
	TrophyTitleIconURL  string `json:"trophyTitleIconUrl"`
	TrophyTitlePlatfrom string `json:"trophyTitlePlatfrom"`
	DefinedTrophies     struct {
		Bronze   int `json:"bronze"`
		Silver   int `json:"silver"`
		Gold     int `json:"gold"`
		Platinum int `json:"platinum"`
	} `json:"definedTrophies"`
	TrophyGroups []struct {
		TrophyGroupID           string `json:"trophyGroupId"`
		TrophyGroupName         string `json:"trophyGroupName"`
		TrophyGroupDetail       string `json:"trophyGroupDetail"`
		TrophyGroupIconURL      string `json:"trophyGroupIconUrl"`
		TrophyGroupSmallIconURL string `json:"trophyGroupSmallIconUrl"`
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
	} `json:"trophyGroups"`
}

// GetTrophyGroups retrieves a user's trophy groups for a specific title.
// accountId: Use "me" for the authenticating account or a numeric accountId to query another user's trophy groups.
func (c *Client) GetTrophyGroups(ctx context.Context, trophyTitleId, accountId string) (*TrophyGroupResponse, error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	h["Accept-Language"] = c.lang

	var response TrophyGroupResponse
	err := c.get(
		ctx,
		fmt.Sprintf(
			"https://%s%s/npCommunicationIds/%s/trophyGroups",
			trophyGroupApi,
			accountId,
			trophyTitleId,
		),
		h,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get trophy groups: %w", err)
	}
	return &response, nil
}
