// Write a program to find the largest number in a slice
package main

import "fmt"

func largestNum(nums []int) int {
	largest := nums[0]
	
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func main() {
	nums := []int{12,23,22,24,56,123,242,232}
	fmt.Println("The largest num of nums is", largestNum(nums))
}