// S45: A sticky error writer.
//
// Stolen from [bradfitz/http2](https://github.com/bradfitz/http2/blob/aa7658c0e9902e929a9ed0996ef949e59fc0f3ab/transport.go#L70).
// See also: https://blog.golang.org/errors-are-values
// Live hacking: https://youtu.be/yG-UaBJXZ80?t=33m50s
//
// OUTPUT:
//
//     $ go run main.go
//     HELLO Alice
//     Protocols often consist of various steps. Some are more chatty than others.
//     mtu:1500;favorites={color:green,dessert:tiramisu}
//
//     OK
//
// Sometimes output will fail, since we are using a FlakyWriter to simulate errors
// (e.g. with a network connection). Then the output will look something like this:
//
//     $ go run main.go
//     HELLO Alice
//
//     2017/03/19 16:57:08 protocol initialization failed: writing failed for some reason
//     exit status 1

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

// stickyErrWriter keeps an error around, so you can *occasionally* check if an error occured.
type stickyErrWriter struct {
	w   io.Writer
	err *error
}

func (sew stickyErrWriter) Write(p []byte) (n int, err error) {
	if *sew.err != nil {
		return 0, *sew.err
	}
	n, err = sew.w.Write(p)
	*sew.err = err
	return
}

// FlakyWriter returns an error with probability p.
type FlakyWriter struct {
	p float64
	w io.Writer
}

// Write will fail with a given probability.
func (w FlakyWriter) Write(p []byte) (n int, err error) {
	if rand.Float64() < w.p {
		return 0, fmt.Errorf("writing failed for some reason")
	}
	return w.w.Write(p)
}

// ProtocolWriter writes various things. Each write operations might return an error.
type ProtocolWriter struct {
	w   io.Writer
	err error
}

// WriteHello says hi.
func (p ProtocolWriter) WriteHello(name string) error {
	_, err := p.w.Write([]byte("HELLO " + name + "\n"))
	return err
}

// WritePreamble writes a chatty preamble.
func (p ProtocolWriter) WritePreamble() error {
	_, err := p.w.Write([]byte("Protocols often consist of various steps. Some are more chatty than others.\n"))
	return err
}

// WriteSettings writes some important protocol settings.
func (p ProtocolWriter) WriteSettings() error {
	_, err := p.w.Write([]byte("mtu:1500;favorites={color:green,dessert:tiramisu}\n"))
	return err
}

func main() {
	// Vary the random seed.
	rand.Seed(time.Now().UnixNano())

	// Temporary buffer for protocol messages.
	var buf bytes.Buffer

	// A new ProtocolWriter.
	pw := ProtocolWriter{}

	// Set the ProtocolWriter's writer to a stickyErrWriter, which uses a FlakyWriter
	// to simulate failing writes (sometimes). Finally use a buffer as sink, so we
	// can inspect the output later.
	pw.w = stickyErrWriter{w: &FlakyWriter{w: &buf, p: 0.2}, err: &pw.err}

	// Here we can see the benefit of using a stickyErrWriter. We can call three
	// functions on our struct, of which each returns an error. Normally, we would
	// have to check them individually, but the stickyErrWriter keep the error around
	// for us.
	pw.WriteHello("Alice")
	pw.WritePreamble()
	pw.WriteSettings()

	// Just show what we have written so far.
	fmt.Println(buf.String())

	// Finally check for any error and exit program if one occured.
	if pw.err != nil {
		log.Fatalf("protocol initialization failed: %v", pw.err)
	}

	// Signal successful initialization.
	fmt.Println("OK")
}
