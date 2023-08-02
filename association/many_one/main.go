package main

import "fmt"

type Classroom struct {
	Name string
}

type Student struct {
	Name      string
	Classroom *Classroom
}

func main() {
	class := &Classroom{Name: "DSA"}
	stud1 := &Student{Name: "Ashutosh", Classroom: class}
	stud2 := &Student{Name: "Rupali", Classroom: class}

	fmt.Println(stud1)
	fmt.Println(stud2)
}
