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

	fmt.Println(alex)

//Output:
//Coop-Admin@DESKTOP-1UU7LCH MINGW64 ~/Desktop/Coop/Go/Projects/goexercises/structs (start-guide)
//$ go run main.go 
//{ }

}