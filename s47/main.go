// S47: Two bufio.Reader, that can you can switch between.
//
// OUTPUT:
//
//     $ go run main.go
//     A
//     1
//     B
//     2
//     C
//     3
//     4
//     5

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

// ToggleReader can flip two readers.
type ToggleReader struct {
	*bufio.Reader
	sync.Mutex
	a, b *bufio.Reader
}

// Toggle switches to the other reader.
func (r *ToggleReader) Toggle() {
	r.Lock()
	defer r.Unlock()
	if r.Reader == r.a {
		r.Reader = r.b
	} else {
		r.Reader = r.a
	}
}

// NewToggleReader creates a new reader.
func NewToggleReader(a, b io.Reader) *ToggleReader {
	ba := bufio.NewReader(a)
	bb := bufio.NewReader(b)
	return &ToggleReader{ba, sync.Mutex{}, ba, bb}
}

func main() {
	a := strings.NewReader("A\nB\nC\n")
	b := strings.NewReader("1\n2\n3\n4\n5\n")

	tr := NewToggleReader(a, b)
	done := 0

	for {
		line, err := tr.ReadString('\n')
		if err == io.EOF {
			done++
			if done == 2 {
				break
			}
			tr.Toggle()
			continue
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(line)

		if done == 0 {
			tr.Toggle()
		}
	}
}
