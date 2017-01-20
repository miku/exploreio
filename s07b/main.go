// S07b: Read sections of a reader.
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
	r := strings.NewReader("some io.Reader stream to be read\n")

	// TODO: Print the string "io.Reader" to stdout. 4 lines.
	s := io.NewSectionReader(r, 5, 9)
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
