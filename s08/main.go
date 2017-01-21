// S08: Strings can be readers, io.ReadFull.
//
//     $ go run main.go
//     Strings
package main

import "strings"

func main() {
	r := strings.NewReader(`Strings can be readers, too.`)
	// TODO: Read the first 7 bytes of the string into buf, the print to stdout. 5 lines.
}
