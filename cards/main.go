package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
	a := newDeck()
	fmt.Println(a)
}

func newCard() string {
	return "Five of Diamonds"
}
