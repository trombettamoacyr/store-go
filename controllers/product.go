package controllers

import (
	"html/template"
	"net/http"
	"store-go/models"
	"store-go/services"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		product := models.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Amount:      amount,
		}

		services.Create(product)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
