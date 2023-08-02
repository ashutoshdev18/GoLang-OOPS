package main

import "fmt"

type Courses struct {
	Title   string
	Student []*Student
}

type Student struct {
	Name    string
	Courses []*Courses
}

func main() {
	ashutosh := &Student{Name: "Ashutosh"}
	pandey := &Student{Name: "Pandey"}

	math := &Courses{Title: "Maths - I"}
	science := &Courses{Title: "Biology"}

	ashutosh.Courses = append(ashutosh.Courses, math, science)
	pandey.Courses = append(pandey.Courses, science)

	math.Student = append(math.Student, ashutosh)

	science.Student = append(science.Student, pandey)

	fmt.Println(ashutosh.Name, ashutosh.Courses[0].Title)
	fmt.Println(pandey.Name, pandey.Courses[0].Title)
	fmt.Println(math.Title, math.Student[0].Name)
	fmt.Println(science.Title, science.Student[0].Name)
}
