// S01: Get bytes from a reader.
//
// Sometimes you want to consume a reader all at once.
//
// OUTPUT:
//
//     $ go run main.go
//     Syscall must always be guarded with build tags.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: We don't need to loop manually, there is a helper
	//       function for that. Replace the next 10 lines
	//       with 5 lines that do the same.
	var contents []byte
	for {
		b := make([]byte, 8)
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		contents = append(contents, b...)
	}
	fmt.Println(string(contents))
}
