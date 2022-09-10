package model

//name of file that stores students info
const FileName = "pkg/data/info.txt"

//model for a student
type Student struct {
	Name    string `json:"name"`
	RollNo  string `json:"roll_no"`
	Section string `json:"section"`
}
