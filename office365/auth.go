package office365

import (
	"encoding/json"
	"fmt"
	"github.com/rprakashg/go-o365api-explorer/sessionstate"
	"github.com/rprakashg/go-o365api-explorer/webconfig"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var OAuthEndPoints = oauth2.Endpoint{
	AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
	TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
}

var (
	resourceString string
	settings       webconfig.Settings
)

func init() {
	settings = webconfig.Settings{}
	settings.Load("webconfig.json")
	resourceString = fmt.Sprintf("?resource=%s", settings.UnifiedApiResourceId)

	OAuthEndPoints.AuthURL += resourceString
	OAuthEndPoints.TokenURL += resourceString

	log.Printf("In Init - Resource String %s", resourceString)
}

//this piece of middleware code checks to see if user has authorized app to access office 365 on behalf of user
func IsAuthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	s, _ := sessionstate.GlobalSessions.SessionStart(w, r)
	defer s.SessionRelease(w)

	if s.Get("token") == nil {
		//get the absolute URL of current request
		requestUrl := r.URL.Path

		//app needs to be authorized, redirect to authorize endpoint that will do the oauth dance with office 365
		http.Redirect(w, r, "/authorize?requestUrl="+requestUrl, http.StatusTemporaryRedirect)
	} else {
		//call the next handler in chain
		next(w, r)
	}
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	config := &oauth2.Config{
		ClientID:     settings.ClientId,
		ClientSecret: settings.ClientSecret,
		RedirectURL:  settings.RedirectUrl,
		Endpoint:     OAuthEndPoints,
	}
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	// redirect to consent page to obtain consent from user
	http.Redirect(w, r, url, http.StatusFound)
}

func HandleOauthCallBack(w http.ResponseWriter, r *http.Request) {

	config := &oauth2.Config{
		ClientID:     settings.ClientId,
		ClientSecret: settings.ClientSecret,
		RedirectURL:  settings.RedirectUrl,
		Endpoint:     OAuthEndPoints,
	}

	state := r.FormValue("state")
	if state != "state" {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", "state", state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//get authorization code
	code := r.URL.Query().Get("code")

	// Exchanging the authorization code for an access token
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//serialize token to json and store in session
	s, _ := sessionstate.GlobalSessions.SessionStart(w, r)
	defer s.SessionRelease(w)

	//serialize struct to json string and store it in session
	b, err := json.Marshal(&token)
	if err != nil {
		fmt.Println("Couldn't serialize struct to json")
		panic(err)
	}
	s.Set("token", string(b))

	//load user profile information and store it in session
	uinfo := UserInfo{}
	uinfo.Get(*token)
	b, err = json.Marshal(&uinfo)
	if err != nil {
		fmt.Println("Couldn't serialize struct to json")
		panic(err)
	}

	s.Set("Profile", string(b))

	//redirect user to redirectUrl in querystring
	redirectUrl := r.URL.Query().Get("redirectUrl")
	if len(redirectUrl) != 0 {
		http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	} else {
		//just redirect to homepage
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
