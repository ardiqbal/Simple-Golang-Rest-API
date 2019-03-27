package models

type Class struct {
	Id        string `json:"id"`
	ClassName string `json:"class_name"`
	ClassTime string `json:"class_time"`
	Room      string `json:"room"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Classes []Class
