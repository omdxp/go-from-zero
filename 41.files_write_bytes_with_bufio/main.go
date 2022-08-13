package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"my_file.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// write to file
	err = bufferedWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// reset buffer for unflushed buffered data (there is no reason after being flushed)
	bufferedWriter.Reset(bufferedWriter)
}

// bufio is useful to manipulate data in memory before dumping it to disk, it is so efficient than writing directly the bytes to disk each time
