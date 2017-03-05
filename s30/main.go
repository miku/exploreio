// S30: A small buffer.
//
// OUTPUT:
//
//     $ echo -n "Hello Buffer" | go run main.go
//     2017/03/04 13:48:22 12 bytes read
//     Hello Buffer
//
// Questions:
//
// (1) This implementation suffers from a small but serious flaw.
//     Can you spot it?
// (2) Can you implement a more efficient version?
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Buffer is a minimal implementation of a variable size byte slice one can read
// from and write into.
type Buffer struct {
	b    []byte
	i, j int
}

// Len returns the current length of this buffer.
func (b *Buffer) Len() int {
	return b.j - b.i
}

// Cap returns the current capacity of this buffer.
func (b *Buffer) Cap() int {
	return len(b.b)
}

// Write writes given byte slice to the buffer.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.b = append(b.b, p...)
	b.j += len(p)
	return len(p), nil
}

// Read reads the current buffer.
func (b *Buffer) Read(p []byte) (n int, err error) {
	sz := len(b.b) - b.i
	if sz > len(p) {
		sz = len(p)
	} else {
		err = io.EOF
	}
	copy(p, b.b[b.i:b.i+sz])
	b.i += sz
	return sz, err
}

func main() {
	var buf Buffer
	if _, err := io.Copy(&buf, os.Stdin); err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes read", buf.Len())
	b, err := ioutil.ReadAll(&buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
