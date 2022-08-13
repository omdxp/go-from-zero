package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // execute this instruction as the last one before main returns

	byteSlice := []byte("I learn Golang!")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	bs := []byte("Go programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644) // if the file does exist it'll truncate it before writing the bytes
	if err != nil {
		log.Fatal(err)
	}

}

// defere statement postpone the exectution of a statement until the surrounding function returns or throw a panic
// os can be used only to write files
// other packages like io, ioutil, bufio does the same with simpler instructions
// but they add a lot of functionalities as well, thus increasing the size of the executable which is staticly linked binary
