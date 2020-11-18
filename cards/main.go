package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()
	fmt.Println(card)

	id := newNumber()
	fmt.Println(id)
}

func newCard() string {
	return "Ace of Spades"
}

func newNumber() int {
	return 5
}
