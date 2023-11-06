package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.twave.io",
		"https://golang.org",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " está KO"
		return
	}
	c <- link + " está OK"
}
