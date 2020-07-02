package main

import (
	"github.com/desferreira/alurago/routes"
	"net/http"
)

func main() {
	routes.InitRoutes()
	http.ListenAndServe(":9000", nil)
}
