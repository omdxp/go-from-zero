package main

import "fmt"

const numberOfSecondsInHour = 3600

func main() {
	fmt.Println("number of seconds in hour is", numberOfSecondsInHour)
}

// go run: run the application immediately (go run main.go)
// go build: build the application to a single executable file named after the folder that resides in it
// GOOS=X GOARCH=Y go build: build the executable for a specific OS and architecture
// go install: build the executable and move it to Go Workspace under go/bin/ folder so it can be accessed anywhere
// gofmt -w main.go: format the file in idiomatic way as in the standard style for Go code
// gofmt -w -l folder/: format all files under folder/
// go fmt: format all Go files in the current directory
