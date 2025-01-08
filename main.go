package main

import (
	"github.com/Castas115/blackjack_practice/game"
)

func main() {
	game := game.StartGame(6, 2)
	game.Print()
}

