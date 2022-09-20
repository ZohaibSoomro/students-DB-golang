package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}
	handleError(http.ListenAndServe("127.0.0.1:"+port, nil))
	defer fmt.Println("Server stopped")
}

// hanldes error
func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
