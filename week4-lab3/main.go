package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonors() bool {
	return s.GPA >= 3.50
}
func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4.00 {
		return errors.New("GPA must be between 0-4")
	}
	return nil
}

func main() {
	// var st Student = Student{ID: "1", Name: "jarudet", Email: "prungprechskun_j@su.ac.th", Year: 4, GPA: 3.75}
	students := []Student{
		{ID: "1", Name: "jarudet", Email: "prungprechskun_j@su.ac.th", Year: 4, GPA: 3.75},
		{ID: "2", Name: "alice", Email: "alice@hotmail.com", Year: 4, GPA: 2.75},
	}

	newStudent := Student{Name: "trudy", Email: "trudy@hotmail.com", Year: 4, GPA: 3.50}
	students = append(students, newStudent)

	for i, student := range students {
		fmt.Printf("%d Honor = %v\n", i, student.IsHonors())
		fmt.Printf("%d Validation = %v\n", i, student.Validate())
	}

}
