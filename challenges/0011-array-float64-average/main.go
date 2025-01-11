// Implement a program to calculate the average of numbers in a slice
package main

import "fmt"

func calculateAvg(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func main() {
	nums := []float64{1.2, 3.5, 23.2, 67.2, 34.4, 56.9}
    fmt.Printf("The average value of %.2f is: %.3f\n", nums, calculateAvg(nums))
}
// The average value of [1.20 3.50 23.20 67.20 34.40 56.90] is: 31.067