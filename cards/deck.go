package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// which is an string array - can grow more
type deck []string

//Creating a set of cards for a deck
// this newDeck() function returns of type deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this function receives a type 'deck'
// which helps to consume this function by deck.print()
func (d deck) print() {
	// For index, variable := range cards which is helping to loop
	// index - index of each item
	// variable - each card item
	//range cards - which is helping to loop the cards
	//
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}

// This function gets a parameter 'deck' and from that returns two different 'deck' slices.
// this function helps to slice the given deck into two based on the handsize
// returns two deck slices (0 to give_size and from given_size to end)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converting deck to slice string into string i.e. []string to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//converting into byte and save it
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if bs != nil {
		cardsText := strings.Split(string(bs), ",")
		return deck(cardsText)
	}
	return nil
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source)

	for index := range d {
		newPosition := randomNumber.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
