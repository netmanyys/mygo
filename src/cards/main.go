package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}

// ❯ go run main.go deck.go
// 0 Ace of Spades
// 1 Two of Spades
// 2 Three of Spades
// 3 Four of Spades
// 4 Ace of Diamonds
// 5 Two of Diamonds
// 6 Three of Diamonds
// 7 Four of Diamonds
// 8 Ace of Hearts
// 9 Two of Hearts
// 10 Three of Hearts
// 11 Four of Hearts
// 12 Ace of Clubs
// 13 Two of Clubs
// 14 Three of Clubs
// 15 Four of Clubs
