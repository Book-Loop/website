package controllers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	return HomeTmpl.Execute(w, nil)
}
