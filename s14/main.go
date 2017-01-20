// S14: Read into variables.
//
//     $ echo 1 1 1 | go run main.go
//     1 (int), 1.0000 (float64), 1 (string)
//
//     $ echo a 1 1 | go run main.go
//     2017/01/20 09:08:18 expected integer
//     exit status 1
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		i int
		f float64
		s string
	)
	// TODO: Read an int, a float and a string from stdin. 3 lines.
	if _, err := fmt.Fscan(os.Stdin, &i, &f, &s); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d (%T), %0.4f (%T), %s (%T)\n", i, i, f, f, s, s)
}
