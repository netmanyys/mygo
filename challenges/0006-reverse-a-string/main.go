// Write a function to reverse a string
package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j] , runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Reverse string:", reverseString("golang"))
}
// Reverse string: gnalog