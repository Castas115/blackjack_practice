package game

import (
	"fmt"
)

type Game struct {
	deck Deck
	dealerHand Hand
	otherHands []Hand
}

func StartGame(decks int) Game {
	game := Game{}
	game.deck = StarterDeck(decks)
	game.deck.Shuffle()
	game.InitialDeal()
	return game
}

func (game *Game) InitialDeal() {
	game.Deal(&game.dealerHand)
	game.Deal(&game.dealerHand)
}

func (game *Game) Print() {
	fmt.Print("             ̈́Dealer:")
	game.dealerHand.Print(true)
	fmt.Print("")
	fmt.Print("")
	game.dealerHand.Print(true)
	game.dealerHand.Print(true)
	game.dealerHand.Print(true)
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}

