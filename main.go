package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

func newCard() string {
	return "Five of Diamonds"
}
