package office365

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/napping"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type UserInfo struct {
	Accountenabled    bool     `json:"accountEnabled"`
	City              string   `json:"city"`
	Companyname       string   `json:"companyName"`
	Country           string   `json:"country"`
	Department        string   `json:"department"`
	Displayname       string   `json:"displayName"`
	Givenname         string   `json:"givenName"`
	Jobtitle          string   `json:"jobTitle"`
	Mail              string   `json:"mail"`
	Mailnickname      string   `json:"mailNickname"`
	Mobile            string   `json:"mobile"`
	Othermails        []string `json:"otherMails"`
	Postalcode        string   `json:"postalCode"`
	Preferredlanguage string   `json:"preferredLanguage"`
	Sipproxyaddress   string   `json:"sipProxyAddress"`
	State             string   `json:"state"`
	Streetaddress     string   `json:"streetAddress"`
	Surname           string   `json:"surname"`
	Telephonenumber   string   `json:"telephoneNumber"`
	Usagelocation     string   `json:"usageLocation"`
	Userprincipalname string   `json:"userPrincipalName"`
}

func (u *UserInfo) Get(token oauth2.Token) {
	ns := napping.Session{}
	ns.Header = &http.Header{}
	ns.Header.Add("Authorization", "Bearer "+token.AccessToken)
	url := settings.UnifiedApiEndpointUrl + "/me"
	resp, err := ns.Get(url, nil, &u, nil)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Status() == 200 {
		fmt.Println("Loaded profile for :", u.Displayname)
	} else {
		fmt.Println("Bad response status from office365 unified api")
		fmt.Println("Raw Response Text : %s", resp.RawText())
	}
}

func Get(url string, token oauth2.Token) string {
	ns := napping.Session{}
	ns.Header = &http.Header{}
	ns.Header.Add("Authorization", "Bearer "+token.AccessToken)
	resp, err := ns.Get(url, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status : %i", resp.Status())
	fmt.Println("Server response %s", resp.RawText())

	dst := new(bytes.Buffer)
	src := []byte(resp.RawText())

	json.Indent(dst, src, "", "\t")

	return dst.String()
}
