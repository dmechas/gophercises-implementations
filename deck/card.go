//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Suit cenas
type Suit uint8

// Heat cenas
const (
	_ Suit = iota
	Spade
	Diamond
	Club
	Heart
	Joker
)

// Rank cenas
type Rank uint8

// Ace cenas
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card cenas
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
