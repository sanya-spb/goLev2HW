package main

import (
	"math/rand"
	"sync"
	"time"
)

type MapMu struct {
	mu sync.Mutex
	m  map[float64]struct{}
}

func NewMapMu() *MapMu {
	return &MapMu{
		m: make(map[float64]struct{}),
	}
}

func (c *MapMu) Read(key float64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.m[key]
	return ok
}

func (c *MapMu) Write(key float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = struct{}{}
}

func RunMutexMap(writeRate float64) {
	testMap := NewMapMu()

	rand.Seed(time.Now().UTC().UnixNano())
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i float64) {
			defer wg.Done()

			if rand.Float64() < writeRate {
				testMap.Write(i)
				return
			}
			_ = testMap.Read(i)
		}(rand.Float64())
	}

	wg.Wait()
}

func main() {
	RunMutexMap(0.5)
}
