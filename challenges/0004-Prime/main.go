// Write a program to check if a number is prime
package main

import "fmt"

func isPrime(n int) bool {
	if n < 1 {
		return false
	}

	for i:=2; i*i<=n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}


func main() {
	num := 29
	fmt.Printf("The num %d is prime: %t\n", num, isPrime(num))
}