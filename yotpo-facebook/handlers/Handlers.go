package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	/*"github.com/gorilla/mux"*/)

/*Authenticate function that returnes json for Authenticate method in yotpo-adds*/
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if strings.Contains(r.URL.Path, "permissions") {

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
	if strings.Contains(r.URL.Path, "accounts") {
		fmt.Fprintln(w, "accounts!")
		fields := r.FormValue("fields")
		fmt.Printf("\n%q\n", fields)
		fieldsarr := strings.Split(fields, ",")
		fmt.Fprintf(w, "\nURL params are: %q\n", fieldsarr)
	}
	if strings.Contains(r.URL.Path, "me?") {
		fmt.Fprintln(w, "me?")
	}

}
