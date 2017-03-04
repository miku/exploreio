// S10: Prettify tabular data with tabwriter.
//
// OUTPUT:
//
//     $ cat hello.tsv | go run main.go
//              Name| Age|Address
//              Paul|  23|1115 W Franklin
//     Bessy the Cow|   5|Big Farm Way
//              Zeke|  45|W Main St
//
package main

import (
	"os"
	"text/tabwriter"
)

func main() {
	// https://golang.org/pkg/text/tabwriter/#NewWriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// TODO: Read tabulated data from standard in
	//       and write it to the tabwriter (3 lines).
	// ...
	// ...
	// ...
	w.Flush()
}
