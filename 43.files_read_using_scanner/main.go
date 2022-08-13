package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords) // it will scan the file word by word (default is line by line)

	success := scanner.Scan()
	if !success {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("First line found:", scanner.Bytes())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
