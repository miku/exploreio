// S02: Copying and standard streams.
//
// You don't need to use byte slices explicitly, if you don't need them.
//
// OUTPUT:
//
//     $ go run main.go
//     The bigger the interface, the weaker the abstraction.
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Write contents of the file to the standard output,
	//       without using a byte slice (3 lines).
	// ...
	// ...
	// ...
}
