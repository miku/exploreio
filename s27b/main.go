// S27b: A reader that times out.
//
// OUTPUT:
//
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

// ErrTimeout signals a timeout.
var ErrTimeout = errors.New("timeout")

// TimeoutReader times out, if read takes too long.
// https://github.com/golang/go/wiki/Timeouts
type TimeoutReader struct {
	r       io.Reader
	timeout time.Duration
}

// TODO: Implement a reader that times out after
//       a certain a given timeout (21 lines).
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
// ...
// ...

// Slow is a sleepy reader. Used for testing.
type Slow struct {
	Sleep time.Duration
}

// Read returns a few bytes after a given duration.
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
	io.Copy(os.Stdout, r)

	// Slow reader will timeout, if Slow sleeps to long.
	r = &TimeoutReader{
		r:       &Slow{Sleep: 200 * time.Millisecond},
		timeout: 150 * time.Millisecond,
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	// Hello World!
	// 2017/01/19 19:56:29 timeout
	// exit status 1
}
