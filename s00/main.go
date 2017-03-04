// S00: A file is a reader.
//
// It has been asked, how to get an reader from a file:
//
//     "golang: Create a io.Reader from a local file"
//         http://stackoverflow.com/q/25677235/89391
//
//     > asked 2014-09-05, 01:05:30Z, viewed 16338 times
//
// Indeed, *os.File implements a `Read` method with the approriate signature,
// so *os.Fileis an io.Reader.
//
// OUTPUT:
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

// TODO: Change a single character in this program so the complete
//       file is read and printed.
func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// A file is an io.Reader.
	var r io.Reader = file

	b := make([]byte, 11)
	n, err := r.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes read: %s\n", n, string(b))
}
