package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling  json files")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJs Bootcamp", 299, "LearncodeOnline", "abc123", []string{"web", "tut"}},
		{"Python Bootcamp", 4000, "LearncodeOnline", "abc123", []string{"web", "tut"}},
		{"DataScience", 4000, "LearncodeOnline", "abc123", nil},
	}

	//package this data as JSON data

	//finalJson, err := json.Marshal(myCourses)
	//finalJson, err := json.MarshalIndent(myCourses, "amr", "\t")
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`

        {
                "coursename": "ReactJs Bootcamp",
                "price": 299,
                "website": "LearncodeOnline",
                "tags": [
                        "web",
                        "tut"
                ]
        }

	`)

	var myCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		//fmt.Println("JSON Was Valid")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		//fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Println(myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v \n", k, v)
	}

}
