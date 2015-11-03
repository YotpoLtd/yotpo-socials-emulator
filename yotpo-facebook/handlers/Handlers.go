package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	/*"github.com/gorilla/mux"*/)

/*MeDataAccount ACCOUNTS DATA*/
type MeDataAccount struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

/*MeDataAccounts ccounts data array*/
type MeDataAccounts []MeDataAccount

/*MeDataPermission permisiins data */
type MeDataPermission struct {
	Permission string `json:"permission"`
	Status     string `json:"status"`
}

/*MeDataPermissions permissions array*/
type MeDataPermissions []MeDataPermission

/*MeData is an owner data*/
type MeData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

/*MeDatas is an owner data*/
type MeDatas []MeData

/*MePermission is /permission json responce structure*/
type MePermission struct {
	Data MeDataPermissions `json:"data"`
}

/*Authenticate function that returnes json for Authenticate method in yotpo-adds*/
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if strings.Contains(r.URL.Path, "permissions") {
		me := MePermission{
			Data: MeDataPermissions{
				MeDataPermission{Permission: "user_friends", Status: "granted"},
				MeDataPermission{Permission: "email", Status: "granted"},
				MeDataPermission{Permission: "ads_management", Status: "granted"},
				MeDataPermission{Permission: "ads_read", Status: "granted"},
				MeDataPermission{Permission: "manage_pages", Status: "granted"},
				MeDataPermission{Permission: "publish_pages", Status: "granted"},
				MeDataPermission{Permission: "public_profile", Status: "granted"},
			},
		}
		json.NewEncoder(w).Encode(me)
	}
	if strings.Contains(r.URL.Path, "accounts") {
		fmt.Fprintln(w, "accounts!")
	}
	if strings.Contains(r.URL.Path, "me?") {
		fmt.Fprintln(w, "me?")
	}

}
