package main

import (
	"fmt"
	"image/gif"
	"image/png"
	"io"
	"os"
)

type R struct {
	io.Reader
	n int
}

func (r *R) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.n += n
	return n, err
}

func main() {
	fmt.Println("attempting to read gif")

	file, err := os.Open("problem.gif")
	if err != nil {
		fmt.Println("Can't open this file:", err)
		return
	}

	r := &R{Reader: file}
	img, err := gif.Decode(r)
	if err != nil {
		fmt.Println("Can't decode this as a gif:", err, r.n, img)
		return
	}

	_ = img.Bounds()

	fmt.Println("gif decoded")
	png.Encode(os.Stdout, img)
}
