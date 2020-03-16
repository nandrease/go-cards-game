package main

func main() {
	cards := newDeckFromFile("kaardipakk")
	cards.shuffle()
	cards.print()
}
