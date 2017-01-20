// S16: Reading the output of a shell command into a buffer.
//
//     $ go run main.go
//     output has 8 bytes: 6d 61 69 6e 2e 67 6f 0a
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var buf bytes.Buffer
	cmd := exec.Command("ls")
	// TODO: Stream output of command directly into the buffer.
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output has %d bytes: %s\n", buf.Len(), buf.String())
}
