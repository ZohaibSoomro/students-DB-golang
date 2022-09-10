package main

import (
	"log"
	"net/http"

	"github.com/zohaibsoomro/database/model"
	"github.com/zohaibsoomro/database/pkg/db"
)

func main() {
	db.LoadAllStudents(model.FileName)
	http.HandleFunc("/students/rollno/", db.GetStudent)
	http.HandleFunc("/students", db.GetAllStudents)
	http.HandleFunc("/students/create", db.CreateStudent)
	handleError(http.ListenAndServe(":8080", nil))
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
