package main

import "fmt"

type Bot interface {
	GetGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb EnglishBot) GetGreeting() string {
	return "Hi There!"
}

func (sb SpanishBot) GetGreeting() string {
	return "Hola!"
}

func printGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}
