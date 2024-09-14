package main

import "fmt"

type Student struct {
	MatricNo   string
	Firstname  string
	Lastname   string
	Department string
	Level      int
	Cgpa       float64
	Courses    []string
}

// type Group struct {
// 	Role           string
// 	Users          []Student
// 	NewestUser     Student
// 	SpaceAvailable bool
// }

func describeUser(u Student) string {
	desc := fmt.Sprintf("Name: %s %s\nDepartment: %s\nLevel: %d\nMatric no: %s\nCGPA: %f", u.Lastname, u.Firstname, u.Department, u.Level, u.MatricNo, u.Cgpa)
	return desc
}

func structs() {
	s1 := Student{MatricNo: "GLY/2018/100", Firstname: "Alabi", Lastname: "Adekunle", Department: "Geology", Level: 300, Cgpa: 4.14}
	desc := describeUser(s1)
	fmt.Println(desc)
}
