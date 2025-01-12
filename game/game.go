package game

import (
	// "fmt"
)

type Game struct {
	Deck Deck
	DealerHand Hand
	PlayerHands []Hand
}

func StartGame(decks int, playerHands int) Game {
	game := Game{}
	game.Deck = StarterDeck(decks)
	game.Deck.Shuffle()
	game.PlayerHands = make([]Hand, playerHands)
	game.DealAll()
	return game
}

func (game *Game) DealAll() {
	game.Deal(&game.DealerHand)
	game.Deal(&game.DealerHand)
	for i := range game.PlayerHands {
		game.Deal(&game.PlayerHands[i])
		game.Deal(&game.PlayerHands[i])
	}
}

func (game *Game) PlayerHandsAsString() []string{
	hands := []string{}
	for _,hand := range game.PlayerHands {
		hands = append(hands, hand.ToString(false))
	}
	return hands
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.Deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}



