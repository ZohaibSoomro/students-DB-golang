package data

import (
	"bufio"
	"os"
	"strings"

	"github.com/zohaibsoomro/database/model"
)

func ReadStudentsFromFile(fileName string) ([]model.Student, error) {
	file, err := os.Open(fileName)
	var students []model.Student
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

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

func WriteStudentToFile(student model.Student) (bool, error) {
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
	file.Sync()
	defer file.Close()
	return true, nil
}

func FindStudent(rollNo string, students []model.Student) (*model.Student, bool) {
	for _, student := range students {
		if strings.EqualFold(student.RollNo, rollNo) {
			return &student, true
		}
	}
	return nil, false
}
