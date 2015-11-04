package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	/*"github.com/gorilla/mux"*/)

/*PermissionsAuthenticate function that returnes json for Authenticate method in yotpo-adds*/
func PermissionsAuthenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	dat, err := ioutil.ReadFile("./yotpo-facebook/static_responses/permissions.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		panic(err)
	}

	jsondata := map[string]interface{}{}
	if err := json.Unmarshal(dat, &jsondata); err != nil {
		panic(err)
	}
	fmt.Printf("%q", jsondata)
	json.NewEncoder(w).Encode(jsondata)

}

/*AccountsAuthenticate function that returnes json for Authenticate method in yotpo-adds*/
func AccountsAuthenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "accounts!")
	fields := r.FormValue("fields")
	fmt.Printf("\n%q\n", fields)
	fieldsarr := strings.Split(fields, ",")
	fmt.Fprintf(w, "\nURL params are: %q\n", fieldsarr)

}

/*MeAuthenticate function that returnes json for Authenticate method in yotpo-adds*/
func MeAuthenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "me?")
}
