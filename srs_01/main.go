package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const n = 1000

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	wg := sync.WaitGroup{}

	go spinner(100 * time.Millisecond)

	// этих мы и будем потом ждать..
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			// эмуляция занятости процесса
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}()
	}

	wg.Wait()
	fmt.Printf("\rDone\n")
}

// анимашка, на то что мы не висим..
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
