package main

func main() {
	// cards := newDeckFromFile("myCards")
	// cards.print()

	cards := newDeck()
	// cards.saveToFile("myCards")

	// fmt.Println(cards.toString())

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
