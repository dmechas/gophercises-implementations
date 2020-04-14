package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleCard(t *testing.T) {
	assert.Equal(t, "Ace of Hearts", Card{Rank: Ace, Suit: Heart}.String())
	assert.Equal(t, "Two of Spades", Card{Rank: Two, Suit: Spade}.String())
	assert.Equal(t, "Nine of Diamonds", Card{Rank: Nine, Suit: Diamond}.String())
	assert.Equal(t, "Jack of Clubs", Card{Rank: Jack, Suit: Club}.String())
	assert.Equal(t, "Joker", Card{Suit: Joker}.String())
}
