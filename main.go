package main

import "fmt"

var card string

func main() {
	cards := newDeckFromFile("meuDeck.txt")
	fmt.Println(cards.toString())
	// cards.saveToFile("meuDeck.txt")
}
