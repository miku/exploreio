// S16: Reading the output of a shell command into a buffer.
//
// OUTPUT:
//
//     $ go run main.go
//     command output has 8 bytes: main.go
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
	// TODO: Stream output of command into the buffer (4 lines).
	// ...
	// ...
	// ...
	// ...
	fmt.Printf("command output has %d bytes: %s", buf.Len(), buf.String())
}
