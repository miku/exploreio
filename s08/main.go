// S08: Strings can be readers, io.ReadFull.
//
//     $ go run main.go
//     Strings
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader(`Strings can be readers, too.`)
	// TODO: Read the first 7 bytes of the string into buf, the print to stdout. 5 lines.
	b := make([]byte, 7)
	io.ReadFull(r, b)
	fmt.Println(string(b))
}
