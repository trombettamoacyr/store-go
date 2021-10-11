package routes

import (
	"net/http"
	"store-go/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
}
