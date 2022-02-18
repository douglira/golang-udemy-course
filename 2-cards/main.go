package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // This 'var' declaration has global scope
	card := "Ace of Spades"   // Declare and assign a value to variable (short syntax). It's a declaration that has scope
	card = "Five of Diamonds" // Assign a new value
	fmt.Println(card)

	var deckSize int
	deckSize = 52
	fmt.Println(deckSize)
}
