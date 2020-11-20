package main

import (
	"log"
	"os"
)


func main() {

    args := os.Args

    if (len(args) <= 1) {
        log.Println("FILE_NAME argument not found, usage: 'go run main.go <FILE_NAME>'")
        os.Exit(1)
    }
    
    log.Println("Args", args)

    fileName := args[1]

    log.Println("Given file name:", fileName)

    file, err := os.Open(fileName)

    if (err != nil) {
        log.Println("Unable read the file:", fileName)
        os.Exit(1)
    }


	data := make([]byte, 10000)
    
    count, err := file.Read(data)
    
	if err != nil {
		log.Fatal(err)
    }
    
	log.Printf("From file %v, read %d bytes\n%q\n", fileName, count, data[:count])

}