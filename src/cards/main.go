package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

// ❯ go run main.go deck.go
// 0 Ace of Spades
// 1 Two of Spades
// 2 Three of Spades
// 3 Four of Spades
// 4 Ace of Diamonds
// 0 Two of Diamonds
// 1 Three of Diamonds
// 2 Four of Diamonds
// 3 Ace of Hearts
// 4 Two of Hearts
// 5 Three of Hearts
// 6 Four of Hearts
// 7 Ace of Clubs
// 8 Two of Clubs
// 9 Three of Clubs
// 10 Four of Clubs
