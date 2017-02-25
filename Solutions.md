S00
---

Change the length of the byte slice, from:

```go
// TODO: Change a single character in this program so the complete file is read and printed.
...
b := make([]byte, 11)
...
```

e.g. to 31:

```go
...
b := make([]byte, 31)
...
```

Any larger number will do as well. The length of the byte slice is the space
that we allow the read method to fill. If this space is too small, we won't be
able to read the whole file.

S01
---

We can shorten these lines:

```go
	// TODO: We don't need to loop manually, there is a helper function for that.
	// TODO: Replace the next 10 lines with 5 that do the same.
	var contents []byte
	for {
		b := make([]byte, 8)
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		contents = append(contents, b...)
	}
	fmt.Println(string(contents))
```

by using [io.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll):

```go
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
```

While [io.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) is useful, it is
sometimes overused. Why is that? Often one wants to read something, process it
and then write it somewhere. Imagine a HTTP request body, that we want to read,
then preprocess and then maybe write to a file.
[io.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) would consume the *whole
data at once*. But for example in the case of a large file upload, there is
seldom a reason, why we would need to load the whole file into memory before
writing it to disk. There are other ways to solve this problem, which are both
more efficient and elegant.

However, [io.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) is in the
standard library and has perfectly fine use cases, too.

Noteworthy: EOF will not be reported by `ioutil.ReadAll` as the
purpose of the method is to consume the reader as a whole:

> ReadAll reads from r until an error or EOF and returns the data it read. A
successful call returns err == nil, not err == EOF. Because ReadAll is defined
to read from src until EOF, it does not treat an EOF from Read as an error to
be reported.

S02
---

Use: [io.Copy](https://golang.org/pkg/io/#Copy) and [os.Stdout](https://golang.org/pkg/os/#pkg-variables).

```go
	// TODO: Write output to Stdout, without using a byte slice (3 lines, including error handling).
	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
```

The importance of [io.Copy](https://golang.org/pkg/io/#Copy) can hardly be overstated:

> Copy copies from src to dst until either EOF is reached on src or an error
occurs. It returns the number of bytes copied and the first error encountered
while copying, if any.

Internally, [io.Copy](https://golang.org/pkg/io/#Copy) uses a [buffer](https://en.wikipedia.org/wiki/Data_buffer)
in an essential sense:

> In computer science, a data buffer (or just buffer) is a region of a physical
memory storage used to temporarily store data while it is being *moved from one
place to another*.

Everywhere, where readers and writers need to connect,
[io.Copy](https://golang.org/pkg/io/#Copy) can be used. As a first example,
here we read from a file and write to one of the [standard
streams](https://en.wikipedia.org/wiki/Standard_streams).

We will see the helpful [io.Copy](https://golang.org/pkg/io/#Copy) over and
over again.

S03
---

Use [os.Stdout](https://golang.org/pkg/os/#pkg-variables) and [os.Stdin](https://golang.org/pkg/os/#pkg-variables).

```go
	// TODO: Read input from standard input and pass it to standard output,
	// TODO: without using a byte slice (3 lines).
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
```

Here, we have the essence of a
[filter](https://en.wikipedia.org/wiki/Filter_(software)), namely a program,
that works with streams, but does not change the stream at all. One can be
reminded of the [identify
function](https://en.wikipedia.org/wiki/Identity_function).

S04
---

Use [gzip.Reader](https://golang.org/pkg/compress/gzip/#Reader).

>  A gzip.Reader is an io.Reader that can be read to retrieve uncompressed data from a gzip-format compressed file.

```go
	// TODO: Read gzip compressed input from standard input and print it to standard output,
	// TODO: without using a byte slice (7 lines).
	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
```

A filter, that decompresses data read from standard input. As soon we get to
[io.Copy](https://golang.org/pkg/io/#Copy), a decompressed stream has the same
*shape* as any other type that implements
[io.Reader](https://golang.org/pkg/io/#Reader).

S05
---

Go comes with an [image package](https://golang.org/pkg/image/) in the standard
library, which implements a basic 2-D image support.

> The fundamental interface is called [Image](https://golang.org/pkg/image/#Image).

There is a [Decode](https://golang.org/pkg/image/#Decode) method, that takes a
reader and turn it into an [Image](https://golang.org/pkg/image/#Image).

In turn, the concrete image subpackages implement an
[Encode](https://golang.org/pkg/image/jpeg/#Encode) method, which take an
[io.Writer](https://golang.org/pkg/io/#Writer) and an
[Image](https://golang.org/pkg/image/#Image) as an argument.

```go
	// TODO: Read the image, encode the image (5 lines).
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, nil)
```

This snippet takes an arbitrary reader (e.g. standard input) and turns it into
an image. The encoding methods are indifferent to the data sink, as long as
they implement [io.Writer](https://golang.org/pkg/io/#Writer).

S06
---

```go
	// TODO: Unmarshal from Stdin into a record struct.
	var rec record
	if err := json.NewDecoder(os.Stdin).Decode(&rec); err != nil {
		log.Fatal(err)
	}
```

S07a
----

```go
	// TODO: Only read the first 27 bytes of stdin. 3 (or 6) lines with error handling.
	if _, err := io.Copy(os.Stdout, io.LimitReader(os.Stdin, 27)); err != nil {
		log.Fatal(err)
	}
```

Also possible:

```go
	// TODO: Only read the first 27 bytes of stdin. 3 (or 6) lines with error handling.
	p := make([]byte, 27)
	_, err := os.Stdin.Read(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(p))
```

Also possible:

```go
	// TODO: Only read the first 27 bytes of stdin. 3 (or 6) lines with error handling.
	if _, err := io.CopyN(os.Stdout, os.Stdin, 27); err != nil {
		log.Fatal(err)
	}
```

S07b
----

```go
	// TODO: Print the string "io.Reader" to stdout. 4 lines.
	s := io.NewSectionReader(r, 5, 9)
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
```

S08
---

```go
	// TODO: Read the first 7 bytes of the string into buf, the print to stdout. 5 lines.
	buf := make([]byte, 7)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
```

S09
---

```go
	// TODO: Copy 24 byte from random source into the encoder. 3 lines including error handling.
	if _, err := io.CopyN(encoder, r, 24); err != nil {
		log.Fatal(err)
	}
```

S10
---

```go
	// TODO: Read tabulated data from standard in and write it to the tabwriter. 3 lines (incl. err).
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatal(err)
	}
```

S11
---

All done.

S12
---

```go
	// TODO: Read from four readers and write to stdout. 4 lines (incl. 1 long and err handling).
	rs := []io.Reader{
		strings.NewReader("Hello\n"),
		strings.NewReader("Gopher\n"),
		strings.NewReader("World\n"),
		strings.NewReader("!\n"),
	}
	r := io.MultiReader(rs...)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
```

S13
---

```go
	// TODO: Write to both, the file and stdout. 4 lines (incl. error handling).
	w := io.MultiWriter(file, os.Stdout)
	if _, err := fmt.Fprintf(w, "SPQR\n"); err != nil {
		log.Fatal(err)
	}
```

S14
---

```go
	// TODO: Read an int, a float and a string from stdin. 3 lines.
	if _, err := fmt.Fscan(os.Stdin, &i, &f, &s); err != nil {
		log.Fatal(err)
	}
```

S15
---

```go
	// TODO: Read one byte at a time from buffer and print the hex value on stdout. 10 lines (incl. error handling).
	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stdout, "% x\n", b)
	}
```

S16
---

```go
	// TODO: Stream output of command directly into the buffer.
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
```

S17
---

All done.

S18a
----

```go
	// TODO: Like curl, print to stdout. 4 (5) lines (with err handling).
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
```

S18b
----

```go
	// TODO: Send a GET request, read the reponse and print to stdout.
	if _, err := io.WriteString(conn, "GET / HTTP/1.0\r\n\r\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
```

S19
---

All done.

S20
---

```go
// TODO: Implement the Read interface, always return EOF. 3 lines.
func (r *Empty) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}
```

S21
---

```go
// TODO: Implement UpperReader, a reader that converts all Unicode letter mapped to their upper case. 11 lines.
type UpperReader struct {
	r io.Reader
}

func (r *UpperReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	copy(p, bytes.ToUpper(p))
	return len(p), nil
}
```

S22
---

```go
// TODO: Implement Discard, that throws away everything that is written. 4 lines.
type Discard struct{}

func (r *Discard) Write(p []byte) (n int, err error) {
	return len(p), nil
}
```

S23
---

```go
type UpperWriter struct {
	w io.Writer
}

func (w *UpperWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(bytes.ToUpper(p))
}
```

S24a
----

```go
// TODO: implement a reader that counts the total number of bytes read. 9 lines.
type CountingReader struct {
	r     io.Reader
	count uint64
}

func (r *CountingReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	atomic.AddUint64(&r.count, uint64(n))
	return
}

func (r *CountingReader) Count() uint64 {
	return atomic.LoadUint64(&r.count)
}
```

S24b
----

All done.

S25
---

All done.

S26
---

All done.

---

S27a
----

All done.

S27b
----

```go
// TODO: Implement a reader that times out after a certain a given timeout. 19 lines.
type readResult struct {
	b   []byte
	err error
}

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
		return len(p), res.err
	}
}
```

S28
---

All done.

S29
---

All done.

S30
---

All done.

S40
---

All done.

S41
---

All done.

S42
---

All done.

S43
---

All done.

S44
---

All done.
