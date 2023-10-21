package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("server in go")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("content", string(content))
	fmt.Println("byte count: ", byteCount)
	fmt.Println("content using strings.Builder : ", responseString.String())

}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{	
			"coursename": "golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("contenti", string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"
	// form data

	data := url.Values{}
	data.Add("firstname", "mahir")
	data.Add("lastname", "neema")
	data.Add("email", "mahir@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("contennt", string(content))
}
