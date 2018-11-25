package main

import "sync"

type Cache interface {
	Get(key int) int
	Put(key int, value int)
}
type SC struct {
	sync.Map
}
type LC struct {
	sync.RWMutex
	m map[int]int
}

func newSC(size int) Cache {
	sc := &SC{}
	for i := 0; i < size; i++ {
		sc.Put(i, i)
	}
	return sc
}
func newLC(size int) Cache {
	c := &LC{m: map[int]int{}}
	for i := 0; i < size; i++ {
		c.Put(i, i)
	}
	return c
}

func (m *LC) Get(key int) int {
	m.RLock()
	val := m.m[key]
	m.RUnlock()
	return val
}
func (m *LC) Put(key int, value int) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}
func (m *SC) Get(key int) int {
	val, ok := m.Load(key)
	if !ok {
		return 0
	}
	return val.(int)
}
func (m *SC) Put(key int, value int) {
	m.Store(key, value)
}

func main() {}
