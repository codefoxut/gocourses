package main

import (
	"fmt"
	"course2/cards"
)

func main() {
	// var card string = "Ace of Spades"
	cards := cards.NewDeck()
	fmt.Println(cards.toString())
	// cards.print()
	hand, remainingCards := cards.deal(9)
	hand.print()
	fmt.Println("Remaining Cards...")
	remainingCards.print()
	cards.saveToFile("my_cards")

	newCards := cards.NewDeckFromFile("my_cards")
	newCards.print()
	fmt.Println("SHUffle the cards.")
	newCards.shuffle()
	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
