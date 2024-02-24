package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Alice Smith"
	pers1.age = 28
	pers1.address = "742 Evergreen Terrace"

	// Pers2 specification
	pers2.name = "Bob Johnson"
	pers2.age = 35

	//  Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Address: ", pers1.address)

	//  Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
}
