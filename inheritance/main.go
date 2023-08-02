package main

import "fmt"

type Speaker interface {
	Speak()
}

type Animal struct{}

type Dog struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Println("Dog barks")
}

type Cat struct {
	Animal
}

func (c *Cat) Speak() {
	fmt.Println("Cat meows")
}
func main() {
	var s Speaker
	s = &Dog{}
	s.Speak()

	s = &Cat{}
	s.Speak()
}
