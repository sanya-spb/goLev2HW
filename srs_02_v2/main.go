/*
вариант с задержкой
doSomething(ctx) не запускал в отдельный поток, думаю это не принципиально.
*/
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		cancel()
		select {
		case <-time.After(time.Second * 1):
			log.Fatal("server finished unexpectedly")
		}
	}()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log.Printf("doSomething() started")

exit:
	for {
		select {
		case <-ctx.Done():
			log.Printf("server has gracefully finished, %v", ctx.Err())
			break exit
		default:
			for i := 0; i < 10; i++ {
				time.Sleep(200 * time.Millisecond)
				if i%10 == 0 {
					fmt.Print("|")
				} else if i%10 > 5 {
					fmt.Print("_")
				} else {
					fmt.Print(".")
				}
			}

		}
	}

	log.Printf("doSomething() stopped")
}
