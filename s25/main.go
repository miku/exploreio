// S25: Generate data.
//
// OUTPUT:
//
//     $ go run main.go | head -10
//     2017-26-02 19:23:03    1.6047
//     2017-26-02 19:23:03    2.2692
//     2017-26-02 19:23:03    1.8446
//     2017-26-02 19:23:03    1.9102
//     2017-26-02 19:23:03    1.8133
//     2017-26-02 19:23:03    1.2980
//     2017-26-02 19:23:03    1.5123
//     2017-26-02 19:23:03    1.1942
//     2017-26-02 19:23:03    0.9112
//     2017-26-02 19:23:03    0.2321
//     ...

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// linesToFill is the numnber of lines generated for the internal buffer.
const linesToFill = 1000

// EndlessStream generates a stream of endless data. It uses an internal buffer to
// decouple data production and consumption. When the buffer is empty, we first
// fill the buffer (with linesToFill lines of data), then pass control back to the
// Read method, which then drains the buffer with each call to Read.
type EndlessStream struct {
	buf   bytes.Buffer
	cur   time.Time
	value float64
}

// Read will provide the reader with additional data. This reader will never signal an EOF.
func (r *EndlessStream) Read(p []byte) (n int, err error) {
	if r.buf.Len() == 0 {
		if err := r.fill(); err != nil {
			return 0, err
		}
	}
	lr := io.LimitReader(&r.buf, int64(len(p)))
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		return 0, err
	}
	copy(p, b)
	return len(b), nil
}

// fill is a helper method to fill up the internal buffer with the actual data,
// that is then read via Read.
func (r *EndlessStream) fill() error {
	if r.cur.IsZero() {
		r.cur = time.Now()
	}
	if r.value == 0 {
		r.value = 1 + rand.Float64()
	}
	for i := 0; i < linesToFill; i++ {
		if _, err := fmt.Fprintf(&r.buf, "%s\t%0.4f\n",
			r.cur.Format("2006-02-01 15:04:05.999"),
			r.value); err != nil {
			return err
		}
		if rand.Float64() > 0.50 {
			r.value += rand.Float64()
		} else {
			r.value -= rand.Float64()
		}
		r.cur = r.cur.Add(1 * time.Millisecond)
	}
	return nil
}

func main() {
	pr := &EndlessStream{}
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}
