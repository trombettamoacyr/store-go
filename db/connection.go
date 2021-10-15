package db

import (
	"database/sql"
	//command: go mod init
	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connection := "user=postgres dbname=postgres password=Test123* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
