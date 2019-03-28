package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ardiqbal/Simple-Golang-Rest-API/database"
	"github.com/ardiqbal/Simple-Golang-Rest-API/models"
	"github.com/ardiqbal/Simple-Golang-Rest-API/util"
	"github.com/gorilla/mux"
)

func CreateNewClass(w http.ResponseWriter, r *http.Request) {
	log.Print("Endpoint Hit: CreateNewClass")
	newClass := models.Class{}
	_ = json.NewDecoder(r.Body).Decode(&newClass)

	db := database.Connect()
	defer db.Close()

	stmt1, _ := db.Prepare("INSERT INTO class (id, class_name, class_time, room, created_at, updated_at) values (UUID(), ?, TIMESTAMP(?), ?, NOW(), NOW())")
	_, err1 := stmt1.Exec(newClass.ClassName, newClass.ClassTime, newClass.Room)

	results, _ := db.Query("SELECT id FROM class WHERE class_name = ?", newClass.ClassName)

	if err1 != nil {
		log.Print(err1.Error())
		respondError(w, http.StatusInternalServerError, util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		for results.Next() {
			_ = results.Scan(&newClass.Id)
			stmt2, _ := db.Prepare("INSERT INTO attendance (attendance_id, class_id, attend_time) values (UUID(), (SELECT id FROM class WHERE id = ?), NOW())")
			_, err2 := stmt2.Exec(newClass.Id)
			if err2 != nil {
				log.Print(err1.Error())
				respondError(w, http.StatusInternalServerError, util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
			} else {
				log.Print(http.StatusOK)
				respondJSON(w, http.StatusOK, util.PostHttpResponse{Success: true, Message: []string{"OK"}})
			}
		}
	}
}

func AddAttendance(w http.ResponseWriter, r *http.Request) {
	log.Print("Endpoint Hit: AddAttendance")

	newAttendance := models.Attendance{}
	_ = json.NewDecoder(r.Body).Decode(&newAttendance)

	db := database.Connect()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO attendance (attendance_id, class_id, student_id, attend_time) values (UUID(), (SELECT id FROM class WHERE id = ?), (SELECT id FROM student WHERE id = ?), NOW())")
	_, err := stmt.Exec(newAttendance.ClassId, newAttendance.StudentId)

	if err != nil {
		log.Print(err.Error())
		respondError(w, http.StatusInternalServerError, util.PostHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		log.Print(http.StatusOK)
		respondJSON(w, http.StatusOK, util.PostHttpResponse{Success: true, Message: []string{"OK"}})
	}
}

func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	log.Print("Endpoint Hit: GetAllClasses")
	var allClasses []models.Class

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
		respondError(w, http.StatusInternalServerError, util.GetHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		log.Print(http.StatusOK)
		respondJSON(w, http.StatusOK, util.GetHttpResponse{Success: true, Message: []string{"OK"}, Data: models.Classes{allClasses}})
	}
}

func GetClassDetail(w http.ResponseWriter, r *http.Request) {
	log.Print("Endpoint Hit: GetClassDetail")
	var allAttendedStudents []models.Student
	var classDetail models.ClassDetail
	vars := mux.Vars(r)

	db := database.Connect()
	defer db.Close()
	results, errq := db.Query("SELECT class.id, class.class_name, class.class_time, class.room, class.created_at, class.updated_at, student.id, student.name, student.phone, student.created_at, student.updated_at FROM attendance LEFT JOIN class ON attendance.class_id = class.id LEFT JOIN student ON attendance.student_id = student.id WHERE attendance.class_id = ?", vars["id"])

	for results.Next() {
		var student models.Student
		err := results.Scan(&classDetail.Id, &classDetail.ClassName, &classDetail.ClassTime, &classDetail.Room, &classDetail.CreatedAt, &classDetail.UpdatedAt, &student.Id, &student.Name, &student.Phone, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			classDetail.ClassTime = util.Format(classDetail.ClassTime)
			classDetail.CreatedAt = util.Format(classDetail.CreatedAt)
			classDetail.UpdatedAt = util.Format(classDetail.UpdatedAt)
		} else {
			classDetail.ClassTime = util.Format(classDetail.ClassTime)
			classDetail.CreatedAt = util.Format(classDetail.CreatedAt)
			classDetail.UpdatedAt = util.Format(classDetail.UpdatedAt)
			student.CreatedAt = util.Format(student.CreatedAt)
			student.UpdatedAt = util.Format(student.UpdatedAt)
			allAttendedStudents = append(allAttendedStudents, student)
		}
	}
	classDetail.Attendance = allAttendedStudents
	if errq != nil {
		log.Print(errq.Error())
		respondError(w, http.StatusInternalServerError, util.GetHttpResponse{Success: false, Message: []string{"FAILED"}})
	} else {
		log.Print(http.StatusOK)
		respondJSON(w, http.StatusOK, util.GetHttpResponse{Success: true, Message: []string{"OK"}, Data: models.ClassDetailItem{classDetail}})
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, payload interface{}) {
	respondJSON(w, code, payload)
}
