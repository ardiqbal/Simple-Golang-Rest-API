package main

import (
	"log"
	"net/http"

	"github.com/ardiqbal/Simple-Golang-Rest-API/app"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", app.Routers()))
}
