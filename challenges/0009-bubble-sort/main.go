// Sort an array of integers using Bubble Sort
package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{23,34,14,156,127,12,53,10}
	fmt.Println("The sorted array is ", bubbleSort(arr))
}

// The sorted array is  [10 12 14 23 34 53 127 156]