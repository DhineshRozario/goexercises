package main

import "fmt"

func main() {
	// 1. Creating and assigning
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"lime":  "#00ff00",
		"blue":  "#0000ff",
	}

	// 2. Normal way
	// var colors map[string]string

	// 3. using make()
	// colors := make(map[string]string)

	//Adding addional values
    colors["white"] = "#ffffff"
    
    //delete the value from map
    delete(colors, "lime")

	fmt.Println(colors)
}
