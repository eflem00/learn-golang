package main

import (
	"fmt"
	"log"
)

func main() {
	deck := newDeck()

	deck.shuffle()

	hand, deck := deal(deck, 5)

	fmt.Println("--- hand")

	hand.print()

	fmt.Println("--- deck")

	deck.print()

	err := deck.saveDeck("test")

	if err != nil {
		log.Fatalf("Error saving deck: %v", err)
	}

	newDeck, err := loadDeck("test")

	if err != nil {
		log.Fatalf("Error loading deck: %v", err)
	}

	fmt.Println("--- new deck")

	newDeck.print()
}