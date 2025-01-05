package main

import (
	"fmt"
	"github.com/Castas115/blackjack_practice/packages"
)

func main() {
	game := StartGame(6)
	game.Print()
}

type Game struct {
	deck deck.Deck
	dealerHand Hand
	otherHands []Hand
}

func StartGame(decks int) Game {
	game := Game{}
	game.deck = deck.StarterDeck(decks)
	game.deck.Shuffle()
	game.InitialDeal()
	return game
}

func (game *Game) InitialDeal() {
	game.Deal(&game.dealerHand)
	game.Deal(&game.dealerHand)
}

func (game *Game) Print() {
	game.dealerHand.Print()
}

func (game *Game) Deal(hand *Hand) {
	card, wasCardDraw := game.deck.Pop()
	if (!wasCardDraw) {
		panic("wtf there was no deck left")
	}
	hand.Deal(card)
}

type Hand struct {
	cards []int
}

func (hand *Hand) Deal(card int) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) Print() {
	for _,card := range hand.cards {
		if (card == 1) {
			fmt.Print("*")
		} else if (card < 10) {
			fmt.Print(" ")
		}
		fmt.Print(card, " ")
	}
	fmt.Println()
}

func (hand *Hand) Count() int {
	count := 0
	for card := range hand.cards {
		count += card
	}
	return count
}
