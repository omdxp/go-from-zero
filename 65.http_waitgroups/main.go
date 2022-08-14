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

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> status code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			file := strings.Split(url, "//")[1] // take second element (example http://www.google.com -> www.google.com)
			file += ".html"

			fmt.Printf("writing response body to %s\n", file)

			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	// this program will save body responses concurrently for all specified urls
	var wg sync.WaitGroup
	urls := []string{"https://go.dev", "https://www.google.com/a.html", "https://www.medium.com"}

	wg.Add(len(urls))
	for _, v := range urls {
		go func(v string) {
			defer wg.Done()
			checkAndSaveBody(v)
			fmt.Println(strings.Repeat("#", 20))
		}(v)
	}

	fmt.Println("number of goroutines:", runtime.NumGoroutine()) // 4 goroutines

	wg.Wait()

}
