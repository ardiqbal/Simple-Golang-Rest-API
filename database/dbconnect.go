package database

import "database/sql"

func Connect() *sql.DB {
	db, err := sql.Open("mysql", Config().FormatDSN())

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
