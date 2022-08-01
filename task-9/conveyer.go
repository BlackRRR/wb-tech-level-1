package main

import (
	"fmt"
	"sync"
)

// square - write the square of numbers to the channel
func square(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
	close(out)
}

// read data from the channel and output to the console
func read(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int, len(arr))
	out := make(chan int, len(arr))

	go square(in, out)
	wg.Add(1)
	go read(out, &wg)

	//write numbers to the channel
	for _, val := range arr {
		in <- val
	}
	close(in)
	wg.Wait()
}
