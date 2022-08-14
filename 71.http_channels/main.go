package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down!\n", url)
		s += fmt.Sprintf("Error: %v", err)
		c <- s // sending into the channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> status code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			file := strings.Split(url, "//")[1] // take second element (example http://www.google.com -> www.google.com)
			file += ".html"

			s += fmt.Sprintf("writing response body to %s\n", file)

			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}
}

func main() {
	// this program will save body responses concurrently for all specified urls
	urls := []string{"https://go.dev", "https://www.google.com", "https://www.medium.com"}

	// step 1
	c := make(chan string)

	for _, v := range urls {
		go func(v string) {
			checkAndSaveBody(v, c)
			fmt.Println(strings.Repeat("#", 20))
		}(v)
	}

	fmt.Println("number of goroutines:", runtime.NumGoroutine()) // 4 goroutines

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
