// S41: Can we read concurrently from a reader?
//
//     $ go run main.go  | head -10
//     0
//     1
//     2
//     3
//     4
//     5
//     6
//     7
//     8
//     9
//
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

// Numbers emits natural numbers, one in a row. Inefficient.
type Numbers struct {
	N   int64
	buf bytes.Buffer
	mu  sync.Mutex
}

func (r *Numbers) Read(p []byte) (n int, err error) {
	// If we do not use a mutex, we get the following panic:
	//
	// $ go run main.go
	// 0
	// 1
	// ...
	// 7
	// panic: bytes.Buffer: truncation out of range [recovered]
	// 	panic: bytes.Buffer: truncation out of range
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.buf.Len() == 0 {
		if err := r.fill(); err != nil {
			return 0, err
		}
	}

	b, err := ioutil.ReadAll(io.LimitReader(&r.buf, int64(len(p))))
	if err != nil {
		return len(b), err
	}
	copy(p, b)
	return len(b), nil
}

func (r *Numbers) fill() error {
	if r.buf.Len() > 0 {
		return nil
	}
	if _, err := io.WriteString(&r.buf, fmt.Sprintf("%d\n", r.N)); err != nil {
		return err
	}
	atomic.AddInt64(&r.N, 1)
	return nil
}

func run(r io.Reader) {
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := Numbers{}
	go run(&r)
	run(&r)
}
