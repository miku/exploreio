// S00: A file is a reader.
//
// It has been asked, how to get an reader from a file:
//
// "golang: Create a io.Reader from a local file"
// http://stackoverflow.com/questions/25677235/golang-create-a-io-reader-from-a-local-file
//
// Indeed, *os.File implements a Read methods with the approriate signature, so *os.File is an io.Reader.
//
//     $ go run main.go
//     31 bytes read: Concurrency is not parallelism.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	var r io.Reader = file

	b := make([]byte, 11)
	n, err := r.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes read: %s\n", n, string(b))
	// Output: 11 bytes read: Concurrency
}

// TODO: Change a single character so the complete file is read and printed.
