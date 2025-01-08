package game

import (
	"fmt"
)

type Game struct {
	deck Deck
	dealerHand Hand
	playerHands []Hand
}

func StartGame(decks int, playerHands int) Game {
	game := Game{}
	game.deck = StarterDeck(decks)
	game.deck.Shuffle()
	game.playerHands = make([]Hand, playerHands)
	game.InitialDeal()
	return game
}

func (game *Game) InitialDeal() {
	game.Deal(&game.dealerHand)
	game.Deal(&game.dealerHand)
	for i := range game.playerHands {
		game.Deal(&game.playerHands[i])
		game.Deal(&game.playerHands[i])
	}
}

func (game *Game) Print() {
	fmt.Print("             ̈́Dealer:")
	game.dealerHand.Print(true)
	fmt.Println("")
	fmt.Println("")
	for i := range game.playerHands {
		fmt.Print("Hand ",i+1,":")
		game.playerHands[i].Print(false)
		fmt.Println("")
	}
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}

