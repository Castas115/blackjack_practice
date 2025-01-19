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
		expected	Result
		err			error
	}

	t.Run("Valid results", func(t *testing.T) {
		tests := []testCase{
			{
				name:		"Blackjack vs normal",
				player:		Player{Hand: Hand{cards: []int{10, 1}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	Blackjack,
			},
			{
				name:		"Blackjack vs Blackjack",
				player:		Player{Hand: Hand{cards: []int{10, 1}}},
				dealerHand:	Hand{cards: []int{10, 1}},
				expected:	Blackjack,
			},
			{
				name:		"Win normal",
				player:		Player{Hand: Hand{cards: []int{10, 9}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	Win,
			},
			{
				name:		"Win dealer bust",
				player:		Player{Hand: Hand{cards: []int{10, 9}}},
				dealerHand:	Hand{cards: []int{10, 8, 8}},
				expected:	Win,
			},
			{
				name:		"Lose vs blackjack",
				player:		Player{Hand: Hand{cards: []int{10, 7}}},
				dealerHand:	Hand{cards: []int{10, 1}},
				expected:	Lose,
			},
			{
				name:		"Lose normal",
				player:		Player{Hand: Hand{cards: []int{10, 7}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	Lose,
			},
			{
				name:		"Busted",
				player:		Player{Hand: Hand{cards: []int{10, 7, 5}}},
				dealerHand:	Hand{cards: []int{10, 8}},
				expected:	Lose,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.player.ResolveRoundOutcome(test.dealerHand)
				actual := test.player.Result
				assert.Equal(t, test.expected, actual)
			})
		}
	})
}
