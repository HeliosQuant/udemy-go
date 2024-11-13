package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://stackoverflow.com/",
		"https://www.facebook.com/",
		"https://twitter.com/",
		"https://go.dev/",
		"https://www.amazon.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // in this line, the program will still wait a value of c to be displayed
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " must be down.")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link + " is up.")
	c <- "Yep it's up"
}
