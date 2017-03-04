// S42: Callbacks.
//
// OUTPUT:
//
//     $ cat main.go | go run main.go
//     ...
//     package main
//
//     import (
//       "io"
//       ...
//       }
//     }
//     2017/01/21 14:10:28 done reading

package main

import (
	"io"
	"log"
	"os"
)

type onEOFreader struct {
	r io.Reader
	f func()
}

func (r *onEOFreader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err == io.EOF {
		r.f()
	}
	return n, err
}

func main() {
	r := onEOFreader{r: os.Stdin, f: func() {
		log.Printf("done reading")
	}}
	if _, err := io.Copy(os.Stdout, &r); err != nil {
		log.Fatal(err)
	}
}
