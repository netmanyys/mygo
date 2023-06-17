package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

// ❯ ./main
// 0 Ace of Diamonds
// 1 Five of Diamonds
// 2 Six of Spades
