package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards.save("my_cards")
	// cards.readFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
