package main

import "fmt"

//After adding the 'getGreeting() string' this method into the interface, now the both englishBot/spanishBot are type of bot now.
//Go does this automatically, as both bots have alreday has the method 'getGreeting() string' with same return type
//And Bot has been passed to the method 'printGreeting'
//So, now from main, when we call the 'printMethod' by passing englishBot, then automatically, bot is mapped calls the respective 'getGreeting' method
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// func printGreeting(eb englishBot) {
//     fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
//     fmt.Println(sb.getGreeting())
// }
