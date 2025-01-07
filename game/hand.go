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

func (hand *Hand) Print(hideFirstCard bool) {
	fmt.Print(" [ ")
	for _,card := range hand.cards {
		if (hideFirstCard) {
			fmt.Print("_ ")
			hideFirstCard = false
		} else {
			if (card == 1) {
				fmt.Print("*")
			} else if (card < 10) {
				fmt.Print(" ")
			}
			fmt.Print(card, " ")
		}
	}
	fmt.Println("]")
}

func (hand *Hand) Count() int {
	count := 0
	for card := range hand.cards {
		count += card
	}
	return count
}
