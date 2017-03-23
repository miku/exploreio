// S08: Strings can be readers, io.ReadFull.
//
// OUTPUT:
//
//     $ go run main.go
//     Strings
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader(`Strings can be readers, too.`)
	// TODO: Read the first 7 bytes of the string into a byte slice,
	//       then print to stdout (5 lines).
}
