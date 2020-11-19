package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person {firstName: "Dewin", lastName: "Dhinesh"}

	fmt.Println(alex)
}