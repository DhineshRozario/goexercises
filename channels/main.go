package main

import (
	"fmt"
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

	//go Routines - have to 'go' keyword only infront of the function.
	for _, link := range links {
		fmt.Println("Verfiying the link:", link)
		//If we add this 'go' keyword, main Routine will create child Routines and exits the program.
		go checkLink(link)
	}


	fmt.Println("Main Routine Exits")
}
func checkLink(link string) {

	res, err := http.Get(link)

	if err != nil {
		fmt.Println("Error: Link might be down:", link, err)
		return
	}

	if res.StatusCode > 0 {
		status := "- is DOWN"
		if res.StatusCode == 200 {
			status = "- is UP"
		}
		fmt.Println("Response:", link, res.StatusCode, status)
	}

	fmt.Println("Chile Routine Exits for link:", link)
}
