package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zohaibsoomro/database/model"
	"github.com/zohaibsoomro/database/pkg/db"
)

func main() {
	//load all students data on startup of server
	db.LoadAllStudents(model.FileName)
	//define all handlers
	http.HandleFunc("/students/rollno/", db.GetStudent)
	http.HandleFunc("/students", db.GetAllStudents)
	http.HandleFunc("/students/create", db.CreateStudent)

	fmt.Println("Server started...")
	handleError(http.ListenAndServe(":8080", nil))
	defer fmt.Println("Server stopped")
}

// hanldes error
func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
