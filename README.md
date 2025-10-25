[![Go Reference](https://pkg.go.dev/badge/github.com/sizovilya/go-psn-api.svg)](https://pkg.go.dev/github.com/sizovilya/go-psn-api)
![build status](https://github.com/sizovilya/go-psn-api/actions/workflows/go.yml/badge.svg?branch=main)
![build status](https://github.com/sizovilya/go-psn-api/actions/workflows/golangci-lint.yml/badge.svg?branch=main)
<p align="center"><img src="assets/gopher_ps_gamer.png" width="250"></p>

# go-psn-api
A Playstation Network API wrapper written in Go.
## Read first
Corresponding to my research how PSN works you need npsso to interact with Sony servers.
Instructions how to get it below.  
### How to get npsso  
Fully described here - https://tusticles.com/psn-php/first_login.html
<details>
<summary>
If link above doesn't work
</summary>

Copy this js code:   
```javascript
(function(open) {
    XMLHttpRequest.prototype.open = function(method, url, async, user, pass) {

        this.addEventListener("readystatechange", function() {
            if (this.readyState == XMLHttpRequest.DONE) {
                let response = JSON.parse(this.responseText);

                if (response && "npsso" in response) {
                    console.log('found npsso', response.npsso);
                }
            }
        }, false);

        open.call(this, method, url, async, user, pass);
    };

    window.onbeforeunload = function(){
        return 'Are you sure you want to leave?';
    };

})(XMLHttpRequest.prototype.open);
```
 - Navigate to https://account.sonyentertainmentnetwork.com/ in your browser and open your browser’s developer console
 - Paste the above Javascript into the console and then login.
 - After the login flow is completed, you should see a new log in the developer console that looks like: found npsso <64 character code>. Copy that 64 character code.
</details>

### Functionality
- Get user profile information
- Get trophy titles for a user
- Get trophy groups for a specific title
- Get trophies for a specific title and group

### Example
```go
package main

import (
	"context"
	"fmt"
	"github.com/sizovilya/go-psn-api"
)

func main() {
	ctx := context.Background()
	opts := &psn.Options{
		Lang:   "ru", // See https://github.com/sizovilya/go-psn-api/blob/main/langs.go
		Region: "ru", // See https://github.com/sizovilya/go-psn-api/blob/main/regions.go
		Npsso:  "<your_npsso_code>",
	}

	client, err := psn.NewClient(opts)
	if err != nil {
		panic(err)
	}

	// Authenticate to get access and refresh tokens.
	err = client.AuthWithNPSSO(ctx, opts.Npsso)
	if err != nil {
		panic(err)
	}

	// You can also use a refresh token for subsequent authentications.
	// refreshToken, _ := client.RefreshToken()
	// err = client.AuthWithRefreshToken(ctx, refreshToken)
	// if err != nil {
	// 	panic(err)
	// }

	// Get user profile information by Online ID (username).
	profile, err := client.GetProfile(ctx, "your_psn_username")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Profile: %+v\n", profile)

	// Get trophy titles for your account.
	// Use "me" for your own account, or a numeric accountId for other users.
	trophyTitles, err := client.GetTrophyTitles(ctx, "me", 50, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Trophy Titles: %+v\n", trophyTitles)

	// Get trophy groups for a specific title.
	if len(trophyTitles.TrophyTitles) > 0 {
		trophyTitleID := trophyTitles.TrophyTitles[0].NpCommunicationID
		trophyGroups, err := client.GetTrophyGroups(ctx, trophyTitleID, "me")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Trophy Groups: %+v\n", trophyGroups)
	}

	// Get trophies for a specific title and group.
	// Parameters: npCommunicationId, trophyGroupId (e.g., "default" or "001"), accountId
	trophies, err := client.GetTrophies(ctx, "NPWR20188_00", "default", "me")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Trophies: %+v\n", trophies)
}
```

### Important Notes

**API Changes (v2 API Migration):**
- `GetTrophyTitles()` now requires an `accountId` parameter instead of `username`. Use `"me"` for your own account.
- `GetTrophies()` now uses the updated PSN API endpoint (`m.np.playstation.com/api/trophy/v1`).
- The API no longer uses region-specific endpoints for trophy operations.

**Account ID vs Online ID:**
- Some methods accept `onlineId` (username like "VaultTec_Trading")
- Trophy-related methods require `accountId` (use `"me"` for your account or a numeric account ID for others)
This project highly inspired by https://github.com/Tustin/psn-php. Some useful things like auth headers and params found in `Tustin/psn-php`. 
<p align="center"> <img src="assets/gopher-dance.gif"> </p>
