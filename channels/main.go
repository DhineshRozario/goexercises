package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://golang.org",
		"http://facebook.com",
		"https://www.co-operativebank.co.uk/help-and-support/ways-to-bank",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	// To sove this problem, 'go' has 'Channel', which acts as a bridge between different 'go' Routines (all child and main routines)
	// Channel, is type specific, if we create a channel for 'string' msg communication, then we cannot use 'int' to communicate.

	//Creating a channel
	c := make(chan string)

	//go Routines - have to 'go' keyword only infront of the function.
	for _, link := range links {
		// fmt.Println("Verfiying the link:", link)
		//If we add this 'go' keyword, main Routine will create child Routines and exits the program.
		go checkLink(link, c) // added channel c, to communicate with main routine
	}

	printResponse(links, c)

	fmt.Println("Main Routine Exits")
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		// fmt.Println("Error: Link might be down:", link, err)
		//Adding message into channel, which can be read by Main Routine
		c <- " Might be down, I think"
		return
	}

	// if res.StatusCode > 0 {
	// 	status := "- is DOWN"
	// 	if res.StatusCode == 200 {
	// 		status = "- is UP"
	// 	}
	// 	// fmt.Println("Response:", link, res.StatusCode, status)
	// }
	//Adding message into channel, which can be read by Main Routine
	c <- " Yep, its up"
	// fmt.Println("Chile Routine Exits for link:", link)
}

func printResponse(links []string, c chan string) {
	// Printing the data comes into channel to console.
	for _, link := range links {
		log.Println("Response:", link, <-c)
	}
}
