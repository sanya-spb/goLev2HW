/*
этот вариант наверное неправельный, т.к. тут мы ничего не ждем, а просто завершаем все при сигнале SIGTERM
вариант с задержкой попробую реализовать в другом варианте..
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
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Print(".")
		}
	}()

	log.Printf("doSomething() started")
	<-ctx.Done()
	log.Printf("doSomething() stopped")
}
