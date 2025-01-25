package main

import "fmt"

func findSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	largest , secondLargest := 0, 0
	for _, v := range arr {
		if v > largest {
			secondLargest = largest
			largest = v
		} else if v > secondLargest {
			secondLargest = v
		}
	}
	return secondLargest
}

func main() {
	nums := []int{10, 5, 8, 2, 1, 9, 3, 7, 4, 6}
	fmt.Println(findSecondLargest(nums))
}
