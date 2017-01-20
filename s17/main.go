// S17: An urgent request.
//
// $ cat gopherbw.png.gz | go run main.go | sort -nr | head -10
//  8543296|0
//  6353501|65535
//     1346|5140
//      881|21588
//      789|5397
//      751|14135
//      677|21331
//      607|34181
//      547|11822
//      506|45489

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	r, _ := gzip.NewReader(os.Stdin)
	img, _, _ := image.Decode(r)

	// Get dimensions.
	rectangle := img.Bounds()
	width, height := rectangle.Size().X, rectangle.Size().Y

	// Store distribution of "r" values in a map.
	rdist := make(map[uint32]int)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			rdist[r]++
		}
	}

	// Write tabulated values into buf.
	var buf bytes.Buffer
	for k, v := range rdist {
		fmt.Fprintf(&buf, "%d\t%d\n", v, k)
	}

	// Write to stdout.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	io.Copy(w, &buf)
	w.Flush()
}
