// main.go
package main

import (
	"fmt"
	"net/http"

	"goapi/controller" // Import your controller package

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register handlers dynamically from controller package
	registerHandlers(r)

	// Start the server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func registerHandlers(r *mux.Router) {
	// Register home handler
	r.HandleFunc("/", controller.HomeHandler)

	// Register about handler
	r.HandleFunc("/about", controller.AboutHandler)
}
