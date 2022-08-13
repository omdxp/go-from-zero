package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) // this is the idiomatic way to handle errors in Go
	}

	err = os.Truncate("a.txt", 0) // 0 means completly empty the file (for a specified number it'll truncate the file to that amount of bytes)
	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644) // creates file if it doesn't exist or append to file if it exists
	if err != nil {
		log.Fatal(err)
	}
	file.Close() // always close files when finishing

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	p := fmt.Println
	p("File name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is directory?", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	_, err = os.Stat("b.txt") // b.txt does not exist
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}

	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

}

// os.Create() will create a file if not existed, otherwise it'll truncate that file
// log is thread safe and adds timing information to the message automatically
// os.Open() is used to open a file
// os.OpenFile() is used to open a file with specfied options like mode and permissions
