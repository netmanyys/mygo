# Struct 
Data structure, Collection of properties that are related together.


Card Struct Field

suit -> String 
value -> string
---
Card Struct
suit -> "Spades"
value -> "Ace"

&variable turn value to `address`
*pointer turn `address` to value

```
                     type description (it means we're working with a pointer to a person)
func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname
}     ^
      |
  This is an operator - it means we want to manupulate the value the pointer is referncing
```



# playground
```
package main

import "fmt"

func main() {
  mySlice := []string{"Hi", "there", "How", "Are", "You"}
  updateSlice(mySlice)
  fmt.Println(MySlice)
}

func updateSlice(s []string) {
  s[0] = "Bye"
}

\\ [Bye There How Are You]
\\ it modified the original slice


```

Difference between Arryas and  Slices

   Arrays        |     Slices
Primitive data   |     Can gow and shrink
Can't be resized |     Used 99% of the time for lists of elements
Rarely used 
directly

# very important in go!
# Value type vs Reference Type
                        Value Types     |  Reference Types
  use                     /    int      |  slices     \
  pointers                |    float    |  maps        | Don't worry about pointers
  to change  -------------|    string   |  channels    | with these
  these things            |    bool     |  pointers    |
  in a function            \   structs  |  functions  /

-------------------------------------------------------------------
 RAM                                                              |
 Address     Value                                                |
 0000                                                             |
 0001      length|cap|ptr to head -------------|       MySlice    |
 0002    []string{"Hi","there", "How","are", "you?"}              |
 0003                                          |                  |
 0004      length|cap|ptr to head --- ---------         copy      |
 ....                                                             |
-------------------------------------------------------------------

```
package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
\\0xc000012028
\\0xc000012038

```



```
package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&*namePointer)
}
\\0xc00010c230
\\0xc00010c230
```
  # map 
  a hash map 
   Map 
   key -> value
   key -> value
   key -> value