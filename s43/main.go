// S43: Like /dev/zero.
//
// OUTPUT:
//
//     $ go run main.go | xxd
//     00000000: 0000 0000 0000 0000 0000 0000 0000 0000  ................
//     00000010: 0000 0000 0000 0000 0000 0000 0000 0000  ................
//     00000020: 0000 0000 0000 0000 0000 0000 0000 0000  ................
//     00000030: 0000 0000 0000 0000 0000 0000 0000 0000  ................

package main

import (
	"io"
	"log"
	"os"
)

var devZero = zeroReader(0)

type zeroReader int

func (r zeroReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	r := io.LimitReader(devZero, 64)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
