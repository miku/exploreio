// S11: Read a file, line by line.
//
// Like `cat -n`.
//
// OUTPUT:
//
//     $ cat hello.txt | go run main.go
//         1  Don't communicate by sharing memory, share memory by communicating.
//         2  Concurrency is not parallelism.
//         3  Channels orchestrate; mutexes serialize.
//         4  The bigger the interface, the weaker the abstraction.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	var i int
	for {
		i++
		s, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("% 6d  %s", i, s)
	}
}
