// S13: Write to more than one writer at once.
//
// OUTPUT:
//
//     $ rm -f output.txt && go run main.go
//     SPQR
//
//     $ cat output.txt
//     SPQR
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// TODO: Write to both, the file and standard output (4 lines).
	// ...
	// ...
	// ...
	// ...
}
