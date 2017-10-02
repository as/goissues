package main

import (
	"os"
	"io"
)

func main() {
	io.Copy(os.Stdout, os.Stdin)
}
