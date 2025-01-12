package game

import (
    "fmt"
)

type Hand struct {
	cards			[]int
	wager		int
}

func (hand *Hand) Deal(card int) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) ToString(hideFirstCard bool) string {
	s := " [ "
	for _, card := range hand.cards {
		if (hideFirstCard) {
			s += "_ "
			hideFirstCard = false
		} else {
			if (card == 1) {
				s += "*"
			} else if (card < 10) {
				s += " "
			}
			s += fmt.Sprintf("%d ", card)
		}
	}
	s += "]"
	return s
}

func (hand *Hand) Count() int {
	count := 0
	for _, card := range hand.cards {
		count += card
	}
	return count
}
