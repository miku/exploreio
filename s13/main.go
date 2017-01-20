// S13: Write to more than one writer at once.
//
//     $ rm -f hello.txt && go run main.go
//     SPQR
//
//     $ cat hello.txt
//     SPQR
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// TODO: Write to both, the file and stdout. 4 lines (incl. error handling).
	// ...
	// ...
	// ...
	// ...
}
