// S21: A reader that converts all Unicode letter mapped to their upper case.
//
//     $ echo "Hello Gophers" | go run main.go
//     HELLO GOPHERS
package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

// TODO: Implement UpperReader, a reader that converts all Unicode letter mapped to their upper case. 11 lines.
type UpperReader struct {
	r io.Reader
}

func (r *UpperReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	copy(p, bytes.ToUpper(p))
	return
}

func main() {
	_, err := io.Copy(os.Stdout, &UpperReader{r: os.Stdin})
	if err != nil {
		log.Fatal(err)
	}
}
