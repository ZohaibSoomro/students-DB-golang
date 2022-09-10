package data

import (
	"bufio"
	"os"
	"strings"

	"github.com/zohaibsoomro/database/model"
)

//this will read student info from the info file
//and then creates slice of students and returns it

func ReadStudentsFromFile(fileName string) ([]model.Student, error) {
	//open file in read mode only
	file, err := os.Open(fileName)
	var students []model.Student
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	//split file content to read it line by line
	scanner.Split(bufio.ScanLines)

	//read students info create student and then append to the 'students' slice
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			name := text
			scanner.Scan()
			text = scanner.Text()
			if len(text) > 0 {
				rollNo := text
				scanner.Scan()
				text := scanner.Text()
				if len(text) > 0 {
					section := text
					student := model.Student{Name: name, RollNo: rollNo, Section: section}
					students = append(students, student)
				}
			}
		}
	}
	return students, nil
}

// this will write the 'student' student's info to the info file
func WriteStudentToFile(student model.Student) (bool, error) {
	//open file in append mode
	file, err := os.OpenFile(model.FileName, os.O_APPEND, 0644)
	if err != nil {
		return false, err
	}
	_, err = file.WriteString("\n" + strings.Title(student.Name))
	if err != nil {
		return false, err
	}
	_, err = file.WriteString("\n" + strings.ToUpper(student.RollNo))
	if err != nil {
		return false, err
	}
	_, err = file.WriteString("\n" + student.Section)
	if err != nil {
		return false, err
	}
	//force it to write to the file
	file.Sync()
	defer file.Close()
	return true, nil
}

// it checks if any student with roll number 'rollNo' exists in 'students' slice
// if exists then it returns that pointer to student struct and true
// otherwise returns nil and false
func FindStudent(rollNo string, students []model.Student) (*model.Student, bool) {
	for _, student := range students {
		if strings.EqualFold(student.RollNo, rollNo) {
			return &student, true
		}
	}
	return nil, false
}
