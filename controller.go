package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type testStruct struct {
	// parameters in struct must be uppcercase so JSON libraries can access the parameters
	Test  string
	Test2 string
}

// Get Requests to handle web pages
func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /")
	renderHomePage(res)
}
func about(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /about")
	renderAboutPage(res)
}

// Post Request handler to handle input data in body of request
func formInput(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		//varify content header
		if req.Header["Content-Type"][0] == "application/x-www-form-urlencoded" {
			// infoke before trying to parse
			//ParseForm populates r.Form and r.PostForm.
			//For all requests, ParseForm parses the raw query from the URL and updates r.Form.
			req.ParseForm()

			postDataLength := req.ContentLength
			fmt.Println("POST /postData ContentLength:", postDataLength)
			// iterate over map of key value pairs in the Post Form Request
			for k, v := range req.PostForm {
				fmt.Printf("key[%s] value[%s]\n", k, v)
			}
			res.Write([]byte("{response: \"Data Submitted\"}"))
		}
	}
}

// Post request to handle json input
func jsonInput(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		//varify content header
		if req.Header["Content-Type"][0] == "application/json" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))
			var t testStruct
			err = json.Unmarshal(body, &t)
			if err != nil {
				panic(err)
			}
			log.Println(t.Test)
			log.Println(t.Test2)
		}
	}
}
