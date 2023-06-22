package main

import "fmt"

func main() {
	colors := make(map[int]string)

	colors[10] = "#ffffff"

	delete(colors, 10)
	fmt.Println(colors)
}

// map[]
