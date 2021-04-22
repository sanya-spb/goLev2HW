package NewMapRWMu

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type MapRWMu struct {
	mu sync.RWMutex
	m  map[float64]struct{}
}

func NewMapRWMu() *MapRWMu {
	return &MapRWMu{
		m: make(map[float64]struct{}),
	}
}

func (c *MapRWMu) Read(key float64) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.m[key]
	return ok
}

func (c *MapRWMu) Write(key float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = struct{}{}
}

func runRWMutexMap(writeRate float64) {
	testMap := NewMapRWMu()

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

func BenchmarkRunRWMutexMap_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.1)
	}
}

func BenchmarkRunRWMutexMap_05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.5)
	}
}

func BenchmarkRunRWMutexMap_09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.9)
	}
}
