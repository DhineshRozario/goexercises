package main

import (
	"fmt"
)

func main() {
	values := [] int { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

	fmt.Println("Given Values: ", values)
	
	for _, value := range values {
		
		if value % 2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}