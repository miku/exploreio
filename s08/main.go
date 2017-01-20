// S08: Strings can be readers, io.ReadFull.
//
//     $ go run main.go
//     Strings
package main

import "strings"

func main() {
	r := strings.NewReader(`Strings can be readers, too.`)

}
