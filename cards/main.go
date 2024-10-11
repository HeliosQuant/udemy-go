package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("================")
	// remainingCards.print()
	cards := newDeckFromFile("my_cards")
	cards.print()
	fmt.Println("===============================")
	cards.shuffle()
	cards.print()
}
