// Find the sum of all elements in an array
package main

import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	arr := []int{23,12,234,12,2324,12,1}
	fmt.Printf("The sum of arr is %d\n", sumArray(arr))
}

// The sum of arr is 2618