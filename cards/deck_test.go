package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("wrong deck length: %v, expected 52", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("not expecting %v as first card", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("not expecting %v as last card", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 card in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
