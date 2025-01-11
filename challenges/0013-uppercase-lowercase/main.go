// Convert a string to uppercase and lowercase
package main

import (
    "fmt"
	"strings"
)


func main() {
	s := "GoLang"
	fmt.Println("Uppercase:", strings.ToUpper(s))
	fmt.Println("Lowercase:", strings.ToLower(s))
}


// Uppercase: GOLANG
// Lowercase: golang