// S15: Hello Buffer.
//
// > In computer science, a data buffer (or just buffer) is a region of a physical
// memory storage used to temporarily store data while it is being moved from one
// place to another. https://en.wikipedia.org/wiki/Data_buffer
//
// OUTPUT:
//
//     $ go run main.go
//     61
//     62
//     63
//     2e
//     78
//     79
//     7a

package main

import (
	"bytes"
	"log"
)

func main() {
	var buf bytes.Buffer
	if _, err := buf.WriteString("abc.xyz"); err != nil {
		log.Fatal(err)
	}
	// TODO: Read one byte at a time from the buffer
	//       and print the hex value on stdout (10 lines).
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
}
