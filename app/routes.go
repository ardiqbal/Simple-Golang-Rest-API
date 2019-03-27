package app

import (
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/add-class", CreateNewClass)
	return myRouter
}
