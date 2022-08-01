package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Workers struct {
	amount int
}

func (w *Workers) Read(ch chan interface{}) {
	//Create goroutines by the number of workers that will read data from the channel
	for i := 0; i <= w.amount; i++ {
		go func(i int) {
			for {
				data := <-ch
				fmt.Printf("Worker %d read data from channel: Type %T, Value %v\n", i, data, data)
			}
		}(i)
	}
}

func Write(ch chan interface{}) {
	//Create a channel with system signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		//At the end of the work through Ctrl + C The program will display a message
		case s := <-sigs:
			fmt.Printf("Programm has been stopped by signal %d\n", s)
			return
		default:
			//Записываем данные в канал
			ch <- rand.Intn(1000)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var n int

	// select the number of workers at start
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(err)
	}

	work := &Workers{
		amount: n,
	}

	ch := make(chan interface{})
	work.Read(ch)
	Write(ch)
}
