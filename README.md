Explore Golang IO
=================

Explore IO with Golang, workshop at [Golab](http://golab.io) 2017.

Prerequisites
-------------

To keep the amount of preparation at a minimum, all you'll need is a working
[Go installation](https://golang.org/doc/install). Exercise programs will be
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

If this worked you are all set.

From simple to complex: one quiz at a time
------------------------------------------

The workshop uses the format of a *quiz* to introduce various concepts. Each
directory (s00, s01, ...) contains one exercise. There is one `main.go` file
for each exercise. There might be auxiliary files in the exercise directory.

Inside `main.go` you find code and comments. Among the comment lines are two
special kinds of comments, marked *TODO* and *OUTPUT*.

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
	// TODO: Unmarshal from Stdin into a record struct (4 lines).
	// ...
	// ...
	// ...
	// ...
	fmt.Printf("It's around %s now and we are at wonderful %s in %s! @golab_conf\n",
		rec.Date, rec.Name, rec.Location)
}

```

The first line gives a one-line summation of the theme of this exercise:

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

The other section is marked TODO:

```go
	// TODO: Unmarshal from Stdin into a record struct (4 lines).
	// ...
	// ...
	// ...
	// ...
```

It contains a short instruction of what should be accomplished in that
exercise. In parentheses you will find a hint, as for how many lines are needed
in idiomatic Go to implement the task. It is just a hint, there are often various ways to get to 
the same result. A variation in the implementation can be good occasion for a discussion.

If not noted otherwise, all tasks should be implemented with basic error
handling, that is: if any method used returns an error, you should check it.
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
reviewing your code.

Solutions
---------

Solutions to all exercises can be found in
[Solutions.md](https://github.com/miku/exploreio/blob/master/Solutions.md).
Each exercise has a solution and potentially a short explanation that
comes with it.

The explanations strive for conciseness. I hope the aim to be brief does not
make the comments too cryptic.

To the instructor
-----------------

If you are using this material for an own course, maybe you like a setup like this:

* Have your laptop open as you show the exercises on a screen. Explain what is required and let people work on the TODOs.
* Have a separate screen (tablet or other device) open with the solutions manual,
so you can jump in with a hint or solution.

List of examples (x) and exercises
----------------------------------

* S00: A file is a reader.
* S01: Get bytes from a reader.
* S02: Copying and standard streams.
* S03: Stdin is a file, too.
* S04: Decompression with gzip and a simple filter chain.
* S05: An simple image converter.
* S06: Besides marshaling, JSON (and XML) can also be decoded.
* S07a: Package io contains many useful readers.
* S07b: Read sections from a reader.
* S08: Strings can be readers, io.ReadFull.
* S09: Random reader and Base64 encoder.
* S10: Prettify tabular data.
* S11: Read a file, line by line.
* S12: Read from multiple readers in turn.
* S13: Write to more than one writer at once.
* S14: Read into variables.
* S15: Hello Buffer.
* S16: Reading the output of a shell command into a buffer.
* S17: An urgent request. (x)
* S18a: Response bodies.
* S18b: HTTP on the TCP level.
* S19: An atomic file implementation. (x)
* S20: The Reader interface.
* S21: A reader that converts all unicode letter mapped to their upper case.
* S22: A writer that discards everything that is written to it.
* S23: An uppercase writer.
* S24a: A counting reader.
* S24b: A simple language guesser. (x)
* S25: Generate data. (x)
* S26: A slow reader. (x)
* S27a: BlackBar censors given words in a stream. (x)
* S27b: A reader that times out.
* S28: Round robin multi-reader.
