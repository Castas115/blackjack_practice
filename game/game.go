package game

import (
	// "fmt"
)

type Game struct {
	Deck		Deck
	DealerHand	Hand
	Players		[]Player
}

func StartGame(decks int, players int) Game {
	game := Game{}
	game.Deck = StarterDeck(decks)
	game.Deck.Shuffle()
	game.Players = make([]Player, players)
	game.DealAll()
	return game
}

func (game *Game) DealAll() {
	game.Deal(&game.DealerHand)
	game.Deal(&game.DealerHand)
	for i := range game.Players {
		game.Deal(&game.Players[i].Hand)
		game.Deal(&game.Players[i].Hand)
	}
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.Deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}

func (game *Game) PlayerHandsAsString() []string{
	hands := []string{}
	for _,player := range game.Players {
		hands = append(hands, player.Hand.ToString(false))
	}
	return hands
}
