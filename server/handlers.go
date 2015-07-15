package server

import (
	"encoding/json"
	"fmt"
	"github.com/rprakashg/go-o365api-explorer/office365"
	"github.com/rprakashg/go-o365api-explorer/sessionstate"
	"golang.org/x/oauth2"
	"html/template"
	"net/http"
	"time"
)

type Link struct {
	Class string
	Href  string
	Text  string
}

type HomeData struct {
	EndPointUri     string
	Result          string
	RequestDuration string
}

const (
	UrlRequestKey   string = "urlRequest"
	RestEndPointKey string = "restEndPoint"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//parse query string
	qs := r.URL.Query()

	s, _ := sessionstate.GlobalSessions.SessionStart(w, r)
	defer s.SessionRelease(w)

	uinfo := office365.UserInfo{}
	hl := Link{Class: "fa fa-sign-in fa-fw", Href: "/authorize", Text: "Sign in"}

	p := s.Get("Profile")
	if p != nil {
		if profileString, ok := p.(string); ok {
			err := json.Unmarshal([]byte(profileString), &uinfo)
			if err != nil {
				fmt.Println("Error deserializing profile")
				panic(err)
			}
		}
		hl.Class = "fa fa-sign-out fa-fw"
		hl.Href = "/signout"
		hl.Text = "Signout"

	}
	var resultString string
	var requestDuration string
	var endPointUri string = "https://graph.microsoft.com/beta/"

	//check a rest query needs to be send to Office 365 unified API
	if len(qs[UrlRequestKey]) > 0 && len(qs[RestEndPointKey]) > 0 {
		token := oauth2.Token{}
		t := s.Get("token")

		if t != nil {
			if tokenString, ok := t.(string); ok {
				err := json.Unmarshal([]byte(tokenString), &token)
				if err != nil {
					fmt.Println("Error deserializing token")
					panic(err)
				}
			}
		}
		endPointUri = qs[RestEndPointKey][0]
		t0 := time.Now()
		resultString = office365.Get(endPointUri, token)
		t1 := time.Now()
		requestDuration = fmt.Sprintf("%v", t1.Sub(t0))
	}

	//set up parameters for home page
	d := HomeData{EndPointUri: endPointUri, Result: resultString, RequestDuration: requestDuration}
	params := map[string]interface{}{"Data": d}

	//render page content
	markup := executeTemplate("home", params)

	//merge with layout content
	layoutParams := map[string]interface{}{"PageContent": template.HTML(string(markup)), "User": uinfo, "Link": hl}
	markup = executeTemplate("layout", layoutParams)

	//write the output to response
	fmt.Fprintf(w, string(markup))
}
