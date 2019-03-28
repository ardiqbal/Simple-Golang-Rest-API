package models

type Class struct {
	Id        string `json:"id"`
	ClassName string `json:"class_name"`
	ClassTime string `json:"class_time"`
	Room      string `json:"room"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Attendance struct {
	AttendanceId string `json:"attendance_id"`
	ClassId      string `json:"class_id"`
	StudentId    string `json:"student_id"`
	AttendTime   string `json:"attend_time"`
}

type Classes []Class
