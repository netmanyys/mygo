// An Armstrong number (also known as a narcissistic number, pluperfect digital invariant (PPDI), or pluperfect number) is a number that is equal to the sum of its own digits raised to the power of the number of digits.
// Formula: x = d1^n + d2^n + .. + dn^n, x is a n bit digit, d1,d2,...,dn are the digits of the number.
// Example: 153 is an armstrong number, 153 = 1^3 + 5^3 + 3^3

package main

import (
	"fmt"
	"math"
)

// isArmstrong checks if a given number is an Armstrong number.
func isArmstrong(num int) bool {
	originalNum := num
	digitCount := digitLength(num)
	sum := 0

	for num > 0 {
		digit := num % 10
		sum += int(math.Pow(float64(digit), float64(digitCount)))
		num /= 10
	}
	return sum == originalNum
}

// countDigits returns the number of digits in a number.
func digitLength(num int) int {
	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count
}

func main() {
	num := 0
	fmt.Print("Please input a num: ")
	fmt.Scan(&num)

    fmt.Printf("%d is armstrong number: %v\n", num, isArmstrong(num))
}