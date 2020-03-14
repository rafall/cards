package main

var card string

func main() {
	// cards := newDeckFromFile("meuDeck.txt")
	// fmt.Println(cards.toString())
	// cards.saveToFile("meuDeck.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
