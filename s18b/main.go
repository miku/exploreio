// S18b: HTTP on the TCP level.
//
// OUTPUT:
//
//     $ go run main.go
//     HTTP/1.0 302 Found
//     Cache-Control: private
//     Content-Type: text/html; charset=UTF-8
//     Location: http://www.google.it/?gfe_rd=cr&ei=YoCCWNObL8HCXvCrjtgH
//     Content-Length: 256
//     Date: Fri, 20 Jan 2017 21:25:54 GMT
//
//     <HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
//     <TITLE>302 Moved</TITLE></HEAD><BODY>
//     <H1>302 Moved</H1>
//     The document has moved
//     <A HREF="http://www.google.it/?gfe_rd=cr&amp;ei=YoCCWNObL8HCXvCrjtgH">here</A>.
//     </BODY></HTML>
//
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// TODO: Send a GET request, read the reponse and print it
	//       to standard output (6 lines).
	// ...
	// ...
	// ...
	// ...
	// ...
	// ...
}
