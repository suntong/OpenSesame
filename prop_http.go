package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func getPublicIPv4() string {
	// Create an IPv4 connection to the server, following `curl -v4 ifconfig.me`
	conn, err := net.Dial("tcp4", "ifconfig.me:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send the request ("Host:" is needed besides GET)
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: ifconfig.me\r\n\r\n")
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)

	// Splitting the headers and body
	headersAndBody := strings.SplitN(string(buffer[:n]), "\r\n\r\n", 2)
	//headers := headersAndBody[0]
	body := headersAndBody[1]
	//fmt.Println("Body:", body)

	return body
}
