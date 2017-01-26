// S22: A writer that discards everything that is written to it.
//
// $ echo "Hello" | go run main.go
package main

import (
	"io"
	"log"
	"os"
)

// TODO: Implement Discard, that throws away everything that is written. 4 lines.
// ...
type Discard struct{}

// TODO: 0 or len(p)
func (w *Discard) Write(p []byte) (n int, err error) {
	return 0, nil
}

// ...
// ...
// ...

func main() {
	if _, err := io.Copy(&Discard{}, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

// Stdlib implementation:
// https://github.com/golang/go/blob/ad26bb5e3098cbfd7c0ad9a1dc9d38c92e50f06e/src/io/ioutil/ioutil.go#L158
