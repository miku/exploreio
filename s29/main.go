// S29: Round robin multireader, that can handle broken readers.

//     $ go run main.go
//     0
//     2017/01/21 00:46:03 retrying
//     1
//     2017/01/21 00:46:03 retrying
//     2
//     ...
//     2017/01/21 00:46:21 retrying
//     98
//     2017/01/21 00:46:21 retrying
//     99
//     2017/01/21 00:46:21 retrying
//     2017/01/21 00:46:21 retrying
//     2017/01/21 00:46:21 retrying
//     2017/01/21 00:46:21 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 retrying
//     2017/01/21 00:46:22 max retries exceeded
//     exit status 1

package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	// DefaultTimeout for the a reader.
	DefaultTimeout = 100 * time.Millisecond

	// DefaultMaxRetry is default number of read that may fail in a row.
	DefaultMaxRetry = 3
)

// RoundRobin reads from readers round robin until all are exhausted. If a reader
// does not respond in time, if will be retried a given number of times. If this reader
// encounters maxRetry failed reads in a row, it will finally give up.
type RoundRobin struct {
	rs       []*bufio.Reader
	delim    byte
	cur      int
	buf      bytes.Buffer
	maxRetry int
}

// NewRoundRobinReader creates a new reader.
func NewRoundRobinReader(rs ...io.Reader) *RoundRobin {
	rr := &RoundRobin{delim: '\n', maxRetry: DefaultMaxRetry}
	for _, r := range rs {
		tr := &TimeoutReader{r: r, timeout: DefaultTimeout}
		rr.rs = append(rr.rs, bufio.NewReader(tr))
	}
	return rr
}

// Read reads from the current reader until delim or EOF is reached. If it's EOF,
// remove the reader for the list.
func (r *RoundRobin) Read(p []byte) (n int, err error) {
	if r.buf.Len() == 0 {
		if len(r.rs) == 0 {
			// Neither buffer nor readers to read from.
			log.Println("Read: successfully read from all readers")
			return 0, io.EOF
		}
		// There are still active readers.
		var i int
		for {
			if i == r.maxRetry {
				return 0, fmt.Errorf("max retries (%d) exceeded", r.maxRetry)
			}
			i++
			if err := r.fill(); err != nil {
				if err != ErrTimeout {
					return 0, err
				}
				// Timeout: Switch to next reader.
				log.Printf("switching to next reader ...")
				r.cur = (r.cur + 1) % len(r.rs)
				continue
			}
			break
		}
	}
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
		return nil
	}
	b, err := r.rs[r.cur].ReadBytes(r.delim)
	if err != nil {
		if err != io.EOF {
			return err
		}
		r.rs = append(r.rs[:r.cur], r.rs[r.cur+1:]...)
		log.Printf("fill: removed exhausted reader, readers left: %d", len(r.rs))
	}
	if _, err := r.buf.Write(b); err != nil {
		return err
	}
	if len(r.rs) > 0 {
		r.cur = (r.cur + 1) % len(r.rs)
	}
	return nil
}

// ErrTimeout signals a timeout.
var ErrTimeout = errors.New("timeout")

// TimeoutReader times out, if read takes too long. https://github.com/golang/go/wiki/Timeouts
type TimeoutReader struct {
	r       io.Reader
	timeout time.Duration
}

// readResult wraps result of a Read.
type readResult struct {
	b   []byte
	err error
}

// Read behaves as usual, except it returns an ErrTimeout if Read takes too long.
func (r *TimeoutReader) Read(p []byte) (n int, err error) {
	ch := make(chan readResult, 1)
	go func() {
		pp := make([]byte, len(p))
		_, err := r.r.Read(pp)
		ch <- readResult{pp, err}
	}()
	select {
	case <-time.After(r.timeout):
		return 0, ErrTimeout
	case res := <-ch:
		copy(p, res.b)
		return len(res.b), res.err
	}
}

// SlowAndFlaky is a sleepy, flaky reader.
type SlowAndFlaky struct {
	ID    int
	Sleep time.Duration
}

func (r *SlowAndFlaky) Read(p []byte) (n int, err error) {
	if rand.Float64() > 0.5 {
		time.Sleep(r.Sleep)
	}
	copy(p, []byte(fmt.Sprintf("SlowAndFlaky #%d\n", r.ID)))
	return len(p), io.EOF
}

func main() {
	n := flag.Int("n", 100, "number of each good and flaky readers")
	flag.Parse()

	var rs []io.Reader
	for i := 0; i < *n; i++ {
		rs = append(rs, strings.NewReader(fmt.Sprintf("Reader #%d\n", i)))
		rs = append(rs, &SlowAndFlaky{ID: i, Sleep: 1000 * time.Millisecond})
	}
	rr := NewRoundRobinReader(rs...)
	if _, err := io.Copy(os.Stdout, rr); err != nil {
		log.Fatal(err)
	}
}
