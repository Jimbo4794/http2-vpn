package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

const url = "http://localhost:1010"

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get(url)
	checkErr(err, "during get")

	fmt.Printf("Client Proto: %d\n", resp.ProtoMajor)
}
