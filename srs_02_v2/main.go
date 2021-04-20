/*
вариант с задержкой
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log.Printf("doSomething() started")
	ctx2, exitFunc := context.WithCancel(context.Background())

	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				log.Println("server has gracefully finished")
				return
			default:
				for i := 0; i < 9; i++ {
					time.Sleep(200 * time.Millisecond)
					fmt.Print(".")
				}
				fmt.Print("|")
			}
		}
	}(ctx2)

	<-ctx.Done()
	exitFunc()
	time.Sleep(time.Second)
	log.Printf("doSomething() stopped")
}
