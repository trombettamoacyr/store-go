package main

import (
	"database/sql"
	"net/http"
	"text/template"

	//command: go mod init
	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func dbConnection() *sql.DB {
	connection := "user=postgres dbname=postgres password=Teste123* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
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

	templates.ExecuteTemplate(w, "Index", products)
	db.Close()
}
