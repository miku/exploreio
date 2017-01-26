// S03: Stdin is a file, too.
//
// A filter that does nothing.
//
//     $ cat hello.txt | go run main.go
//     Cgo is not Go.
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// TODO: Read input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (3 lines, including error handling).
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}
