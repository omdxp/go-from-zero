package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2) // make 2 bytes slice

	numberBytesRead, err := io.ReadFull(file, byteSlice) // read first 2 bytes in file
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read all data from a file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))

	// another way to read data from file
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data read: %s\n", data)
}
