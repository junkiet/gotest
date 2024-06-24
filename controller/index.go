package controller

import "net/http"

type IndexController struct{}

func (a *IndexController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/ Handler"))
}
