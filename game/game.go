package game

import (
	// "fmt"
)

type Game struct {
	Deck		Deck
	DealerHand	Hand
	Players		[]Player
	BetSize		float32
}

func StartGame(decks int, players int) Game {
	game := Game{}
	game.Deck = StarterDeck(decks)
	game.Deck.Shuffle()
	game.Players = make([]Player, players)
	game.DealTurn()
	return game
}

func (game *Game) DealTurn() {
	game.DealerHand.Empty()
	game.Deal(&game.DealerHand)
	game.Deal(&game.DealerHand)
	for i := range game.Players {
		game.Players[i].Result = None
		game.Players[i].FinishTurn()
		game.Deal(&game.Players[i].Hand)
		game.Deal(&game.Players[i].Hand)
		if (game.Players[i].Hand.Count() == 21) {
			game.Players[i].Action = "bl"
		}
	}
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.Deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}

func (game *Game) ResolveRoundOutcome() {
	game.dealerDraws()
	for i := range game.Players {
		game.Players[i].Wager = game.BetSize
		game.Players[i].ResolveRoundOutcome(game.DealerHand)
	}
}

func (game *Game) dealerDraws() {
	if (game.DealerHand.Count() < 17) {
		game.Deal(&game.DealerHand)
		game.dealerDraws()
	}
}
