package main

import (
	"fmt"

	"github.com/dmechas/gophercises-implementations/deck"
)

func ExampleHand_Score() {
	fmt.Println(Hand{deck.Card{Rank: deck.King}}.Score())
	fmt.Println(Hand{deck.Card{Rank: deck.King}, deck.Card{Rank: deck.King}}.Score())
	fmt.Println(Hand{deck.Card{Rank: deck.Ace}, deck.Card{Rank: deck.King}}.Score())
	fmt.Println(Hand{deck.Card{Rank: deck.Ace}}.Score())
	fmt.Println(Hand{deck.Card{Rank: deck.Ace}, deck.Card{Rank: deck.Ace}}.Score())

	// Output:
	// 10
	// 20
	// 21
	// 11
	// 12
}
