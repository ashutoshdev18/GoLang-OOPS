package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

type Employee struct {
	Person
	Salary     float64
	Profession string
}

type Student struct {
	Person
	EducationLevel string
	Degree         string
	Major          string
}

func main() {
	employee := Employee{
		Person:     Person{Name: "Ashutosh", Age: 24, Address: "Krishnagiri"},
		Salary:     160000.00,
		Profession: "Software Developer",
	}

	student := Student{
		Person:         Person{Name: "Rupali", Age: 23, Address: "Gaya"},
		EducationLevel: "Graduate",
		Degree:         "B.E.,",
		Major:          "Computer Science and Engineeringz",
	}

	fmt.Println(employee)
	fmt.Println()
	fmt.Println(student)
}
