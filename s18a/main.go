// S18a: Response bodies.
//
//     $ go run main.go
//     <!DOCTYPE html>
//     ...
//     <head>
//     <title>GoLab - The Italian conference on Go</title>
//     <meta name="description" content="GoLab - The Italian conference on Go">
//     <meta name="author" content="themecube">

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.golab.io/")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Like curl, print to stdout. 4 (5) lines (with err handling).
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
