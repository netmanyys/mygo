// Create a struct to represent a rectangle and calculate its area and perimeter
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area is", rect.Area())
	fmt.Println("Perimeter is", rect.Perimeter())
}

// Area is 50
// Perimeter is 30
