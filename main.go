package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " está KO")
		c <- link
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%s está OK, status: %d\n", link, resp.StatusCode)
	if resp.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		file := strings.Split(link, "//")[1]
		file += ".txt"
		fmt.Printf("Guardando el body como %s\n", file)

		err = os.WriteFile(file, bodyBytes, 0664)
		if err != nil {
			fmt.Println(err)
		}
	}
	c <- link
}
