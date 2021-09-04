package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	fmt.Println(cards.toString())
	// cards.print()
	hand, remainingCards := cards.deal(9)
	hand.print()
	fmt.Println("Remaining Cards...")
	remainingCards.print()
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.print()
	fmt.Println("SHUffle the cards.")
	newCards.shuffle()
	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
