package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected foirst card of Ace of Spades, but ogt %v", d[0])
	}

	if d[len(d)-1] != "Queen of Hearts" {
		t.Errorf("Expected last card of Queen of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.shuffle()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != len(d) {
		t.Errorf("Expected same of cards in both decks, but got %v , %v", len(loadedDeck), len(d))
	}

	os.Remove("_decktesting")
}
