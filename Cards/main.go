package main

func main() {
	// cards := newDeckFromFile("my_cardss")
	// deck.print(cards)

	cards := newDeck()
	cards.shuffle()
	deck.print(cards)

}
