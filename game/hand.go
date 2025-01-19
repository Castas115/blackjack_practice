package game

import (
    "fmt"
)

type Hand struct {
	cards []int
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
	aceCount := 0
	for _, card := range hand.cards {
		if (card == 1) {
			aceCount++
		}
		count += card
	}
	for i := 0; i < aceCount; i++ {
		if (count + 10 <= 21){
			count += 10
		}
	}
	return count
}

func (hand *Hand) Empty()  {
	hand.cards = hand.cards[:0]
}
