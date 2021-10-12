package services

import (
	"fmt"
	"store-go/db"
	"store-go/models"
)

func Create(product models.Product) {

	name := product.Name
	description := product.Description
	price := product.Price
	amount := product.Amount

	fmt.Println(product)

	db := db.DbConnection()
	query, err := db.Prepare("insert into product(name, description, price, amount) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("batataasss")
	query.Exec(name, description, price, amount)
	defer db.Close()
}
