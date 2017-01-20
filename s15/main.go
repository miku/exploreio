// S15: Hello Buffer.
//
// > In computer science, a data buffer (or just buffer) is a region of a physical
// memory storage used to temporarily store data while it is being moved from one
// place to another. /wiki/Data_buffer
//
//     $ go run main.go
//      31
//      32
//      33

package main

import (
	"bytes"
	"log"
)

func main() {
	var buf bytes.Buffer
	if _, err := buf.WriteString("123"); err != nil {
		log.Fatal(err)
	}
	// TODO: Read one byte at a time from buffer and print the hex value on stdout. 10 lines (incl. error handling).
}
