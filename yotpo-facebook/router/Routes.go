package router

import (
	"net/http"
	"yotpo-socials-emulator/yotpo-facebook/handlers"
)

/*Route structure*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*Routes array*/
type Routes []Route

var routes = Routes{
	Route{
		"Authenticate",
		"GET",
		"/me/permissions",
		handlers.PermissionsAuthenticate,
	},
	Route{
		"Authenticate",
		"GET",
		"/me",
		handlers.MeAuthenticate,
	},
	Route{
		"Authenticate",
		"GET",
		"/me/accounts",
		handlers.AccountsAuthenticate,
	},
}
