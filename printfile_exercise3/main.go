package main

import (
	"io"
	"log"
	"os"
)


func main() {
    log.Println("Args: ", os.Args)
    
    f, err := os.Open(os.Args[1])

    if (err != nil) {
        log.Fatal("Error:", err)
        os.Exit(1)
    }

    //when checking the functions in File, we do have similar Reader interface implemented method is available as 'file.Read()'
    // - here, on the 2nd argument, we can pass directly 'f' as an argument 
    // - 'go' internally maps the interface methods and reads data from file and passes it into 'Stdout'
    io.Copy(os.Stdout, f)
}