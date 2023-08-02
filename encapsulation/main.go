package encapsulation

import "fmt"

type Person struct {
	name string //unexported field
	Name string //exported field
}

func MyName() string { //expported method
	return ""
}

func myName() string { //unexported method
	return ""
}

func main() {
	p1 := Person{name: "Ashutosh", Name: "Ashutosh"}
	fmt.Println(p1)
}
