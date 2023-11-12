package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome API")
	fmt.Println("http://localhost:3000")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 299, Author: &Author{FullName: "Shubham", Website: "Sp.in"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Angular", CoursePrice: 899, Author: &Author{FullName: "Shubham", Website: "Sp.in"}})

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

// Controller
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome Go</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	// grab id
	params := mux.Vars(r)

	// loop courses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with give id")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course~")
	w.Header().Set("Content-Type", "application/json")

	// body empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data")
		return
	}

	// {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No actual data")
		return
	}

	for _, v := range courses {
		if course.CourseName == v.CourseName {
			json.NewEncoder(w).Encode("Already exist")
			return
		}
	}

	rand.Seed(time.Now().Unix())
	// course.CourseId = string(rand.Intn(100))
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var c Course
			json.NewDecoder(r.Body).Decode(&c)
			c.CourseId = params["id"]
			courses = append(courses, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	// TODO: If not found
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}

	// TODO: If not found
}
