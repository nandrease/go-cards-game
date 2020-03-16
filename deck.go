package main

import "fmt"

// Create a new type of 'deck
// which is a slice of strings

type deck []string

func (s deck) print() {
	for i, card := range s {
		fmt.Println(i, card)
	}
}
