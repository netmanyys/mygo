// Implement linear search to find an element in a slice
package main

import "fmt"

func linearSearch(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}


func main() {
	nums := []int{156,24,15,667,123,133}
	target := 667
	index := linearSearch(nums, target)
	if index != -1 {
        fmt.Printf("Element %d found at index %d\n", target, index)
    } else {
        fmt.Printf("Element %d not found\n", target)
	}
}