package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	ch := make(chan string)
	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}

	fmt.Println(link, "is up!")
	ch <- link

}

// ❯ time go run main.go
// Success: http://stackoverflow.com is up
// Success: http://google.com is up
// Success: http://golang.org is up
// Success: http://facebook.com is up
// Success: http://amazon.com is up
// go run main.go  0.28s user 0.32s system 21% cpu 2.816 total
