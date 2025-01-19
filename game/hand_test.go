package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	type testCase struct {
		name     string
		hand     Hand
		expected int
		err      error
	}

	t.Run("Valid results", func(t *testing.T) {
		tests := []testCase{
			{
				name:     "Blackjack with Ace and 10",
				hand:     Hand{cards: []int{10, 1}},
				expected: 21,
			},
			{
				name:     "Simple hand without Ace",
				hand:     Hand{cards: []int{10, 5}},
				expected: 15,
			},
			{
				name:     "Hand with Ace counted as 1",
				hand:     Hand{cards: []int{1, 8, 8}},
				expected: 17,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := test.hand.Count()
				assert.Equal(t, test.expected, actual)
			})
		}
	})
}
