package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json: "fullname`
	Website  string `json:"website"`
}

// fake Db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grabbing id form request
	params := mux.Vars(r)

	//lop through the courses and find the matching id
	for _, course := range courses {
		if course.courseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Noe course found with the given id")
	return
}
