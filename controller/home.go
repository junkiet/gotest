package controller

import "net/http"

type HomeController struct{}

func (h *HomeController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Handler"))
}
