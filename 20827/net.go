package main

// Compile locally

import (
	"net"
	"os"
)

var conn net.Conn

func main() {
	os.Stderr.Close()
	
	net.Dial("tcp", "localhost:4444")
	
	// Linux and Plan 9: Panic writes to the network endpoint
	// Doesn't work on Windows
	panic("writes this message to the network endpoint")
}
