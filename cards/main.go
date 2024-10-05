package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("================")
	// remainingCards.print()
	cards := newDeck()
	a := cards.saveToFile("my_cards")
	fmt.Println(a)
}
