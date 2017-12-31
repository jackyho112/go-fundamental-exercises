package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person is a pointer type in type annotation but a retriever of the value in memory as an operator
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// & gets the memory address of a value
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.updateName("jimmy")
	jim.print()
}
