package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func closeByChan(ch chan int) {
	//create a channel to complete
	done := make(chan struct{})

	//start the goroutine that writes data to the channel and ends the program when the data arrives in the channel done
	go func() {
		for {
			select {
			case ch <- rand.Intn(1000):
				fmt.Println("Sending data ...")
			case <-done:
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// one second after starting, the program will end
	go func() {
		time.Sleep(time.Second * 1)
		done <- struct{}{}
	}()

	for data := range ch {
		fmt.Printf("Read data ... %d\n", data)
	}
}

func closeByCtx(ch chan int) {
	ctx, cancel := context.WithCancel(context.Background())

	//run a goroutine that writes data to the channel and terminates the program when the data arrives in the context
	// about completion
	go func(ctx context.Context) {
		for {
			select {
			case ch <- rand.Intn(1000):
				fmt.Println("Sending data ...")
			case <-ctx.Done():
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)

	// one second after starting, the program will end
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()

	for data := range ch {
		fmt.Printf("Read data ... %d\n", data)
	}
}

func main() {
	ch := make(chan int)
	closeByCtx(ch)
	closeByChan(ch)
}
