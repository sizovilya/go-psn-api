<p align="center"><img src="assets/gopher_ps_gamer.png" width="250"></p>

# go-psn-api(WIP)
A Playstation Network API wrapper written in Go.
##Read first
Corresponding to my research how PSN works you need several things to interact with Sony servers:  
- npsso - some secret, you need it to obtain access token
- client_id - identifier of client
- client_secret - secret  

To get them please follow steps below.  
###How to get npsso  
Fully described here - https://tusticles.com/psn-php/first_login.html
###How to get client_id and client_secret
- Go to https://account.sonyentertainmentnetwork.com/ and log in with your own credentials
- Open Chrome Dev Tools, go to Network tab and find `token` request, url - https://auth.api.sonyentertainmentnetwork.com/2.0/oauth/token  
<img src="assets/gopher_ps_gamer.png" width="250">