package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.yadavamrendra789.com.np/"

func main() {
	fmt.Println("My web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	dataByte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content, "Content is this")

}
