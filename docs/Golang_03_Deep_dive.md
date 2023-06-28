## Project Overview

Cards
  - newDeck         -> Create a list of playing cards. Essentially an array of strings
  - print           -> Log out the contents of a deck of cards
  - shuffle         -> Shuffles all the cards in a deck
  - deal            -> Create a 'hand' of cards
  - saveToFile      -> Save a list of cards to a file on the local machine
  - newDeckFromFile -> Load a list of cards from the local machine

cteate a folder `cards` under `src`

Static type of programming language same as c++/java.

Basic Go  Types

  - bool           |   true false 
  - string         |   "Hi!"  "Hows it going?"
  - int            |   0    -10000    99999
  - float64        |   10.000001    

two types of variable declarations
- var variable_name variable_type = "value"
- variable_name := "value" .   # can only be in a function

variables should only initial once and the shortcut method should only declared inside of functions

The Go playground https://go.dev/play

## Array and Slice
Array : Fixed length list of things
Slice : An array that can grow or shrink

Every element in a slice must be of same type

# How to initial a string
cards := []string{"Ace of Diamonds", newCard()}

# how to add element to slice
cards = append(cards, "Six of Spades")

# How to iterate a slice
```
for i, card := range cards {
    fmt.Println(i, card)
}
```


Cards - OO Approach

Deck Class -> Deck instance (string/functions)

Cards - Go Approach
  - Base Go Types
    * string
    * integer
    * float
    * array
    * map 
  - type deck []string
    Tell go we want to create an array of strings and add a bunch of functions specifically made to work with it
  - Functions with 'deck' as a 'receiver'
     A function with a receiver is like a "methods" - a function that belongs to an "instance"

type and receivers, type is similiar with `class`, receivers is like `methods`.


'cards' folder
[main.go] [deck.go] [deck_test.go]

receiver argument, d is similar to `self` in py
`d` is the actual copy of the deck, we're woring with is available in the function as a variable called `d`

`deck` Every variable of type 'deck' can call this function on itself


slice[start:end] \\ up to end but not include end
slice[:3] \\ 0,1,2
slice[3:] \\ 3

func (d deck) print() {}
Any variable of type "deck" now gets access to the "print" method

## Golang standard packages
https://golang.org/pkg/

## func WriteFile
https://pkg.go.dev/os#WriteFile
func WriteFile(name string, data []byte, perm FileMode) error

```
package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
```

https://pkg.go.dev/strings#Join
https://pkg.go.dev/os#ReadFile
https://pkg.go.dev/strings#Split

## Go test
Go testing is not RSpec, mocha, jasmine, selenium, etc!

To make a test, create a new file ending in _test.go
  deck_test.go
To run all tests in a package, run the command...
  go test


Deciding what to test
let's check the length of the deck that return
2nd check the first card is equal to ace of spades
3rd check the last card is equal to Four of Spades
#   
