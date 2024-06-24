package controller

import "net/http"

type AboutController struct{}

func (a *AboutController) AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Handler"))
}
