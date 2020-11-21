package main

import (
	"log"
	"net/http"
	"time"
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

    // all the response received in 'channel c'
	// looping with the data comes into channel multiple times.
	for l := range c {
		go anonymus(l, c)
	}

	log.Println("Main Routine Exits")
}
func anonymus(link string, c chan string) {  
    time.Sleep(5 * time.Second)
    checkLink(link, c)
}

func checkLink(link string, c chan string) {
    

	res, err := http.Get(link)

	if err != nil {
		log.Println("Error: Link might be down:", link, err)
		//Adding message into channel, which can be read by Main Routine
		c <- link
		return
	}

	if res.StatusCode > 0 {
		status := "- is DOWN"
		if res.StatusCode == 200 {
			status = "- is UP"
		}
		log.Println("Response:", link, res.StatusCode, status)
	}
	//Adding message into channel, which can be read by Main Routine
	c <- link
	// fmt.Println("Chile Routine Exits for link:", link)
}
