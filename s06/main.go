// S06: Besides Marshaling, JSON (and XML) can also be decoded.
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
	var rec record
	if err := json.NewDecoder(os.Stdin).Decode(&rec); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("It's around %s now and we are at wonderful %s in %s! @golab_conf\n",
		rec.Date, rec.Name, rec.Location)
}
