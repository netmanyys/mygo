package main

import "fmt"

type bot interface {
	getGreeting() string
}

// if you are a type in thie program with a function named "getGreeting"
// and you return a string then you are now an honorary member of type 'bot'
// and now you can call function printGreeting

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

//	func printGreeting(sb spanishBot) {
//		fmt.Println(sb.getGreeting())
//	}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

// ❯ go run main.go
// Hi There!
// Hola

// struct  englishbot spansihbot          interface  bot
// func    getGreeting getGreeting           printGreeting
//                                              main
