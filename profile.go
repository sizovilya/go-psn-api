package psn

import (
	"context"
	"fmt"
)

const usersApi = "-prof.np.community.playstation.net/userProfile/v1/users/"

type AvatarUrls struct {
	Size      string `json:"size"`
	AvatarURL string `json:"avatarUrl"`
}
type EarnedTrophies struct {
	Platinum int `json:"platinum"`
	Gold     int `json:"gold"`
	Silver   int `json:"silver"`
	Bronze   int `json:"bronze"`
}
type TrophySummary struct {
	Level          int            `json:"level"`
	Progress       int            `json:"progress"`
	EarnedTrophies EarnedTrophies `json:"earnedTrophies"`
}
type PersonalDetail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Presences struct {
}
type ConsoleAvailability struct {
	AvailabilityStatus string `json:"availabilityStatus"`
}
type Profile struct {
	OnlineID              string              `json:"onlineId"`
	NpID                  string              `json:"npId"`
	AvatarUrls            []AvatarUrls        `json:"avatarUrls"`
	Plus                  int                 `json:"plus"`
	AboutMe               string              `json:"aboutMe"`
	LanguagesUsed         []string            `json:"languagesUsed"`
	TrophySummary         TrophySummary       `json:"trophySummary"`
	IsOfficiallyVerified  bool                `json:"isOfficiallyVerified"`
	PersonalDetail        PersonalDetail      `json:"personalDetail"`
	PersonalDetailSharing string              `json:"personalDetailSharing"`
	PrimaryOnlineStatus   string              `json:"primaryOnlineStatus"`
	Presences             []Presences         `json:"presences"`
	FriendRelation        string              `json:"friendRelation"`
	Blocking              bool                `json:"blocking"`
	MutualFriendsCount    int                 `json:"mutualFriendsCount"`
	Following             bool                `json:"following"`
	FollowerCount         int                 `json:"followerCount"`
	ConsoleAvailability   ConsoleAvailability `json:"consoleAvailability"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

// Method retrieves user profile info by PSN id
func (p *psn) GetProfileRequest(ctx context.Context, name string) (profile *Profile, err error) {
	var h = headers{}
	h["authorization"] = fmt.Sprintf("Bearer %s", p.accessToken)

	userResponse := &ProfileResponse{}
	err = p.get(
		ctx,
		fmt.Sprintf(
			"https://%s%s%s/profile2?fields=onlineId,aboutMe,consoleAvailability,languagesUsed,avatarUrls,personalDetail,personalDetail(@default,profilePictureUrls),primaryOnlineStatus,trophySummary(level,progress,earnedTrophies),plus,isOfficiallyVerified,friendRelation,personalDetailSharing,presences(@default,platform),npId,blocking,following,currentOnlineId,displayableOldOnlineId,mutualFriendsCount,followerCount&profilePictureSizes=s,m,l&avatarSizes=s,m,l&languagesUsedLanguageSet=set4",
			p.region,
			usersApi,
			name,
		),
		h,
		&userResponse,
	)
	if err != nil {
		return nil, fmt.Errorf("can't do GET request: %w", err)
	}
	return &userResponse.Profile, nil
}
