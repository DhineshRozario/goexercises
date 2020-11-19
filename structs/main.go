package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string

	// we can remove this variable 'contact'
	// contact contactInfo
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	//First way of declaring the person
	// dewin := person {
	// 	firstName: "Dewin",
	// 	lastName: "Dhinesh",
	// 	contactInfo : contactInfo {
	// 		email: "dewindhinesh@gmail.com",
	// 		zip: 638503,
	// 	},
	// }

	//another way of declaring the person
	var dewin person

	dewin.firstName = "Dewin"
	dewin.lastName = "Dhinesh"

	dewin.contactInfo.email = "dewindhinesh@gmail.com"
	dewin.contactInfo.zip = 638503

	dewin.print()
}

func (p person) print() {
	// %+v will print verbose
	fmt.Printf("%+v", p)
}
