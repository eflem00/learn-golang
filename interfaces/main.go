package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct {
	test int
}

func (bot EnglishBot) getGreeting() string {
	return "Hello there"
}

type SpanishBot struct{}

func (bot SpanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(bot Bot) {
	fmt.Println(bot.getGreeting())
}

func main() {
	englishBot := EnglishBot{}

	spanishBot := SpanishBot{}

	printGreeting(englishBot)

	printGreeting(spanishBot)
}
