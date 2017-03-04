// S40: Draining a body.
//
// Example from: https://github.com/golang/go/blob/2815045a50862276082048714337f95c46e98605/src/net/http/httputil/dump.go#L26
//
// OUTPUT:
//
//     $ go run main.go
//     Hello Gophers
//     Hello Gophers
package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
//
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func main() {
	r := ioutil.NopCloser(strings.NewReader("Hello Gophers\n"))
	r1, r2, err := drainBody(r)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, io.MultiReader(r1, r2)); err != nil {
		log.Fatal(err)
	}
}
