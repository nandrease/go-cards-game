package main

func main() {
	cards := deck{newCard(), "Ace", "King"}
	cards = append(cards, "Queen", "Jack")

	cards.print()
}

func newCard() string {
	return "Joker"
}
