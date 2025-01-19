package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestResolveRoundOutcome(t *testing.T) {
	type testCase struct {
		name		string
		player		Player
		dealerHand	Hand
		expected	string
		err			error
	}

	t.Run("Valid results", func(t *testing.T) {
		tests := []testCase{
			{
				name:		"Blackjack vs normal",
				player:		Player{Hand: Hand{cards: []int{10, 1}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	"BL",
			},
			{
				name:		"Blackjack vs Blackjack",
				player:		Player{Hand: Hand{cards: []int{10, 1}}},
				dealerHand:	Hand{cards: []int{10, 1}},
				expected:	"BL",
			},
			{
				name:		"Win normal",
				player:		Player{Hand: Hand{cards: []int{10, 9}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	"wi",
			},
			{
				name:		"Win dealer bust",
				player:		Player{Hand: Hand{cards: []int{10, 9}}},
				dealerHand:	Hand{cards: []int{10, 8, 8}},
				expected:	"wi",
			},
			{
				name:		"Lose vs blackjack",
				player:		Player{Hand: Hand{cards: []int{10, 7}}},
				dealerHand:	Hand{cards: []int{10, 1}},
				expected:	"lo",
			},
			{
				name:		"Lose normal",
				player:		Player{Hand: Hand{cards: []int{10, 7}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	"lo",
			},
			{
				name:		"Busted",
				player:		Player{Hand: Hand{cards: []int{10, 7, 5}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	"bu",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.player.ResolveRoundOutcome(test.dealerHand)
				actual := test.player.Action
				assert.Equal(t, test.expected, actual)
			})
		}
	})
}
