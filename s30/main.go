// S30: A small buffer.
//
// $ go run main.go
// Hello
// Golab 2017
package main

import (
	"io"
	"log"
	"os"
)

// Buffer minimal.
type Buffer struct {
	b    []byte
	i, j int
}

// Len TODO(miku): Check.
func (b *Buffer) Len() int {
	return b.j - b.i
}

// Cap TODO(miku): Check.
func (b *Buffer) Cap() int {
	return len(b.b)
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.b = append(b.b, p...)
	b.j += len(p)
	return len(p), nil
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	sz := len(b.b) - b.i
	if sz > len(p) {
		sz = len(p)
	} else {
		err = io.EOF
	}
	copy(p, b.b[b.i:])
	b.i += sz
	return sz, err
}

func main() {
	var buf Buffer
	if _, err := io.WriteString(&buf, "\n░♡░┳┣I░A░Ⓝ░К▒❀▒УOЦ░♡░ for your attention!\n\n"); err != nil {
		log.Fatal(err)
	}
	r := io.LimitReader(&buf, 6)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, &buf); err != nil {
		log.Fatal(err)
	}
}
