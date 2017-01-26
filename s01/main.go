// S01: Get bytes from a reader.
//
// Sometimes you want to consume a reader all at once.
//
//     $ go run main.go
//     Syscall must always be guarded with build tags.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: We don't need to loop manually, there is a helper function for that.
	// TODO: Replace the next 10 lines with 5 that do the same.
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
