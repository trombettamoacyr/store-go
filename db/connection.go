package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	PRODUCT_TABLE = "CREATE TABLE public.product (id SERIAL UNIQUE, name varchar NOT NULL, description varchar NOT NULL, " +
		"price decimal NOT NULL, amount int NOT NULL, CONSTRAINT product_pk PRIMARY KEY (id));"
)

func DbConnection() *sql.DB {
	connection := "user=postgres dbname=postgres password=Test123* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Migration() {
	connection := "user=postgres dbname=postgres password=Test123* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	_, err = db.Query("select * from product where 1=1 limit 1;")
	if err != nil {
		createTable(db)
	}
}

func createTable(db *sql.DB) {
	_, err := db.Exec(PRODUCT_TABLE)
	if err != nil {
		panic(err.Error())
	}
}
