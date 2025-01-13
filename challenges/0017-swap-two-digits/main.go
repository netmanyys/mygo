// Write a function to swap two integers without using a third variable
package main

import "fmt"

func swap(a, b int) (int, int) {
    a = a + b        
    b = a - b     // b = (a + b) - b = a
    a = a - b     // a = a + b - a = b
    return a, b
}

func main() {
    a, b := 5, 10
    fmt.Println("Before the swap:",a, b)
    a, b = swap(a,b)
    fmt.Println("After the swap ",a ,b)
}


// Before the swap: 5 10
// After the swap  10 5