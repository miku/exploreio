// S24a: A counting reader, count the number of bytes read in total.
//
//     $ cat main.go | go run main.go
//     n (io.Copy) = 550, n (CountingReader) = 550
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync/atomic"
)

// TODO: implement a reader that counts the total number of bytes read. 12 lines.
type CountingReader struct {
	r     io.Reader
	count uint64
}

func (r *CountingReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	atomic.AddUint64(&r.count, uint64(n))
	return
}

func (r *CountingReader) Count() uint64 {
	return atomic.LoadUint64(&r.count)
}

func main() {
	cr := &CountingReader{r: os.Stdin}
	n, err := io.Copy(ioutil.Discard, cr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n (io.Copy) = %d, n (CountingReader) = %d\n", n, cr.Count())
}
