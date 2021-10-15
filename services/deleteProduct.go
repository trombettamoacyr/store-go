package services

import (
	"store-go/db"
)

func Delete(id string) {
	db := db.DbConnection()
	query, err := db.Prepare("delete from product where id=$1")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
	defer db.Close()
}
