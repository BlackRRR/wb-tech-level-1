package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup) {
	//Create an array of elements
	array := []int{2, 4, 6, 8, 10}

	//Iterate through the array and create goroutines for each operation of calculating
	//the square of a number in the array
	for i := range array {
		//call the waitgroup Add method which increments the counter at each iteration by one
		//Which blocks the goroutine in which the wait method is implemented
		wg.Add(1)
		go func(i int) {
			fmt.Println("square = ", array[i]*array[i])
			//call the waitgroup Done method which decrements the counter on each iteration by one
			//When the counter reaches zero, the goroutine in which the wait method is implemented is unlocked
			wg.Done()
		}(i)
	}
}

func main() {
	//create a waitgroup for synchronization
	wg := &sync.WaitGroup{}

	square(wg)

	//run the wait method which waits for the end of the counter and blocks the goroutine
	wg.Wait()
}
