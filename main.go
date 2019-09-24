package main

import (
	"log"
	"net/http"
)

func main() {

	//initDb()

	// serve up a webpage
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/about", http.HandlerFunc(about))
	http.Handle("/formInput", http.HandlerFunc(formInput))
	http.Handle("/jsonInput", http.HandlerFunc(jsonInput))

	// place all resources in the public folder, includes js, css, etc
	fs := http.FileServer(http.Dir("./public"))
	// Anything that has a resources folder,
	// strip off the resources and look for everything in public
	http.Handle("/resources/", http.StripPrefix("/resources", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
