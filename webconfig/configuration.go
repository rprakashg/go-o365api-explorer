package webconfig

import (
	"encoding/json"
	"os"
)

type Settings struct {
	ClientId              string `json:"clientId"`
	ClientSecret          string `json:"clientSecret"`
	RedirectUrl           string `json:"redirectUrl"`
	Tenant                string `json:"tenant"`
	UnifiedApiEndpointUrl string `json:"unifiedApiEndpointUrl"`
	UnifiedApiResourceId  string `json:"unifiedApiResourceId"`
}

func (s *Settings) Load(filename string) {
	dir := os.Getenv("GOPATH") + "/src/github.com/rprakashg/go-o365api-explorer/webconfig/"
	//load app config from appconfig.json file
	file, err := os.Open(dir + filename)
	if err != nil {
		panic(err)
	}
	parser := json.NewDecoder(file)
	if err = parser.Decode(&s); err != nil {
		panic(err)
	}
}
