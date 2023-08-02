package main

import "fmt"

//to create a class in golang we use struct.
type Person struct {
	name       string
	age        int
	profession string
	salary     float32
}

func NewPerson(name string, age int, profession string, salary float32) *Person {
	return &Person{
		name:       name,
		age:        age,
		profession: profession,
		salary:     salary,
	}
}
func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetProfession() string {
	return p.profession
}

func (p *Person) GetSalary() float32 {
	return p.salary
}

func (p *Person) SetAge(age int) {
	p.age = age
}
func (p *Person) SetName(name string) {
	p.name = name
}
func (p *Person) SetProfession(profession string) {
	p.profession = profession
}
func (p *Person) SetSalary(salary float32) {
	p.salary = salary
}

func main() {
	p1 := Person{"Ashutosh", 24, "Software Developer", 60000.00}
	fmt.Println(p1.GetName())
	p1.SetName("Ashutosh Pandey")
	fmt.Println(p1.GetName())
	p2 := Person{}
	fmt.Println(p2.GetAge())
}
