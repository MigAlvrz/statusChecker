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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " puede que esté caido")
		return
	}
	fmt.Println(link + " está funcionando")
}
