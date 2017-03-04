// S05: An simple image converter.
//
// Read PNG from standard input, write JPG to standard output.
//
// Inspired by Donovan, Kernighan (2016), page 287.
// https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// OUTPUT:
//
//    $ cat gopherbw.png | go run main.go > gopherbw.jpg
//
// You should find a new file gopherbw.jpg in the directory.
package main

import (
	_ "image/gif" // register gif decoder
	_ "image/png" // register png decoder
	"io"
	"log"
	"os"
)

// toJPG converts a GIF or PNG image int JPEG.
func toJPG(r io.Reader, w io.Writer) error {
	// TODO: Decode the image and encode it to JPEG, write it
	//       to the given writer (5 lines).
	// Hint: Utilize methods taking io.Reader or io.Writer
	//       in https://golang.org/pkg/image/.
	// ...
	// ...
	// ...
	// ...
	// ...
}

func main() {
	if err := toJPG(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
