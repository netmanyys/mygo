package main

func main() {
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	cards := newDeck()
	cards.saveToFile("my_cards")
}
