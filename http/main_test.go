package main

import (
    "fmt"
    "os"
)


func TestsomeMain() {
    n, err := main()
    if (err!=nil) {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
	fmt.Println(n)
}