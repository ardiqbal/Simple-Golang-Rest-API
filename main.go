package main

import (
	"./dbconfig"
	"./routes"
)

func main() {
	// myRouter := routes.Routers()
	// log.Fatal(http.ListenAndServe(":10000", myRouter))

	dbConfig := dbconfig.Config()
	routes.Init(dbConfig)
}
