package deck

import (
    "fmt"
    "math/rand"
)

func StarterDeck(deckNum int) Deck {
	if deckNum <= 0 { // TODO: How many decks should be played?
		panic("choose a sensible deck count")
    }

    deck := Deck{}
	deck.deckNum = deckNum

    for i := 1; i <= 9; i++ {
		for j := 0; j < 4*deck.deckNum; j++ {
			deck.Push(i) 
		}
    }
	for j := 0; j < 4*4*deck.deckNum; j++ {
		deck.Push(10)
	}

    return deck
}

type Deck struct {
    items []int
    deckNum int
}

func (s *Deck) Push(item int) {
    s.items = append(s.items, item)
}


func (s *Deck) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false 
    }

    top := s.items[len(s.items)-1]

    s.items = s.items[:len(s.items)-1]

    return top, true
}

func (s *Deck) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false // Return 0 and false if the stack is empty
    }
    return s.items[len(s.items)-1], true
}

func (s *Deck) IsEmpty() bool {
    return len(s.items) == 0
}

func (deck *Deck) Print() {
    for i := 0; i < len(deck.items); i++ {
		if (deck.items[i] == 1) {
			fmt.Print("*")
		} else if (deck.items[i] < 10) {
			fmt.Print(" ")
		}
		fmt.Print(deck.items[i], " ")

		if ((i+1) % (4*deck.deckNum) == 0) {
			fmt.Println()
		}
	}
}

func (deck *Deck) Shuffle() {
	for i := range deck.items {
		j := rand.Intn(i + 1)
		deck.items[i], deck.items[j] = deck.items[j], deck.items[i]
	}
}

