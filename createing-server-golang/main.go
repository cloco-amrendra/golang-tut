package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to  the handling request tut")

	PerformGetReuquest()
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

	conent, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(conent)
	fmt.Println("Byte count", byteCount)
	fmt.Println(responseString.String())

	/* fmt.Println(string(conent)) */

}
