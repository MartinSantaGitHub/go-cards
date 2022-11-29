package main

// Golang is a pass by value language

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   struct {
		email   string
		zipCode int
	}
}

func main() {
	//var alex Person

	alex := Person{
		lastName:  "Anderson",
		firstName: "Alex",
		contact: ContactInfo{
			email:   "aa@registry.com",
			zipCode: 1551,
		},
	}

	alex.updateName("Pablo")

	alex.print()
}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
