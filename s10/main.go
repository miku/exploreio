// S10: Prettify tabular data with tabwriter.
//
//     $ cat hello.tsv | go run main.go
//              Name| Age|Address
//              Paul|  23|1115 W Franklin
//     Bessy the Cow|   5|Big Farm Way
//              Zeke|  45|W Main St
//
package main

import (
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	// https://golang.org/pkg/text/tabwriter/#Writer
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// TODO: Read tabulated data from standard in and write it to the tabwriter. 3 lines (incl. err).
	io.Copy(w, os.Stdin)
	// ...
	// ...
	// ...
	w.Flush()
}
