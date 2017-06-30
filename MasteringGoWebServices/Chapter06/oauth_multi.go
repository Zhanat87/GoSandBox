package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
)

var (

)

type OauthService struct {
	clientID string
	clientSecret string
	scope string
	redirectURL string
	authURL string
	tokenURL string
	requestURL string
	code string
}



func main() {

	OauthServices := map[string] OauthService{}
	
	OauthServices["facebook"] = OauthService {
		clientID:  "1435244946724353",
		clientSecret: "b8d4783df9a4011728e92145b73cdb40",
		scope: "",
		redirectURL: "http://www.mastergoco.com",
		authURL: "https://www.facebook.com/dialog/oauth",
		tokenURL: "https://graph.facebook.com/oauth/access_token",
		requestURL: "https://graph.facebook.com/me",
		code: "",
	}
	OauthServices["google"] = OauthService {
		clientID:  "576771593450-kk4fccoumsr25lu6m65le7dqu24r60cn.apps.googleusercontent.com",
		clientSecret: "3kJI4CHi73cOclvx-dKiDm5D",
		scope: "https://www.googleapis.com/auth/plus.login",
		redirectURL: "http://www.mastergoco.com",
		authURL: "https://accounts.google.com/o/oauth2/auth",
		tokenURL: "https://accounts.google.com/o/oauth2/token",
		requestURL: "https://graph.facebook.com/me",
		code: "",
	}	

	for key, value := range OauthServices {
		fmt.Println("New Service",key)
		fmt.Println(value)

		oauthConnection := &oauth.Config{
			ClientId:     OauthServices[key].clientID,
			ClientSecret: OauthServices[key].clientSecret,
			RedirectURL:  OauthServices[key].redirectURL,
			Scope:        OauthServices[key].scope,
			AuthURL:      OauthServices[key].authURL,
			TokenURL:     OauthServices[key].tokenURL,
		}

		url := oauthConnection.AuthCodeURL("")
		fmt.Print("Visit this URL to get a code, then run again with -code=YOUR_CODE\n\n")
		fmt.Println(url)


	}


}
