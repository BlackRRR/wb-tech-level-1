package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writer(n int, ch chan int) {
	//create a timer
	timer := time.NewTimer(time.Duration(n) * time.Second)
	//write data to the channel, after the time expires, close the channel and exit
	for {
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("time out")
			return
		default:
			ch <- rand.Intn(1000)
		}
	}

}

func reader(ch chan int) {
	counter := 1
	// read data from the channel and output to the console
	for data := range ch {
		fmt.Println("read data: ", data, " ", counter, " time(s)")
		counter++
	}
	fmt.Println("channel closed")
}

func main() {
	var n int
	fmt.Scan(&n)

	ch := make(chan int)

	go reader(ch)
	writer(n, ch)
}
