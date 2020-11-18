package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	value := int(52)
	if len(d) != value {
		t.Errorf("Expected deck length of %d, but got %v", value, len(d))
	}
}

func TestFirstCardNewDeck(t *testing.T) {
	d := newDeck()

	firstCard := d[0]
	expected := string("Ace of Spades")
	if firstCard != expected {
		t.Errorf("Expected first deck card as %v, but got %v", expected, firstCard)
	}
}

func TestLastCardNewDeck(t *testing.T) {
	d := newDeck()

	lastCard := d[len(d)-1]
	expected := string("King of Clubs")
	if lastCard != expected {
		t.Errorf("Expected last deck card as %v, but got %v", expected, lastCard)
	}
}

func TestSaveFileAndNewDeckFromFile(t *testing.T) {
	fileName := string("_decktesting")

	// Cleanup the file if exists
	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	value := int(52)
	if len(loadedDeck) != value {
		t.Errorf("Expected %v cards in deck, got %v", value, len(loadedDeck))
	}

	// Cleanup the file, after testing
	os.Remove(fileName)
}
