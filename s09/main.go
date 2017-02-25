// S09: Random reader and Base64 encoder.
//
// OUTPUT:
//
//     $ go run main.go
//     8UBc7YuZaLr5EJJZ
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

func main() {
	// base64.URLEncoding is safe for URLs and filenames
	encoder := base64.NewEncoder(base64.URLEncoding, os.Stdout)
	defer encoder.Close() // The encoder has to be closed to write any partial data.
	r := rand.New(rand.NewSource(123))

	// TODO: Copy 12 byte from random source into the encoder (3 lines).
	// ...
	// ...
	// ...
	fmt.Println()
}
