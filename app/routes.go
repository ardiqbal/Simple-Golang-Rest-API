package app

import (
	"fmt"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/add-class", CreateNewClass)
	myRouter.HandleFunc("/add-attendance", AddAttendance)
	myRouter.HandleFunc("/classes", GetAllClasses)
	myRouter.HandleFunc("/class/{id}", GetClassDetail)
	fmt.Println("Listening on Port 8000")
	return myRouter
}
