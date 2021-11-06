package controllers

import (
	"github.com/trombettamoacyr/store-go/db"
	"html/template"
	"net/http"
	"strconv"

	"github.com/trombettamoacyr/store-go/models"
	"github.com/trombettamoacyr/store-go/services/product"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := product.FindAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		newProduct := models.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Amount:      amount,
		}

		product.Create(newProduct)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product.Delete(productId)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	currentProduct := product.Find(productId)
	templates.ExecuteTemplate(w, "Edit", currentProduct)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		newProduct := models.Product{
			Id:          id,
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Amount:      amount,
		}
		product.Edit(newProduct)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Migration(w http.ResponseWriter, r *http.Request) {
	db.Migration()
	products := product.FindAll()
	templates.ExecuteTemplate(w, "Index", products)
}
