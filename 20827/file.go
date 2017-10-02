package main

// Compile locally

import (
	"net"
	"os"
)

var conn net.Conn

func main() {
	os.Stderr.Close()
	
	os.Create("test.txt")
	
	// Writes to the file "test.txt", works on all platforms tested (plan9, linux, windows)
	panic("writes this message to the network endpoint")
}
