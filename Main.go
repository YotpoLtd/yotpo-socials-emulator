package main

import (
	"log"
	"net/http"
	"yotpo-socials-emulator/yotpo-facebook/router"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
