// S09: Random reader and Base64 encoder.
//
//
package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"os"
)

func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	r := rand.New(rand.NewSource(123))

	// TODO: Copy 24 byte from random source into the encoder. 3 lines including error handling.
	// ...
	// ...
	// ...

	// Note: The encoder has to be closed to write any partial data.
	if err := encoder.Close(); err != nil {
		log.Fatal(err)
	}
}
