package db

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/zohaibsoomro/database/model"
	"github.com/zohaibsoomro/database/pkg/data"
)

// slice to hold all students' info
var students []model.Student

// loads data from the file 'fileName'
func LoadAllStudents(fileName string) {
	stds, err := data.ReadStudentsFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	students = stds
}

// get all students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, "GET") {
		http.Error(w, "Invalid Request!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	jsonBody, err := json.MarshalIndent(students, "", "\t")
	handleError(w, err)
	w.Write(jsonBody)
}

// get students with a roll number
func GetStudent(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, "GET") {
		http.Error(w, "Invalid Request!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	rollNum := r.URL.Path[len("students/rollno/")+1:]
	//check if student exists
	student, found := data.FindStudent(rollNum, students)
	if !found {
		w.Write([]byte("Student not found!"))
		return
	} else {
		//convert student data into json
		jsonBody, err := json.MarshalIndent(student, "", "\t")
		handleError(w, err)
		w.Write(jsonBody)
	}
}

// create a student with post request
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, "POST") {
		http.Error(w, "Invalid Request!", http.StatusBadRequest)
		return
	}
	//set content type of data to be in json
	w.Header().Set("Content-type", "application/json")
	var student model.Student
	defer r.Body.Close()
	handleError(w, json.NewDecoder(r.Body).Decode(&student))
	//check if students exists already
	_, found := data.FindStudent(student.RollNo, students)
	//if the student record is already present then don't add it
	if found {
		w.Write([]byte("Students already exists."))
		return
	} else {
		students = append(students, student)
		//also write this student into
		writtenToFile, err := data.WriteStudentToFile(student)
		handleError(w, err)
		if writtenToFile {
			//if new students record added then reload 'students' slice
			LoadAllStudents(model.FileName)
			println("New Length: ", len(students))
			//convert added student's data to json
			jsonBody, err := json.MarshalIndent(student, "", "\t")
			handleError(w, err)
			//write stuendent's json data to browser
			w.Write(jsonBody)
		} else {
			w.Write([]byte("Some Error occurred!"))
		}
	}
}

// function to handle any kind of errors
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		//write to browser
		w.Write([]byte(err.Error()))
		log.Fatalf("Error: %s", err.Error())
	}
}
