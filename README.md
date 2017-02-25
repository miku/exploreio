Explore Golang IO
=================

Explore IO with Golang, workshop at [Golab](http://golab.io) 2017.

Prerequisites
-------------

To keep the amount of preparation you need to a minimum, all you'll need is a
working [Go installation](https://golang.org/doc/install). All programs in the
exercises will be mostly self-contained (in the spirit of a
[SSCCE](http://sscce.org/)).

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
	// TODO: Unmarshal from Stdin into a record struct. 4 lines with error handling.
	// ...
	// ...
	// ...
	// ...
	fmt.Printf("It's around %s now and we are at wonderful %s in %s! @golab_conf\n",
		rec.Date, rec.Name, rec.Location)
}

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