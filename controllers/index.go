package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	return IndexTmpl.Execute(w, nil)
}
