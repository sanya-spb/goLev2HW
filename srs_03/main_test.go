package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

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

func BenchmarkRunMutexMap_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.1)
	}
}

func BenchmarkRunMutexMap_05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.5)
	}
}

func BenchmarkRunMutexMap_09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runRWMutexMap(0.9)
	}
}

func runRWMutexMap(writeRate float64) {
	var (
		globalMap   = map[float64]struct{}{}
		globalMapMu = sync.RWMutex{}
	)
	rand.Seed(time.Now().UTC().UnixNano())
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i float64) {
			defer wg.Done()
			if rand.Float64() < writeRate {
				globalMapMu.Lock()
				globalMap[i] = struct{}{}
				globalMapMu.Unlock()
				return
			}
			globalMapMu.RLock()
			globalMapMu.RUnlock()
		}(rand.Float64())
	}

	wg.Wait()
}

func runMutexMap(writeRate float64) {
	var (
		globalMap   = map[float64]struct{}{}
		globalMapMu = sync.RWMutex{}
	)
	rand.Seed(time.Now().UTC().UnixNano())
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i float64) {
			defer wg.Done()
			globalMapMu.Lock()
			defer globalMapMu.Unlock()
			if rand.Float64() < writeRate {
				globalMap[i] = struct{}{}
				return
			}
		}(rand.Float64())
	}

	wg.Wait()
}
