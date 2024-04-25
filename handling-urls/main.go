package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.yadavamrendra789.com.np:8000/?coursename=reactjs&payment=esewa"

func main() {

	/* fmt.Println("Handling web urls")
	fmt.Println(myUrl) */

	//parsing the url

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	/* fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port()) */

	qParams := result.Query()

	fmt.Println(qParams)

	fmt.Println(qParams["coursename"][0])

	for _, val := range qParams {
		fmt.Println("Params is :", val[0])
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "test.dev",
		Path:     "/tutcss",
		RawQuery: "user=amrendra",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
