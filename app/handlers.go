package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ardiqbal/Simple-Golang-Rest-API/database"
	"github.com/ardiqbal/Simple-Golang-Rest-API/models"
	"github.com/ardiqbal/Simple-Golang-Rest-API/util"
)

func CreateNewClass(w http.ResponseWriter, r *http.Request) {
	newClass := models.Class{}
	_ = json.NewDecoder(r.Body).Decode(&newClass)

	db := database.Connect()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO class (id, class_name, class_time, room, created_at, updated_at) values (UUID(), ?, TIMESTAMP(?), ?, NOW(), NOW())")
	_, err := stmt.Exec(newClass.ClassName, newClass.ClassTime, newClass.Room)

	fmt.Println(newClass)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
	}

	fmt.Println("Endpoint Hit: CreateNewClass")

	json.NewEncoder(w).Encode(util.PostHttpResponse{Success: true, Message: []string{"OK"}})
}

func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	allClasses := models.Classes{}

	db := database.Connect()
	defer db.Close()
	results, err := db.Query("SELECT * FROM class")

	for results.Next() {
		var class models.Class
		err = results.Scan(&class.Id, &class.ClassName, &class.ClassTime, &class.Room, &class.CreatedAt, &class.UpdatedAt)
		if err != nil {
			json.NewEncoder(w).Encode(util.GetHttpResponse{Success: false, Message: []string{"FAILED"}})
		}
		class.ClassTime = util.Format(class.ClassTime)
		class.CreatedAt = util.Format(class.CreatedAt)
		class.UpdatedAt = util.Format(class.UpdatedAt)
		allClasses = append(allClasses, class)
	}

	json.NewEncoder(w).Encode(util.GetHttpResponse{Success: true, Message: []string{"OK"}, Data: util.Data{allClasses}})
}
