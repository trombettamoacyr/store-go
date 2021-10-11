package main

import (
	"net/http"
	"store-go/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
