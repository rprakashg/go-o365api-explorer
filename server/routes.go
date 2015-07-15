package server

import (
	"github.com/rprakashg/go-o365api-explorer/office365"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		homeHandler,
	},
	Route{
		"Authorize",
		"GET",
		"/authorize",
		office365.Authorize,
	},
	Route{
		"Callback",
		"GET",
		"/callback",
		office365.HandleOauthCallBack,
	},
}
