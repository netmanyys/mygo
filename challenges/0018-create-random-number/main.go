// Generate a random number in Go
package main

import (
	    "fmt"
		"time"
		"math/rand"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number:", rand.Intn(100))
}

// $ go run  main.go 
// Random number: 61
// $ go run  main.go 
// Random number: 9
// $ go run  main.go 
// Random number: 75
// $ go run  main.go 
// Random number: 19
// $ go run  main.go 
// Random number: 33