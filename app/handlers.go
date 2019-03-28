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
	fmt.Println("Endpoint Hit: CreateNewClass")
	newClass := models.Class{}
	_ = json.NewDecoder(r.Body).Decode(&newClass)

	db := database.Connect()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO class (id, class_name, class_time, room, created_at, updated_at) values (UUID(), ?, TIMESTAMP(?), ?, NOW(), NOW())")
	_, err := stmt.Exec(newClass.ClassName, newClass.ClassTime, newClass.Room)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		json.NewEncoder(w).Encode(util.PostHttpResponse{Success: true, Message: []string{"OK"}})
	}
}

func AddAttendance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: AddAttendance")

	newAttendance := models.Attendance{}
	_ = json.NewDecoder(r.Body).Decode(&newAttendance)

	db := database.Connect()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO attendance (attendance_id, class_id, student_id, attend_time) values (UUID(), (SELECT id FROM class WHERE id = ?), (SELECT id FROM student WHERE id = ?), NOW())")
	_, err := stmt.Exec(newAttendance.ClassId, newAttendance.StudentId)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		json.NewEncoder(w).Encode(util.PostHttpResponse{Success: true, Message: []string{"OK"}})
	}
}

func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllClasses")
	allClasses := models.Classes{}

	db := database.Connect()
	defer db.Close()
	results, err := db.Query("SELECT * FROM class")

	for results.Next() {
		var class models.Class
		_ = results.Scan(&class.Id, &class.ClassName, &class.ClassTime, &class.Room, &class.CreatedAt, &class.UpdatedAt)
		class.ClassTime = util.Format(class.ClassTime)
		class.CreatedAt = util.Format(class.CreatedAt)
		class.UpdatedAt = util.Format(class.UpdatedAt)
		allClasses = append(allClasses, class)
	}

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(util.GetHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		json.NewEncoder(w).Encode(util.GetHttpResponse{Success: true, Message: []string{"OK"}, Data: util.Data{allClasses}})
	}
}
