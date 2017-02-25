// S07b: Read sections from a reader.
//
// OUTPUT:
//
//     $ go run main.go
//     io.Reader
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader to be read\n")

	// TODO: Print the string "io.Reader" to stdout (4 lines).
	// ...
	// ...
	// ...
	// ...
	fmt.Println() // for readability
}
