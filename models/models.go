package models

type Student struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

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

type Classes struct {
	Items []Class `json:"classes"`
}

type ClassDetailItem struct {
	ClassDetail `json:"class"`
}

type ClassDetail struct {
	Id         string      `json:"id"`
	ClassName  string      `json:"class_name"`
	ClassTime  string      `json:"class_time"`
	Room       string      `json:"room"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
	Attendance interface{} `json:"attendance"`
}
