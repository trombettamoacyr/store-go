package routes

import (
	"net/http"

	"github.com/trombettamoacyr/store-go/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/product", controllers.New)
	http.HandleFunc("/product/insert", controllers.Save)
	http.HandleFunc("/product/delete", controllers.Delete)
	http.HandleFunc("/product/edit", controllers.Edit)
	http.HandleFunc("/product/update", controllers.Update)

	http.HandleFunc("/migration", controllers.Migration)
}
