// S44: A flaky reader that flips bytes.
//
//     $ go run main.go
//     Hello World!
//     Hello World!
//     Hello World!
//     Hello Wosld"
//     Hellp Xorld!
//     Hello World!
//     Hemlp World!
//     Helmo World!
//     Hello World!
//     Hello World"

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Flaky struct {
	r    io.Reader
	prob float64
}

func (r Flaky) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := range p {
		if rand.Float64() < r.prob {
			p[i] = (p[i] + 1) % 255
		}
	}
	return n, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		flaky := &Flaky{r: strings.NewReader("Hello World!"), prob: 0.1}
		b, _ := ioutil.ReadAll(flaky)
		fmt.Println(string(b))
	}
}
