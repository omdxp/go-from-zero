package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	scanner.Scan() // scan a new line

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text:", text)
	fmt.Println("Input bytes:", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)
		if strings.EqualFold(text, "exit") {
			fmt.Println("Exiting the scanning...")
			break
		}
	}

	fmt.Println("Just a message after for loop")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
