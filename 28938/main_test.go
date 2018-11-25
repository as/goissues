package main

import (
	"fmt"
	"testing"
	"time"
)

const (
	E         = 20
	Size      = 1 << E
	maxWriter = 2
)

func hammer(c Cache, done chan bool) {
	r := newXor(Size)
	for {
		select {
		case <-done:
			return
		default:
		}
		v := r.next()
		if v&3 == 0 {
			// happens with probability 1/8 (~13%)
			c.Put(v, v)
		}
	}
}
func getter(c Cache, done chan bool) {
	r := newXor(Size)
	for {
		select {
		case <-done:
			return
		default:
		}
		c.Get(r.next())
	}
}

func TestMap(t *testing.T) {
	t.Skip("usability test skipped")
	l := newLC(Size)
	s := newSC(Size)
	l.Put(4, 4)
	if l.Get(4) != 4 {
		t.Fatal("l 4!=4")
	}
	s.Put(4, 4)
	if s.Get(4) != 4 {
		t.Fatal("s 4!=4")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go hammer(l, done)
		go hammer(s, done)
	}
	defer close(done)
	for i := 0; i < 1024*10; i++ {
		l.Get(i * 65537)
		s.Get(i * 65537)
	}
}
func BenchmarkMap(b *testing.B) {
	const k = 4 // skip by 2^4

	for w := 0; w <= maxWriter; w++ {
		for nm, new := range map[string]func(int) Cache{
			"Lock": newLC,
			"Sync": newSC,
		} {
			// Iterate in series and randomly
			for _, ctrFunc := range []func(uint32) seq{
				newCtr, newXor,
			} {
				c := new(Size)
				done := make(chan bool)
				for i := 0; i < w; i++ {
					// Mutate the cache as dictated by the outermost loop
					go hammer(c, done)

					// Add 10 readers for every writer
					for i := 0; i < 10; i++ {
						go getter(c, done)
					}
				}

				// Start with a narrow access window, expand the mask
				// outward until the Get can stratify the entire map. Early
				// iterations will be localized and expand as 'e' grows
				for e := uint(0); e <= E; e += k {
					ctr := ctrFunc((1 << e) - 1)
					b.Run(fmt.Sprintf("%dW%dR/%s/%s", w, w*10+1, nm, ctr), func(b *testing.B) {
						for n := 0; n < b.N; n++ {
							c.Get(ctr.next())
						}
					})
				}

				close(done)
			}

		}
	}
}

type seq interface {
	next() int
}

type ctr struct {
	int  uint32
	mask uint32
}
type xorshift32 struct {
	v    uint32
	mask uint32
}

func newCtr(mask uint32) seq { return &ctr{^uint32(0), mask} }
func (c *ctr) next() int {
	c.int++
	return int(c.int & c.mask)
}

func newXor(mask uint32) seq { return &xorshift32{uint32(time.Now().UnixNano()), mask} } // prime
func (x *xorshift32) next() int {
	x.v ^= x.v << 13
	x.v ^= x.v >> 17
	x.v ^= x.v << 5
	return int(x.v & x.mask)
}

func (x xorshift32) String() string { return fmt.Sprintf("Rand/Mask%x", x.mask) }
func (c ctr) String() string        { return fmt.Sprintf("Incr/Mask%x", c.mask) }
