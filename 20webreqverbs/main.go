package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// performGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("bytecount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename": "let's go with golang",
		"price": 0,
		"platform": "learnCodeOnline.in"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Suyash")
	data.Add("lastname", "Tiwari")
	data.Add("email", "suyash@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
