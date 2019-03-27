package models

import (
	"time"
)

type Class struct {
	Id        string    `json:"id"`
	ClassName string    `json:"class_name"`
	ClassTime string    `json:"class_time"`
	Room      string    `json:"room"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Classes []Class
