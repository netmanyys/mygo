package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

// ❯ go run main.go stevenfile.txt
// This is just for testing.

// As of now we are not sure if this works or not.

// But we will see.

// yeah!!!%
