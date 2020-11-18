package main

import (
	"fmt"
)

func main() {

	//Declaring Slice - array of elements which can grow.
	cards := newDeck()

	// hand1, remainingDeck1 := deal(cards, 13)

	// hand1.print()

	// hand2, remainingDeck2 := deal(remainingDeck1, 13)

	// hand2.print()

	// hand3, remainingDeck3 := deal(remainingDeck2, 13)

	// hand3.print()

	// remainingDeck3.print()


	// fmt.Println("Before Suffling ", cards.toString())

	//Shuffle
	cards.shuffle()

	fmt.Println("After Suffling \n", cards.toString())

	// cards.saveToFile("my_cards.txt")

	// deck := newDeckFromFile("my_cards.txt")

	// deck.print()

}
