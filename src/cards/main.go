package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	cards.print() // calling receiver print
}

func newCard() string {
	return "Five of Diamonds"
}

// ❯ ./main
// 0 Ace of Diamonds
// 1 Five of Diamonds
// 2 Six of Spades
