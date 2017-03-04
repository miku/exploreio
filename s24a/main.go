// S24a: A counting reader.
//
// OUTPUT:
//
//     $ cat main.go | go run main.go
//     n (io.Copy) = 775, n (CountingReader) = 775
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// TODO: Implement a reader that counts the total
//       number of bytes read.
//       It should have a Count() uint64 method,
//       that returns the number of bytes read so far.
//       (12 lines).
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
// ...
// ...

func main() {
	cr := &CountingReader{r: os.Stdin}
	n, err := io.Copy(ioutil.Discard, cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n (io.Copy) = %d, n (CountingReader) = %d\n", n, cr.Count())
}
