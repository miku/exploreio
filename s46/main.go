// S46: Append a final newline to a bytestream.
//
// Via: https://git.io/vyblf
//
// Without using finalNewlineReader, the program is sensitive to the final newline:
//
// OUTPUT:
//
//     $ echo -n "text without newline" | go run main.go
//     $ echo "text without newline" | go run main.go
//     text without newline
//
// When using finalNewlineReader, the program gets a bit more robust, by being
// independent of the explicit newline at the end:
//
// OUTPUT:
//
//     $ echo -n "text without newline" | go run main.go
//     text without newline
//     $ echo "text without newline" | go run main.go
//     text without newline
//
// This is a small difference, but sometimes important, when working with newline delimited
// data files, where the final newline might be missing.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// finalNewlineReader appends a final newline to a byte stream, but only if there is not already one.
type finalNewlineReader struct {
	r    io.Reader
	done bool // true, when r has been fully read
}

// Read delegates reads to the wrapped reader. Only when EOF is encountered,
// check, if the last byte is a newline. If it is not, do not signal EOF just yet.
func (r *finalNewlineReader) Read(p []byte) (n int, err error) {
	if r.done {
		if len(p) > 0 {
			p[0] = 10
			return 1, io.EOF
		}
		return 0, nil
	}
	n, err = r.r.Read(p)
	if err == io.EOF && (n == 0 || p[n-1] != 10) {
		r.done = true
		return n, nil
	}
	return
}

func main() {
	r := os.Stdin // will print out string, only if it ends with a newline
	// r := &finalNewlineReader{r: os.Stdin} // appends newline, if missing
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(line)
	}
}
