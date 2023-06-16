## Five big questions
* How do we run the code in our project?
  ```
   ❯ go run main.go
     Hi there!      \\ compile and execute the go file
  ```
* What does 'package main' mean?
* What does 'import "fmt"' mean?
* What is that 'func' thing?
* How is the main.go file organized?

## Go CLI
* go build   -> Compiles a bunch of go source code files
* go run     -> Compiles and executes one or two files
* go fmt     -> Formats all the code in each file in the current directory
* go install -> Compiles and "Installs" a package.
* go get     -> Downloads the raw source code of someone else's package
* go test    -> Runs any tests associated with the current project

```
 $ ls
main.go
 $ go run main.go  \\ compile and execute the go file
Hi there!
 $ ls
main.go
 $ go build main.go \\ only compile the go file
 $ ls
main    main.go \\ execute it
 $ ./main
Hi there!
```

Package == Project == Workspace


Types of packages
* Executable, generates a file that we can run, must have a function named main
* Reusable, code used as 'helpers', good place to put reusable logic, library

[golang package](golang.org/pkg/) 