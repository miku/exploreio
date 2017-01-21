S00
---

Change the length of the byte slice.

```go
b := make([]byte, 11)
```

e.g. to 41:

```go
b := make([]byte, 41)
```

S01
---

Use io.ReadAll:

```go
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
```

S02
---

* io.Copy, os.Stdout (os.Stderr, os.Stdin)

```go
	// TODO: Write output to Stdout, without using a byte slice (3 lines, including error handling).
	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
```

S03
---

```go
	// TODO: Read input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (3 lines, including error handling).
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
```

S04
---

Use a gzip reader.

```go
	// TODO: Read compressed input from stdin and pass it to Stdout.
	// TODO: without using a byte slice (7 lines, including error handling).
	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
```

S05
---

```go
	// TODO: Read the image, encode the image. 5 lines with error handling.
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 10})
```

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
