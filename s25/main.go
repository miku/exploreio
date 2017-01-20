// S25: Generate data.
//
//     $ go run main.go | head -10
//     1/20/2017 23:45:29, 1.6047
//     1/20/2017 23:45:29, 2.2692
//     1/20/2017 23:45:29, 1.8446
//     1/20/2017 23:45:29, 1.9102
//     1/20/2017 23:45:29, 1.8133
//     1/20/2017 23:45:29, 1.2980
//     1/20/2017 23:45:29, 1.5123
//     1/20/2017 23:45:29, 1.1942
//     1/20/2017 23:45:29, 0.9112
//     1/20/2017 23:45:29, 0.2321
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

type LastPrice struct {
	buf   bytes.Buffer
	cur   time.Time
	price float64 // bad idea
}

func (r *LastPrice) Read(p []byte) (n int, err error) {
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
	return len(p), nil
}

func (r *LastPrice) fill() error {
	if r.cur.IsZero() {
		r.cur = time.Now()
	}
	if r.price == 0 {
		r.price = 1 + rand.Float64()
	}
	for i := 0; i < 100; i++ {
		if _, err := fmt.Fprintf(&r.buf, "%s, %0.4f\n", r.cur.Format("1/2/2006 15:04:05"), r.price); err != nil {
			return err
		}
		if rand.Float64() > 0.50 {
			r.price += rand.Float64()
		} else {
			r.price -= rand.Float64()
		}
		r.cur = r.cur.Add(1 * time.Millisecond)
	}
	return nil
}

func main() {
	pr := &LastPrice{}
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}
