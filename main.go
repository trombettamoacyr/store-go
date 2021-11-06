package main

import (
	"net/http"

	"github.com/trombettamoacyr/store-go/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
