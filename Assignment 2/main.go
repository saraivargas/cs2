package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (person Person) Print(personNum int) {
	fmt.Printf("Person %v:\n", personNum)
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)

	if person.address != "" {
		fmt.Println("Address: ", person.address)
	}
	fmt.Println()
}

func (person Person) SCREAM() {
	fmt.Println("AAAAHHHHHHHHHHHHHHHHH")
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

	// printPersonInfo(pers1, 1)
	pers1.Print(1)
	pers2.Print(2)

	pers1.address = "320 Borden Rd"
	pers2.age = 36

	pers1.Print(1)
	pers2.Print(2)

	pers1.SCREAM()
}
