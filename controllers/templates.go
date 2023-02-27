package controllers

import "html/template"

func MakeTemplate(path string) *template.Template {
	files := []string{path, "views/templates/base.html"}
	return template.Must(template.ParseFiles(files...))
}

var HomeTmpl = MakeTemplate("views/templates/home.html")
