// S22: A reader that discards everything that is written to it.
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
