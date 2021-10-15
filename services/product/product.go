package product

import (
	"store-go/db"
	"store-go/models"
)

func Find(productId string) models.Product {
	db := db.DbConnection()
	findProduct, err := db.Query("select * from product where id = $1", productId)

	checkErro(err)

	p := models.Product{}

	for findProduct.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = findProduct.Scan(&id, &name, &description, &price, &amount)

		checkErro(err)

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount
	}
	defer db.Close()
	return p
}

func FindAll() []models.Product {
	db := db.DbConnection()
	findAllProducts, err := db.Query("select * from product")

	checkErro(err)

	p := models.Product{}
	products := []models.Product{}

	for findAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = findAllProducts.Scan(&id, &name, &description, &price, &amount)

		checkErro(err)

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func Create(product models.Product) {
	name := product.Name
	description := product.Description
	price := product.Price
	amount := product.Amount

	db := db.DbConnection()
	query, err := db.Prepare("insert into product(name, description, price, amount) values ($1, $2, $3, $4)")

	checkErro(err)

	query.Exec(name, description, price, amount)
	defer db.Close()
}

func Edit(product models.Product) {
	db := db.DbConnection()
	query, err := db.Prepare("update product set name = $1, description = $2, price = $3, amount = $4 where id = $5")

	checkErro(err)

	query.Exec(product.Name, product.Description, product.Price, product.Amount, product.Id)
	defer db.Close()
}

func Delete(id string) {
	db := db.DbConnection()
	query, err := db.Prepare("delete from product where id=$1")

	checkErro(err)

	query.Exec(id)
	defer db.Close()
}

func checkErro(err error) {
	if err != nil {
		panic(err.Error())
	}
}
