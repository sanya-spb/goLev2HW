/*
вариант с универсальным мультипликатором
*/
package main

import (
	"fmt"
)

var (
	counter    uint64 = 0
	maxCounter uint64 = 1000
	multiplier uint64 = 2
)

func main() {
	fmt.Println("counter = ", counter)
	ch := make(chan uint64)

	for i := 1; uint64(i) <= maxCounter; i++ {
		go func(out chan<- uint64) {
			out <- uint64(multiplier)
		}(ch)
	}

	for counter < maxCounter*multiplier {
		if val, ok := <-ch; ok {
			counter += val
		}
	}

	fmt.Println("counter = ", counter/multiplier)
}
