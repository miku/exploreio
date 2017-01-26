// S02: Copying and standard streams.
//
// You don't need to use byte slices, if you don't need them.
//
//     $ go run main.go
//     The bigger the interface, the weaker the abstraction.
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Write output to Stdout, without using a byte slice (3 lines, including error handling).
	// ...
	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
}
