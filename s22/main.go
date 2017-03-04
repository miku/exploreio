// S22: A writer that discards everything that is written to it.
//
// OUTPUT:
//
//     $ echo "Hello" | go run main.go
//
package main

import (
	"io"
	"log"
	"os"
)

// TODO: Implement type Discard, which throws away
//       everything that is written to it (4 lines).
// ...

// ...
// ...
// ...

func main() {
	if _, err := io.Copy(&Discard{}, os.Stdin); err != nil {
		log.Fatal(err)
	}
}

// Standard library implementation:
// https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go#L158
