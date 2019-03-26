package main

import (
	"log"
	"net/http"

	"./routes"
)

func main() {
	myRouter := routes.Routers()
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
