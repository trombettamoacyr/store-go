package routes

import (
	"net/http"
	"store-go/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/product", controllers.NewProduct)
	http.HandleFunc("/product/insert", controllers.SaveProduct)
}
