package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func main() {
	// alex := person {firstName: "Dewin", lastName: "Dhinesh"}

	var alex person

	alex.firstName = "Dewiz"
	alex.lastName = "Dhinesh"

	// %+v will print verbose
	fmt.Printf("%+v", alex)

}