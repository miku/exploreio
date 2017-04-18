Exploring Readers and Writers
=============================

A tour through IO in Golang.

This repository contains material from a workshop held at
[Golab](http://golab.io) 2017.

The goal of this workshop was to show two important interfaces of the standard
library in action: [io.Reader](https://golang.org/pkg/io/#Reader) and
[io.Writer](https://golang.org/pkg/io/#Writer). There are numerous
implementations of these inferfaces.

Where are these interfaces used? In short: everywhere. You can read from a
file, a network connection, an HTTP response body, compressed data, you can read
endlessly from an artificial source or a stream of zeros or random noise.
Image, JSON, XML and other decoders work with readers. You can write to
files, network connections, hashing algorithms or standard output. You can read
and write to memory buffers.

The use cases are broad and there is a chance, that a custom implementation of
a reader or writer in your own project will simplify overall program design by
adhering to these widely used interface.

Slides
------

The [slides](https://github.com/miku/exploreio/blob/master/Slides.md) contain some
introduction and context. Here are a few resources, that are mentioned in the slides:

* The talk, that inspired this workshop: [Go
Proverbs](https://youtu.be/PAAkCSZUG1c?t=5m18s) at Gopherfest 2015.
* A short statement from Ken Thompson on [what a file is in
Unix](https://youtu.be/tc4ROCJYbm0?t=12m55s), from an 1982 AT&T documentary.
* The mantra [everything is a
file](https://en.wikipedia.org/wiki/Everything_is_a_file) remains an amazing
concept; it has been reformulated as [everything is a stream of
bytes](http://yarchive.net/comp/linux/everything_is_file.html).
* *Pipes and Filters* are an important [architectural
pattern](https://john.cs.olemiss.edu/~hcc/csci581oo/notes/pipes.html).
* The whole [Go Tour](https://tour.golang.org) is amazing, among the many
examples, there is also a [rot13Reader](https://tour.golang.org/methods/23).
* The Go documentation is really nice, the [package
io](https://golang.org/pkg/io/) docs contain a number of examples as well.

Duration
--------

The duration of the workshop at [Golab](http://golab.io) 2017 was about three
hours in which we looked at about 70% of the material.

Prerequisites
-------------

To keep the amount of preparation at a minimum, all you'll need is a working
[Go installation](https://golang.org/doc/install). The example programs will be
mostly self-contained.

Here is all you need to do to get started. Clone this repository:

```
$ git clone https://github.com/miku/exploreio.git
$ cd exploreio
```

And run the following command:

```
$ go run check/main.go
Hello Gopher!
```

If this worked, that congratulations: you are all set.

From simple to complex: one quiz at a time
------------------------------------------

This workshop uses the format of a *quiz* to introduce various concepts. Each
directory (s00, s01, ...) contains one exercise or example. There is always one
executable `main.go`. There might be auxiliary files in the directory.

Inside `main.go` you find code and comments. Among the comment lines are two
special kinds of comments, marked

* *TODO* and
* *OUTPUT*.

Here is an real exercise to illustrate the format:

```go
// S06: Besides marshaling, JSON (and XML) can also be decoded.
//
// OUTPUT:
//
//     $ cat hello.json | go run main.go
//     It's around 2017-01-20 17:15:54.603712222 +0100 CET now and ...
//     we are at wonderful Golab 2017 in Firenze! @golab_conf
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type record struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func main() {
	// TODO: Unmarshal from standard input into a record struct (4 lines).
	// ...
	// ...
	// ...
	// ...
	fmt.Printf("It's around %s now and we are at wonderful %s in %s! @golab_conf\n",
		rec.Date, rec.Name, rec.Location)
}

```

The first line gives a one-line summary of the theme of this exercise:

```go
// S06: Besides marshaling, JSON (and XML) can also be decoded.
```


Below the line marked with *OUTPUT* you will find the expected output of the
command, along with the way to call the file. Sometimes you need to pass
parameters, or like here, you need to pipe something into the script:

```go
// OUTPUT:
//
//     $ cat hello.json | go run main.go
//     It's around 2017-01-20 17:15:54.603712222 +0100 CET now and ...
//     we are at wonderful Golab 2017 in Firenze! @golab_conf

```

The other section is marked *TODO*:

```go
	// TODO: Unmarshal from Stdin into a record struct (4 lines).
	// ...
	// ...
	// ...
	// ...
```

It contains a short instruction of what should be accomplished in that
exercise. In parentheses you will find a hint, as for how many lines are needed
(in mostly idiomatic Go) to implement the task. It is just a hint, there are
often various ways to get to the same result.

If not noted otherwise, all tasks should be implemented with basic error
handling, that is: if any method returns an error, you should check it.
Since we write small scripts, we can safely quit the program, if some error
condition occurs.

For example, if a file is not found we can safely quit the program altogether:

```go
file, err := os.Open("filename.txt")
if err != nil {
	log.Fatal(err)
}
```

Once you filled in the code for solving the task, you should call the program
as indicated in the *OUTPUT* section.

```shell
$ cat hello.json | go run main.go
```

If the result is the same you solved the exercise. If the output is off, try
reviewing your code, move to another task or lookup the
[solution](https://github.com/miku/exploreio/tree/master#solutions).

Solutions
---------

Solutions and comments to all exercises can be found in
[Solutions.md](https://github.com/miku/exploreio/blob/master/Solutions.md). The
explanations strive for conciseness. I hope the aim to be brief does not make
the comments too cryptic.

To the instructor
-----------------

If you are using part or all of this material for a course, maybe you like a
setup like this:

* Have your laptop open as you show the exercises on a screen. Explain what is required and let people work on the TODOs.
* Have a separate screen (tablet or other device) open with the solutions manual ([available as PDF](https://github.com/miku/exploreio/blob/master/Solutions.pdf)),
so you can jump in with a hint or solution.
* Most of the examples are self contained; it should be easy to skip exercises,
if they are to easy or there's too little time.

List of exercises and examples (x)
----------------------------------

Example marked with an (x) do not contain a TODO. The are here for
illustration. There are comments in the code and some more context in the
solutions.

* [S00](https://github.com/miku/exploreio/tree/master/s00): A file is a reader.
* [S01](https://github.com/miku/exploreio/tree/master/s01): Get bytes from a reader.
* [S02](https://github.com/miku/exploreio/tree/master/s02): Copying and standard streams.
* [S03](https://github.com/miku/exploreio/tree/master/s03): Stdin is a file, too.
* [S04](https://github.com/miku/exploreio/tree/master/s04): Decompression with gzip and a simple filter chain.
* [S05](https://github.com/miku/exploreio/tree/master/s05): An simple image converter.
* [S06](https://github.com/miku/exploreio/tree/master/s06): Besides marshaling, JSON (and XML) can also be decoded.
* [S07a](https://github.com/miku/exploreio/tree/master/s07a): Package io contains many useful readers.
* [S07b](https://github.com/miku/exploreio/tree/master/s07b): Read sections from a reader.
* [S08](https://github.com/miku/exploreio/tree/master/s08): Strings can be readers, io.ReadFull.
* [S09](https://github.com/miku/exploreio/tree/master/s09): Random reader and Base64 encoder.
* [S10](https://github.com/miku/exploreio/tree/master/s10): Prettify tabular data.
* [S11](https://github.com/miku/exploreio/tree/master/s11): Read a file, line by line. (x)
* [S12](https://github.com/miku/exploreio/tree/master/s12): Read from multiple readers in turn.
* [S13](https://github.com/miku/exploreio/tree/master/s13): Write to more than one writer at once.
* [S14](https://github.com/miku/exploreio/tree/master/s14): Read into variables.
* [S15](https://github.com/miku/exploreio/tree/master/s15): Hello Buffer.
* [S16](https://github.com/miku/exploreio/tree/master/s16): Reading the output of a shell command into a buffer.
* [S17](https://github.com/miku/exploreio/tree/master/s17): An urgent request. (x)
* [S18a](https://github.com/miku/exploreio/tree/master/s18a): Response bodies.
* [S18b](https://github.com/miku/exploreio/tree/master/s18b): HTTP on the TCP level.
* [S19](https://github.com/miku/exploreio/tree/master/s19): An atomic file implementation. (x)
* [S20](https://github.com/miku/exploreio/tree/master/s20): The Reader interface.
* [S21](https://github.com/miku/exploreio/tree/master/s21): A reader that converts all unicode letter mapped to their upper case.
* [S22](https://github.com/miku/exploreio/tree/master/s22): A writer that discards everything that is written to it.
* [S23](https://github.com/miku/exploreio/tree/master/s23): An uppercase writer.
* [S24a](https://github.com/miku/exploreio/tree/master/s24a): A counting reader.
* [S24b](https://github.com/miku/exploreio/tree/master/s24b): A simple language guesser. (x)
* [S25](https://github.com/miku/exploreio/tree/master/s25): Generate data. (x)
* [S26](https://github.com/miku/exploreio/tree/master/s26): A slow reader. (x)
* [S27a](https://github.com/miku/exploreio/tree/master/s27a): BlackBar censors given words in a stream. (x)
* [S27b](https://github.com/miku/exploreio/tree/master/s27b): A reader that times out.
* [S28](https://github.com/miku/exploreio/tree/master/s28): Round robin multireader. (x)
* [S29](https://github.com/miku/exploreio/tree/master/s29): Round robin multireader, that can handle broken readers. (x)
* [S30](https://github.com/miku/exploreio/tree/master/s30): A small buffer.

Feedback
--------

I prepared this workshop, because I was curious about these interfaces myself.
This is a kind of ongoing project as I plan to add more implementations of
readers and writers as I encounter them.

For any idea, correction or extension, please file an issue or send a pull
request or patch. You can also contact me at
[martin.czygan@gmail.com](mailto:martin.czygan@gmail.com).
