// Implement a function to calculate the Fibonacci series up to a given number

package main

import "fmt"

func fibonacci(n int) []int {
	fib := []int{0,1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

func main() {
	num := 10
	fmt.Println("Fibonacci series() %d, %v", num, fibonacci(num))
}