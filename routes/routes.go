package routes

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Init(config *mysql.Config) {
	db, err := sql.Open("mysql", config.FormatDSN())

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Error")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func Routers() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}/{var1}/{var2}/", returnSingleArticle)
	return myRouter
}
