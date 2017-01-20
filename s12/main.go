// S12: Read from multiple readers in turn.
//
//     $ go run main.go
//     Hello
//     Gopher
//     World
//     !
package main

import "strings"

func main() {
	// TODO: Read from these four readers and write to stdout. 4 lines (incl. 1 long and err handling).
	strings.NewReader("Hello\n")
	strings.NewReader("Gopher\n")
	strings.NewReader("World\n")
	strings.NewReader("!\n")
}
