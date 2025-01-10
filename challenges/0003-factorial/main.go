// Create a function to calculate the factorial of a number
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
        return 1
	}
	return n * factorial(n-1)
}

func main() {
	num := 6
	fmt.Printf("The factorial of %d is %d\n", num, factorial(num))
}