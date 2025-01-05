package main

import (
    "github.com/Castas115/blackjack_practice/packages"
)

func main() {
	deck := deck.StarterDeck(6)
	deck.Shuffle()
	deck.Print()
}

