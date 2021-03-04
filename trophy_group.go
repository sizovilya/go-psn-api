package psn

import (
	"context"
	"fmt"
	"time"
)

const trophyGroupApi = "-tpy.np.community.playstation.net/trophy/v1/trophyTitles/"

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

// Method retrieves user's trophy groups
func (p *psn) GetTrophyGroups(ctx context.Context, trophyTitleId, username string) (profile *TrophyGroupResponse, err error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)
	h["Accept"] = "*/*"
	h["Accept-Encoding"] = "gzip, deflate, br"

	trophyTitleResponse := TrophyGroupResponse{}
	err = p.get(
		ctx,
		fmt.Sprintf(
			"https://%s%s%s/trophyGroups?fields=@default,trophyGroupSmallIconUrl&comparedUser=%s&npLanguage=%s",
			p.region,
			trophyGroupApi,
			trophyTitleId,
			username,
			p.lang,
		),
		h,
		&trophyTitleResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &trophyTitleResponse, nil
}
