// Write a function to find the greatest common divisor (GCD) of two numbers
package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	a, b := 108, 24
	fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd(a,b))
}

// GCD of 108 and 24 is 12