package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// getLocalIPv4 returns the LAN (192.168.x.x) address of the interface used
// for outbound traffic.
// The UDP "dial" sends no packets, it only resolves the route.
func getLocalIPv4() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err == nil {
		defer conn.Close()
		return conn.LocalAddr().(*net.UDPAddr).IP.String()
	}

	// no default route (offline LAN): fall back to the first non-loopback IPv4
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return "unknown"
}

// getPublicIPv4 returns the Internet-facing Public IPv4 address
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
