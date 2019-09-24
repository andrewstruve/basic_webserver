package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func renderHomePage(res http.ResponseWriter) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
func renderAboutPage(res http.ResponseWriter) {
	tpl.ExecuteTemplate(res, "about.gohtml", nil)
}
