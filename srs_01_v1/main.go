/*
вариант от "чукчи" :)
раз уж сделал, отправлю на рецензию..
правильные варианты ниже.
*/
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	counter    uint64 = 0
	maxCounter uint64 = 1000
	maxWait           = 10
)

func main() {
	fmt.Println("counter = ", counter)
	for i := 1; uint64(i) <= maxCounter; i++ {
		go func() {
			// counter++ плохое было решение
			// взял атомарный счетчик
			atomic.AddUint64(&counter, 1)
		}()
	}
	// подождем, пока все досчитает, но не дольше чем maxWait
	for i := 1; i <= maxWait; i++ {
		if counter == maxCounter {
			break
		}
		time.Sleep(time.Second)
		fmt.Print(".")
	}

	fmt.Printf("\ncounter = %d\n", counter)
}
