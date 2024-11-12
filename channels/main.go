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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " must be down.")
		return
	}
	fmt.Println(link + " is up.")
}
