//  Count the frequency of each character in a string
package main

import "fmt"

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	return freq
}

func main() {
	s := "hello world!"
	frequencies := charFrequency(s)
	fmt.Println("Character frequencies:", frequencies)
	for char, count := range frequencies {
		fmt.Printf("'%c: %d\n", char, count) // %c displays the character
	}
}

// Character frequencies: map[32:1 33:1 100:1 101:1 104:1 108:3 111:2 114:1 119:1]
// 'l: 3
// 'o: 2
// 'w: 1
// 'r: 1
// 'd: 1
// '!: 1
// 'h: 1
// 'e: 1
// ' : 1