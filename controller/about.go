package controller

import (
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle about route logic here
	w.Write([]byte("About us page"))
}
