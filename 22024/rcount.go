package main

import (
	"os"
	"log"
	"flag"
)

var(
	bs = flag.Int("bs", 1, "block size")
)
func init(){
	flag.Parse()
}

func main() {
	var(
		nr int
		n int
		err error
		b = make([]byte, *bs)
	)
	for {
		n, err = os.Stdin.Read(b)
		nr++
		log.Printf("read #%d: n=%d (err=%s)\n", nr, n, err)
		if err != nil{
			break
		}
	}
	log.Printf("done")
	
}
