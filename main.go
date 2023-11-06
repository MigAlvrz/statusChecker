package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	links := []string{
		"https://www.twave.io",
		"https://golang.org",
		"http://google.com",
		"http://medium.com",
	}

	var wg sync.WaitGroup

	wg.Add(len(links))

	for _, link := range links {
		go checkAndSaveLink(link, &wg)
	}
	fmt.Println("No. rutinas activas:", runtime.NumGoroutine())
	wg.Wait()
}

func checkAndSaveLink(link string, wg *sync.WaitGroup) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " está KO")
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%s está OK, status: %d\n", link, resp.StatusCode)
	if resp.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		file := strings.Split(link, "//")[1]
		file += ".html"
		fmt.Printf("Guardando el body como %s\n", file)
		err = os.WriteFile(file, bodyBytes, 0664)
		if err != nil {
			log.Fatal(err)
		}
	}
	wg.Done()
}
