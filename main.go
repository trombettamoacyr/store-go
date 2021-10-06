package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Mouse", Description: "Microsoft 1.1", Price: 95.00, Amount: 12},
		{"Keyboard", "Logitech", 150.00, 5},
		{"Camera", "Logitech", 250.00, 7},
	}

	templates.ExecuteTemplate(w, "Index", products)
}
