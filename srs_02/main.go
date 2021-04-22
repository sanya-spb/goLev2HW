// переделаем пример из методички под условие задачи
package main

import (
	"fmt"
	"sync"
	"time"
)

const count = 1000

func main() {
	var (
		counter int
		mutex   sync.Mutex

		// Вспомогательная часть нашего кода
		ch = make(chan struct{}, count)
	)
	for i := 0; i < count; i += 1 {
		go func() {
			// Захват мьютекса
			mutex.Lock()

			// переделка тут!
			// используем таким образом чтоб как можно скорее освободить мютекс
			func() {
				// Освобождение мьютекса
				defer mutex.Unlock()
				counter += 1
			}()

			// Фиксация факта запуска горутины в канале
			ch <- struct{}{}
		}()
	}
	time.Sleep(2 * time.Second)
	close(ch)

	i := 0
	for range ch {
		i += 1
	}
	// Выводим показание счетчика
	fmt.Println(counter)
	// Выводим показания канала
	fmt.Println(i)
}
