package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	server := http2.Server{}

	l, err := net.Listen("tcp", "0.0.0.0:1010")
	checkErr(err, "while listening")

	fmt.Printf("Listening [0.0.0.0:1010]...\n")
	for {
		conn, err := l.Accept()
		checkErr(err, "during accept")

		server.ServeConn(conn, &http2.ServeConnOpts{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Printf("New Connection: %+v\n", r)
				fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
			}),
		})
	}
}

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}
