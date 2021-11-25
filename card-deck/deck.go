package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

var cardSuits = []string{
	"Spades",
	"Diamonds",
	"Hearts",
	"Clubs",
}

var cardValues = []string {
	"Ace",
	"Two",
	"Three",
	"Four",
}

func newDeck() Deck {
	deck := Deck{}
	for _, value := range cardValues {
		for _, suit := range cardSuits {
			deck = append(deck, value + " of " + suit)
		}
	}
	return deck
}

func (deck Deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func (deck Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		j := rand.Intn(len(deck))
		k := rand.Intn(len(deck))
	
		deck[k], deck[j] = deck[j], deck[k]
	}
}

func deal(deck Deck, handSize int) (Deck, Deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck Deck) saveDeck(filename string) error {
	fileString := strings.Join([]string(deck), ",")
	return ioutil.WriteFile(filename, []byte(fileString), 0666)
}

func loadDeck(filename string) (Deck, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return Deck(strings.Split(string(bytes), ",")), nil
}