package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to  the handling request tut")

	//PerformGetReuquest()

	//PerformPostJsonReuquest()

	PerfomrPostFormRequest()
}

func PerformGetReuquest() {
	const myUrl = "http://localhost:1111/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	//fmt.Println("Staus Code: ", response.StatusCode)
	//fmt.Println("Content Length is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count", byteCount)
	fmt.Println(responseString.String())

	/* fmt.Println(string(conent)) */

}

func PerformPostJsonReuquest() {
	const myUrl = "http://localhost:1111/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"country" : "Nepal",
		"Code" : "Nep"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerfomrPostFormRequest() {
	const myUrl = "http://localhost:1111/postform"

	//create formdata

	data := url.Values{}
	data.Add("firstname", "amrendra")
	data.Add("lastname", "yadav")
	data.Add("email", "yadavamrendra789@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("The actual response is :", string(content))
}
