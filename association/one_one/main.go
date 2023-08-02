package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	address := &Address{City: "Bengaluru", Country: "India"}
	Person := &Person{Name: "Ashutosh", Address: address}

	fmt.Println(Person.Address.City)
	fmt.Println(Person.Address.Country)
	fmt.Println(Person.Name)

}
