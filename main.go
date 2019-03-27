package main

import (
	"log"
	"net/http"

	"github.com/ardiqbal/Simple-Golang-Rest-API/app"
)

func main() {
	// dbConfig := dbconfig.Config()
	// app.Init(dbConfig)
	log.Fatal(http.ListenAndServe(":8000", app.Routers()))
}
