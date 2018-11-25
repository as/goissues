package main

import (
	"fmt"
	"math/bits"
	"math/rand"
	"testing"
	"time"
)

const (
	E         = 20
	Size      = 1 << E
	maxWriter = 1 // 2 will exceed 10 minutes on my system
)

func hammer(c Cache, done chan bool) {
	r := newRng(Size)
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
	r := newRng(Size)
	for {
		select {
		case <-done:
			return
		default:
		}
		c.Get(r.next())
	}
}

func TestRandtab(t *testing.T) {
	t.Skip("usability test skipped")
	var hit = [Size]int{}
	n := 0
	for _, v := range Randtab {
		hit[v]++
		n++
	}
	ssq := 0.0
	for _, v := range hit {
		p := float64(v) / float64(n)
		ssq += p * p
	}
	t.Log(ssq) //
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
			"Sync": newSC,
			"Lock": newLC,
		} {
			// Iterate in series and randomly
			for _, ctrFunc := range []func(uint32) seq{
				newCtr, newRng, // newMem,
			} {
				c := new(Size)
				done := make(chan bool)
				for i := 0; i < w; i++ {
					// Mutate the cache as dictated by the outermost loop
					go hammer(c, done)

					// Add 8 readers for every writer
					for i := 0; i < 10; i++ {
						go getter(c, done)
					}
				}
				// Add 1 more reader
				go getter(c, done)

				// Start with a narrow access window, expand the mask
				// outward until the Get can stratify the entire map. Early
				// iterations will be localized and expand as 'e' grows
				for e := uint(0); e <= E; e += k {
					ctr := ctrFunc((1 << e))
					b.Run(fmt.Sprintf("%dW%dR/%s/%s", w, w*8+2, nm, ctr), func(b *testing.B) {
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

func init() {
	x := newRng(Size)
	for i := 0; i < len(Randtab); i++ {
		Randtab[i] = x.next()
	}
}

var Randtab = [Size]int{}

type seq interface {
	next() int
}

type mem struct {
	int  uint32
	mask uint32
}

func (c *mem) next() int {
	c.int++
	return Randtab[c.int&c.mask]
}

func cm(mask uint32) uint32 {
	n := bits.OnesCount32(mask + 1)
	if n > 1 {
		panic(fmt.Sprintf("mask %x not p-1 congruent", mask))
	}
	return mask
}

func newCtr(max uint32) seq { return &ctr{^uint32(0), cm(max - 1)} }
func newMem(max uint32) seq { return &mem{^uint32(0), cm(max - 1)} }
func newRng(max uint32) seq { return &rng{rand.New(rand.NewSource(time.Now().UnixNano())), max} }

type ctr struct {
	int  uint32
	mask uint32
}
type rng struct {
	*rand.Rand
	max uint32
}

func (c *ctr) next() int {
	c.int++
	return int(c.int & c.mask)
}
func (x *rng) next() int {
	return x.Rand.Intn(int(x.max))
}

func (m mem) String() string { return fmt.Sprintf("Memo/Mask%x", m.mask) }
func (r rng) String() string { return fmt.Sprintf("Rand/Mask%x", r.max-1) }
func (c ctr) String() string { return fmt.Sprintf("Incr/Mask%x", c.mask) }
