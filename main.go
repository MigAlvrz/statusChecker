package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c {
		go func(l string) {
			time.Sleep(600 * time.Second)
			go checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " está KO")
		c <- link
		return
	}
	fmt.Println(link + " está OK")
	c <- link
}
