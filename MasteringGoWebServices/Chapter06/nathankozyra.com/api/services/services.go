package service

import
(
	"fmt"
	"code.google.com/p/goauth2/oauth"
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

var OauthServices map[string] OauthService

func InitServices() {

	OauthServices = map[string] OauthService{}	

	OauthServices["facebook"] = OauthService {
		clientID:  "1435244946724353",
		clientSecret: "b8d4783df9a4011728e92145b73cdb40",
		scope: "",
		redirectURL: "http://www.mastergoco.com/connect/facebook",
		authURL: "https://www.facebook.com/dialog/oauth",
		tokenURL: "https://graph.facebook.com/oauth/access_token",
		requestURL: "https://graph.facebook.com/me",
		code: "",
	}
	OauthServices["google"] = OauthService {
		clientID:  "576771593450-kk4fccoumsr25lu6m65le7dqu24r60cn.apps.googleusercontent.com",
		clientSecret: "3kJI4CHi73cOclvx-dKiDm5D",
		scope: "https://www.googleapis.com/auth/plus.login",
		redirectURL: "http://www.mastergoco.com/connect/google",
		authURL: "https://accounts.google.com/o/oauth2/auth",
		tokenURL: "https://accounts.google.com/o/oauth2/token",
		requestURL: "https://graph.facebook.com/me",
		code: "",
	}	
}

func PostMessage(service string, authCode string, scope string) (bool) {
	OauthServices[service].scope = scope
	token, err = transport.Exchange(*authCode)
    if err != nil {
            log.Fatal("Exchange:", err)
    }	

}

func GetAccessTokenURL(service string, scope string) (string) {

		oauthConnection := &oauth.Config{
			ClientId:     OauthServices[service].clientID,
			ClientSecret: OauthServices[service].clientSecret,
			RedirectURL:  OauthServices[service].redirectURL,
			Scope:        OauthServices[service].scope,
			AuthURL:      OauthServices[service].authURL,
			TokenURL:     OauthServices[service].tokenURL,
		}

		fmt.Println(OauthServices[service])

		url := oauthConnection.AuthCodeURL("")
		return url

}