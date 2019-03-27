package app

import (
	"fmt"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/add-class", CreateNewClass)
	myRouter.HandleFunc("/classes", GetAllClasses)
	fmt.Println("Listening on Port 8000")
	return myRouter
}
