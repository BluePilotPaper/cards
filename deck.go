package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	cardValues := []string{"Ace", "Two,", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// parameters, returning multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// type conversion and built in functions for strings
	return strings.Join([]string(d), ",")

}

// usage: cards.saveToFile()
func (d deck) saveToFile(filename string) error {
	fmt.Println("creating new deck")
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is full permission
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// Error handling example
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}
