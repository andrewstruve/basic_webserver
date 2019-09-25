package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	Title       string
	CompanyName string
}

var tpl *template.Template
var pd pageData

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	pd = pageData{
		"Struve's Basic Go Server",
		"Andrew Struve",
	}
}
func renderHomePage(res http.ResponseWriter) {
	tpl.ExecuteTemplate(res, "index.gohtml", pd)
}
func renderAboutPage(res http.ResponseWriter) {
	tpl.ExecuteTemplate(res, "about.gohtml", pd)
}
