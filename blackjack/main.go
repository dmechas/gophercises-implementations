package main

import (
	"fmt"
	"strings"

	"github.com/dmechas/gophercises-implementations/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) Score() int {
	minScore := h.MinScore()

	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var player, dealer Hand
	everyone := []*Hand{&player, &dealer}
	cards = eachDrawTwo(everyone, cards)
	var input string
	for input != "s" {
		fmt.Println(askNext(player, dealer))
		fmt.Scanf("%s\n", &input)
		cards, player = processCommand(input, cards, player)
	}

	fmt.Println(finalReport(player, dealer))
}

func processCommand(command string, cards []deck.Card, hand Hand) ([]deck.Card, Hand) {
	var card deck.Card
	switch command {
	case "h":
		card, cards = draw(cards)
		hand = append(hand, card)
	}

	return cards, hand
}

func askNext(player, dealer Hand) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Player: %s\n", player)
	fmt.Fprintf(&b, "Dealer: %s\n", dealer.DealerString())
	fmt.Fprintf(&b, "What will you do? (h)it, (s)tand")
	return b.String()
}

func finalReport(player, dealer Hand) string {
	var b strings.Builder
	fmt.Fprintf(&b, "==FINAL HANDS==\n")
	fmt.Fprintf(&b, "Player: %s\n", player)
	fmt.Fprintf(&b, "Dealer: %s", dealer)
	return b.String()
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func eachDrawTwo(hands []*Hand, cards []deck.Card) []deck.Card {
	var card deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range hands {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	return cards
}
