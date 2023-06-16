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

receiver function, d is similar to `self` in py

slice[start:end] \\ up to end but not include end
slice[:3] \\ 0,1,2
slice[3:] \\ 3

func (d deck) print() {}
Any variable of type "deck" now gets access to the "print" method



#   

                        Value Types     |  Reference Types
  use                     /    int      |  slices     \
  pointers                |    float    |  maps        | Don't worry about pointers
  to change  -------------|    string   |  channels    | with these
  these things            |    bool     |  pointers    |
  in a function            \   structs  |  functions  /

  # map 
  a hash map 
   Map 
   key -> value
   key -> value
   key -> value