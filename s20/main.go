// S20: The Reader interface.
//
// OUTPUT:
//
//     $ go run main.go
//     0
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Empty is an reader that returns nothing.
type Empty struct{}

// TODO: Implement io.Reader interface. Always return EOF (3 lines).
// ...
// ...
// ...

func main() {
	n, err := io.Copy(os.Stdout, &Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n) // Print the number of bytes read.
}
