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
	ch := make(chan uint64)

	// стресс тест..
	for i := 1; uint64(i) <= maxCounter; i++ {
		go func() {
			ch <- uint64(1)
		}()
	}

	for counter != maxCounter {
		if val, ok := <-ch; ok {
			counter += val
		}
	}

	// ждать уже нечего, все сообщения получены, показываем как есть..
	fmt.Println("counter = ", counter)
}
