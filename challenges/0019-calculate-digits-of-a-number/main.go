package main

import "fmt"

func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func main() {
	num := 2379
	fmt.Printf("The %d's digits sum is %d\n", num, digitSum(num))
}