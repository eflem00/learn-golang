package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	fmt.Println(quote.Glass());

	message, err := greetings.Greet("Gladys")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)

	names := []string{
		"evan",
		"niki",
	}
	messages, err := greetings.Greets(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		message := messages[name]

		fmt.Println(message)
	}
}