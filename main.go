package main

import (
    "fmt"
    "math/rand"
    "packages/deck"
)

func main() {
	deck := deck.StarterDeck(6)
	deck.Shuffle()
	deck.Print()
}

