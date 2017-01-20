// S28: Round robin multi-reader.

//     $ go run main.go
//     0
//     1
//     2
//     3
//     4
//     5
//     ...
//     97
//     98
//     99
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// RoundRobin reads from readers round robin until all are exhausted.
type RoundRobin struct {
	rs    []*bufio.Reader
	delim byte
	cur   int
	buf   bytes.Buffer
}

// New creates a new reader.
func New(rs ...io.Reader) *RoundRobin {
	rr := &RoundRobin{delim: '\n'}
	for _, r := range rs {
		rr.rs = append(rr.rs, bufio.NewReader(r))
	}
	return rr
}

// Read reads from the current reader until delim or EOF is reached. If it's EOF, remove the reader for the list.
func (r *RoundRobin) Read(p []byte) (n int, err error) {
	if r.buf.Len() == 0 {
		if len(r.rs) == 0 {
			// Neither buffer nor readers to read from.
			return 0, io.EOF
		}
		// There are still active readers.
		if err := r.fill(); err != nil {
			return 0, err
		}
	}
	b, err := limitBytes(&r.buf, int64(len(p)))
	if err != nil {
		return len(b), err
	}
	copy(p, b)
	return len(b), nil
}

// fill fills the buffer that will be drained by Read.
func (r *RoundRobin) fill() error {
	if r.buf.Len() > 0 {
		return nil
	}
	b, err := r.rs[r.cur].ReadBytes(r.delim)
	if err != nil {
		if err != io.EOF {
			return err
		}
		r.rs = append(r.rs[:r.cur], r.rs[r.cur+1:]...)
	}
	if _, err := r.buf.Write(b); err != nil {
		return err
	}
	if len(r.rs) > 0 {
		r.cur = (r.cur + 1) % len(r.rs)
	}
	return nil
}

// limitBytes reads at most n bytes from reader and returns them
func limitBytes(r io.Reader, n int64) ([]byte, error) {
	lr := io.LimitReader(r, n)
	return ioutil.ReadAll(lr)
}

func main() {
	var rs []io.Reader
	for i := 0; i < 100; i++ {
		rs = append(rs, strings.NewReader(fmt.Sprintf("%d\n", i)))
	}
	rr := New(rs...)
	if _, err := io.Copy(os.Stdout, rr); err != nil {
		log.Fatal(err)
	}
}
