// S26: A slow reader.

package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type SlowReader struct {
	r        io.Reader
	throttle <-chan time.Time
}

func New(r io.Reader, dur time.Duration) *SlowReader {
	return &SlowReader{r: r, throttle: time.Tick(dur)}
}

func (r *SlowReader) Read(p []byte) (n int, err error) {
	<-r.throttle
	if len(p) == 0 {
		return 0, nil
	}
	b := make([]byte, 1)
	n, err = r.r.Read(b)
	if err != nil {
		return
	}
	copy(p, b)
	return
}

func main() {
	s := `Among the primitive concepts of computer programming, and of the high level
languages in which programs are expressed, the action of assignment is familiar
and well understood. In fact, any change of the internal state of a machine
executing a program can be modeled as an assignment of a new value to some
variable part of that machine. However, the operations of input and output,
which affect the external environment of a machine, are not nearly so well
understood. They are often added to a programming language only as an
afterthought.
`
	r := strings.NewReader(s)
	sr := New(r, 50*time.Millisecond)
	if _, err := io.Copy(os.Stdout, sr); err != nil {
		log.Fatal(err)
	}
}
