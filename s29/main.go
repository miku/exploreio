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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const DefaultTimeout = 100 * time.Millisecond

// RoundRobin reads from readers round robin until all are exhausted.
type RoundRobin struct {
	rs       []*bufio.Reader
	delim    byte
	cur      int
	buf      bytes.Buffer
	maxRetry int
}

// New creates a new reader.
func New(rs ...io.Reader) *RoundRobin {
	rr := &RoundRobin{delim: '\n', maxRetry: 10}
	for _, r := range rs {
		tr := &TimeoutReader{r: r, timeout: DefaultTimeout}
		rr.rs = append(rr.rs, bufio.NewReader(tr))
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
		var i int
		for {
			if i == r.maxRetry {
				return 0, fmt.Errorf("max retries exceeded")
			}
			i++
			if err := r.fill(); err != nil {
				if err == ErrTimeout {
					// switch to next reader
					log.Println("retrying ")
					r.cur = (r.cur + 1) % len(r.rs)
					continue
				}
				return 0, err
			}
			break
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

// Slow is a sleepy reader.
type Slow struct {
	Sleep time.Duration
}

func (r *Slow) Read(p []byte) (n int, err error) {
	time.Sleep(r.Sleep)
	copy(p, []byte{0x7a, 0x7a, 0x7a, 0x7a, 0x0a})
	return len(p), io.EOF
}

func main() {
	var rs []io.Reader
	for i := 0; i < 100; i++ {
		rs = append(rs, strings.NewReader(fmt.Sprintf("%d\n", i)))
		rs = append(rs, &Slow{Sleep: 1000 * time.Millisecond})
	}
	rr := New(rs...)
	if _, err := io.Copy(os.Stdout, rr); err != nil {
		log.Fatal(err)
	}
}
