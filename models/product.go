package models

import "store-go/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func FindAllProducts() []Product {
	db := db.DbConnection()
	findAllProducts, err := db.Query("select * from product")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for findAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = findAllProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	db.Close()
	return products
}
