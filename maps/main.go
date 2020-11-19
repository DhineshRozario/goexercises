package main

import "fmt"

func main() {
	// 1. Creating and assigning
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
		"white": "#ffffff",
		"black": "#000000",
	}

	fmt.Println(colors)

	// map is passed by reference
	printMap(colors)

}

// map is passed by reference, if we change any value inside the map, then it will be updated in calling place
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for color %v -> %v\n", color, hex)
	}
}
