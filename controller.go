package main

import (
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
			fmt.Println("POST /jsonInput")
			body, err := ioutil.ReadAll(req.Body)
			fmt.Println(string(body))
			if err != nil {
				fmt.Println("JSON input handling failed")
				fmt.Println(err)
			}
			//Add handling to JSON input

			/* fmt.Println(string(body))
			var t testStruct
			err = json.Unmarshal(body, &t)
			if err != nil {
				panic(err)
			}
			log.Println(t.Test)
			log.Println(t.Test2)
			*/
		}
		res.Write([]byte("{response: \"Data Submitted\"}"))
	}
}

func getRandomNumberHandle(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		fmt.Println("GET /getRandomNumber Num:", r1.Int())
		res.Write([]byte(strconv.Itoa(r1.Intn(1000))))
	}
}

func getRandomNumberWithParametersHandle(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		keys, err := req.URL.Query()["maxNum"]

		if !err || len(keys[0]) < 1 {
			log.Println("Url Param 'maxNum' is missing")
			return
		}

		// Query()["key"] will return an array of items,
		// we only want the single item.
		key := keys[0]

		//log.Println("Url Param 'maxNum' is: " + string(key))
		maxVal, _ := strconv.Atoi(string(key))

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		fmt.Println("GET /getRandomNumberWithParam Num:", r1.Intn(maxVal))
		res.Write([]byte(strconv.Itoa(r1.Intn(maxVal))))
	}
}
