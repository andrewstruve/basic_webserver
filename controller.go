package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type testStruct struct {
	// parameters in struct must be uppcercase so JSON libraries can access the parameters
	Test  string
	Test2 string
}

// Get Requests to handle web pages
func indexPageHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /")
	renderHomePage(res)
}
func aboutPageHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /about")
	renderAboutPage(res)
}
func testDataHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /testData")
	renderTestDataPage(res)
}

// Post Request handler to handle input data in body of request
func formInputHandle(res http.ResponseWriter, req *http.Request) {
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
func jsonInputHandle(res http.ResponseWriter, req *http.Request) {
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
		res.Write([]byte("{response: \"Data Submitted\"}"))
	}
}

func getRandomNumberHandle(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		res.Write([]byte(strconv.Itoa(r1.Int())))
	}
}
