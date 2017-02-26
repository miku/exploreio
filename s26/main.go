// S26: A slow reader.

package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// SlowReader is a reader that throttles reads.
type SlowReader struct {
	r        io.Reader
	throttle <-chan time.Time
}

// NewReader creates a new SlowReader which inserts a given delay between each byte it emits.
func NewReader(r io.Reader, dur time.Duration) *SlowReader {
	return &SlowReader{r: r, throttle: time.Tick(dur)}
}

// Read will insert a slow delay into the reading process. At most one byte will
// be returned, independent of the size of p.
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
	slow := NewReader(strings.NewReader(s), 50*time.Millisecond)
	if _, err := io.Copy(os.Stdout, slow); err != nil {
		log.Fatal(err)
	}
}
