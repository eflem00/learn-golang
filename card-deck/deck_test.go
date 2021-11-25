package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("unexpected deck len of %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("unexpected first card %v", deck[0])
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	err := deck.saveDeck("_decktesting")

	if err != nil {
		t.Errorf("error saving deck")
	}

	deck, err = loadDeck("_decktesting")

	if err != nil {
		t.Errorf("error loading deck")
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("unexpected first card %v", deck[0])
	}

	os.Remove("_decktesting")
}