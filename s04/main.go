// S04: Decompressors are filters can be chained.
//
//     $ cat hello.txt.gz | go run main.go
//     Don't just check errors, handle them gracefully.
package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	// TODO: Read compressed input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (7 lines, including error handling).
	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
