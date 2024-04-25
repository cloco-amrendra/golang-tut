package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course -into another file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname" `
	Website  string `json:"website" `
}

//fake DB

var courses []Course

//middleware,helper -into separted file

func (c *Course) IsEMpty() bool {

	/* return c.CourseId == "" && c.CourseName == "" */

	return c.CourseName == ""

}

func main() {

}

//controllers into another file

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to API HUB</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnceCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	//grab if from request

	params := mux.Vars(r)

	//loop through courses , find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}

	json.NewEncoder(w).Encode("No course found with the given id")

	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEMpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
	}

	//genereate unique id,string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id fro req

	params := mux.Vars(r)

	//loop through remove that value from index , add operation  with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	//Todo : send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop through remove that value from index

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}

}
