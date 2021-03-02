[![Go Reference](https://pkg.go.dev/badge/github.com/sizovilya/go-psn-api.svg)](https://pkg.go.dev/github.com/sizovilya/go-psn-api)
![build status](https://github.com/sizovilya/go-psn-api/actions/workflows/go.yml/badge.svg?branch=main)
![build status](https://github.com/sizovilya/go-psn-api/actions/workflows/golangci-lint.yml/badge.svg?branch=main)
<p align="center"><img src="assets/gopher_ps_gamer.png" width="250"></p>

# go-psn-api(WIP)
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

### How to get client_id and client_secret
- Go to https://account.sonyentertainmentnetwork.com/ and log in with your own credentials
- Open Chrome Dev Tools, go to Network tab and find `token` request, url - https://auth.api.sonyentertainmentnetwork.com/2.0/oauth/token  
  <img src="assets/screen.png" width="450">

### Functions at this moment 
- You can get user profile info
- You can get trophy titles


### Example    
```go
package main

import (
  "fmt"
  "github.com/sizovilya/go-psn-api"
)

func main() {
  lang := "ru" // full list here https://github.com/sizovilya/go-psn-api/blob/main/langs.go
  region := "ru" // full list here https://github.com/sizovilya/go-psn-api/blob/main/regions.go
  npsso := "your npsso"
  psnApi, err := psn.NewPsnApi(
    lang,
    region,
  )
  if err != nil {
    panic(err)
  }

  // This request will get access token and refresh token from Sony's servers
  err = psnApi.AuthWithNPSSO(npsso)
  if err != nil {
    panic(err)
  }

  // If you've obtained refresh token you may use it for next time
  // Next logins should be like this:
  // refreshToken, _ := psnApi.GetRefreshToken() // store refresh token somewhere for future logins by psnApi.AuthWithRefreshToken method
  err = psnApi.AuthWithRefreshToken("your token") // get new access token
  if err != nil {
    panic(err)
  }

  // How to get user's profile info
  profile, err := psnApi.GetProfileRequest("geeek_52rus")
  if err != nil {
    panic(err)
  }
  fmt.Print(profile)

  // How to get trophy titles
  trophyTitles, err := psnApi.GetTrophyTitles("geeek_52rus", 50, 0)
  if err != nil {
    panic(err)
  }
  fmt.Print(trophyTitles)
}

```

<p align="center"> ![gopher-dance](assets/gopher-dance.gif) </p>