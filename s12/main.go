// S12: Read from multiple readers in turn.
//
// OUTPUT:
//
//     $ go run main.go
//     Hello
//     Gopher
//     World
//     !
package main

import (
	"io"
	"strings"
)

func main() {
	// TODO: Read from these four readers
	//       and write to standard output (4 lines).
	rs := []io.Reader{
		strings.NewReader("Hello\n"),
		strings.NewReader("Gopher\n"),
		strings.NewReader("World\n"),
		strings.NewReader("!\n"),
	}
	// ...
	// ...
	// ...
	// ...
}
