// S27b: A reader that times out.

//     $ go run main.go
//     Hello World!
//     2017/01/21 00:42:59 timeout
//     exit status 1

package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var ErrTimeout = errors.New("timeout")

// TimeoutReader times out, if read takes too long. https://github.com/golang/go/wiki/Timeouts
type TimeoutReader struct {
	r       io.Reader
	timeout time.Duration
}

// TODO: Implement a reader that times out after a certain a given timeout. 19 lines.
// ...
// ...
// ...
// ...

// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...
// ...

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
	var r io.Reader

	// Strings are fast.
	r = &TimeoutReader{
		r:       strings.NewReader("Hello World!\n"),
		timeout: 100 * time.Millisecond,
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	r = &TimeoutReader{
		r:       &Slow{Sleep: 100 * time.Millisecond},
		timeout: 70 * time.Millisecond,
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	// Hello World!
	// 2017/01/19 19:56:29 timeout
	// exit status 1
}
