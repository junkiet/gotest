package controller

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle home route logic here
	w.Write([]byte("Welcome to the home page"))
}
