// S28: Round robin multi-reader.
//
// OUTPUT:
//
//     $ go run main.go
//     0
//     1
//     2
//     3
//     0
//     1
//     2
//     3
//
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

// RoundRobin reads from readers round robin until all readers are exhausted.
type RoundRobin struct {
	rs    []*bufio.Reader // the readers
	delim byte            // delimiter on which to stop reading, e.g. a newline
	cur   int             // index of currently active reader
	buf   bytes.Buffer    // internal buffer
}

// NewReader creates a new reader. We use a newline as a delimiter by default.
func NewReader(rs ...io.Reader) *RoundRobin {
	rr := &RoundRobin{delim: '\n'}
	for _, r := range rs {
		// We use a bufio.Reader, so we can use r.ReadBytes later.
		rr.rs = append(rr.rs, bufio.NewReader(r))
	}
	return rr
}

// Read reads from the current reader until delim or EOF is reached. If it's EOF,
// remove the reader for the list.
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
	// Read at most len(p) bytes from the internal buffer. Fewer is ok, too.
	b, err := ioutil.ReadAll(io.LimitReader(&r.buf, int64(len(p))))
	if err != nil {
		return len(b), err
	}
	copy(p, b)
	return len(b), nil
}

// fill fills the buffer that will be drained by Read.
func (r *RoundRobin) fill() error {
	if r.buf.Len() > 0 {
		// If the buffer is not empty yet, there is no need to fill it up.
		return nil
	}
	// Read from the current reader until we hit the delimiter.
	b, err := r.rs[r.cur].ReadBytes(r.delim)
	if err != nil {
		if err != io.EOF {
			// An error occured and it's not EOF, report.
			return err
		}
		// Remove the exhausted reader from the list of readers
		// (https://github.com/golang/go/wiki/SliceTricks).
		r.rs = append(r.rs[:r.cur], r.rs[r.cur+1:]...)
	}
	// Writer bytes into the internal buffer.
	if _, err := r.buf.Write(b); err != nil {
		return err
	}
	// Move reader index to the next reader.
	if len(r.rs) > 0 {
		r.cur = (r.cur + 1) % len(r.rs)
	}
	return nil
}

func main() {
	var rs []io.Reader
	for i := 0; i < 4; i++ {
		rs = append(rs, strings.NewReader(fmt.Sprintf("reader #%d\nreader #%d\n", i, i)))
	}
	rr := NewReader(rs...)
	if _, err := io.Copy(os.Stdout, rr); err != nil {
		log.Fatal(err)
	}
}
