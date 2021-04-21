package main

import (
	"math/rand"
	"sync"
	"time"
)

const n = 1000

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}()
	}

	wg.Wait()
}
