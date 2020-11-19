package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string

	contact contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func main() {
//First way of declaring the person
	alex := person {
		firstName: "Dewin", 
		lastName: "Dhinesh",
		contact : contactInfo {
			email: "dewindhinesh@gmail.com",
			zip: 638503,
		},
	}

// //another way of declaring the person
	// var alex person

	// alex.firstName = "Dewiz"
	// alex.lastName = "Dhinesh"

	// alex.contact.email = "dewizdhinesh@gmail.com"
	// alex.contact.zip = 638503

	// %+v will print verbose
	fmt.Printf("%+v", alex)

}