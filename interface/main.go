package main

import (
	"fmt"
)

type bot interface {
	getGreating() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}

func printGreating(eb bot) {
	fmt.Println(eb.getGreating())
}

func (englishBot) getGreating() string {
	return "Hi there!"
}

func (spanishBot) getGreating() string {
	return "Hola"
}
