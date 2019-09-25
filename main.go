package main

import (
	"log"
	"net/http"
)

func main() {

	//initDb()

	// serve up a webpage
	http.Handle("/", http.HandlerFunc(indexPageHandle))
	http.Handle("/about", http.HandlerFunc(aboutPageHandle))
	http.Handle("/testData", http.HandlerFunc(testDataHandle))

	// example handles for post Request data
	http.Handle("/formInput", http.HandlerFunc(formInputHandle))
	http.Handle("/jsonInput", http.HandlerFunc(jsonInputHandle))

	// example handle for a get request
	http.Handle("/getRandomNumber", http.HandlerFunc(getRandomNumberHandle))

	// place all resources in the public folder, includes js, css, etc
	fs := http.FileServer(http.Dir("./public"))
	// Anything that has a resources folder,
	// strip off the resources and look for everything in public
	http.Handle("/resources/", http.StripPrefix("/resources", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
