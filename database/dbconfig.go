package database

import (
	"github.com/go-sql-driver/mysql"
)

func Config() *mysql.Config {
	return &mysql.Config{
		User:                 "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "class_attendance",
		AllowNativePasswords: true,
	}
}
