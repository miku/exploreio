// S21: A reader that converts all unicode letter mapped to their upper case.
//
//     $ echo "Hello Gophers" | go run main.go
//     HELLO GOPHERS
package main

import (
	"io"
	"log"
	"os"
)

// UpperReader is an uppercase filter.
type UpperReader struct {
	r io.Reader
}

// TODO: Implement UpperReader, a reader that converts
//       all Unicode letter mapped to their upper case (6 lines).
func (r *UpperReader) Read(p []byte) (n int, err error) {
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
}

func main() {
	_, err := io.Copy(os.Stdout, &UpperReader{r: os.Stdin})
	if err != nil {
		log.Fatal(err)
	}
}
