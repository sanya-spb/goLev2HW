/*
вариант когда только на единичку надо инкрементить
*/
package main

import (
	"fmt"
)

var (
	counter    uint64 = 0
	maxCounter uint64 = 1000
)

func main() {
	fmt.Println("counter = ", counter)
	ch := make(chan struct{})

	for i := 1; uint64(i) <= maxCounter; i++ {
		go func(out chan<- struct{}) {
			out <- struct{}{}
		}(ch)
	}

	for counter < maxCounter {
		<-ch
		counter++
	}

	fmt.Println("counter = ", counter)
}
