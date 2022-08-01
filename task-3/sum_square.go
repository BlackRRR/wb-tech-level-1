package main

import (
	"fmt"
)

func sumSquares(array []int) int {
	//Check if the array is empty, then return the sum = 0
	if len(array) < 1 {
		return 0
	}

	newChan := make(chan int)
	var sum int
	//Iterate through the array and create goroutines for each operation of calculating the square of a number in the array
	for i := range array {
		go func(i int) {
			newChan <- array[i] * array[i]

		}(i)

		value := <-newChan

		sum += value
	}

	close(newChan)

	return sum
}

func sumSquares1(nums []int) int {
	//Check if the array is empty, then return the sum = 0
	if len(nums) < 1 {
		return 0
	}

	var sum int

	//create a buffered channel with a buffer size equal to the length of the array
	ch := make(chan int, len(nums))

	go func() {
		//Iterate through the array and create goroutines for each operation of calculating the square of a number in the array
		for i := 0; i < len(nums); i++ {
			go func(num int) {
				ch <- num * num
			}(nums[i])
		}
	}()

	//We pass through the array and sum up the value that came from the channel
	for i := 0; i < len(nums); i++ {
		sum += <-ch
	}

	close(ch)

	return sum
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	fmt.Println(sumSquares(array))
	fmt.Println(sumSquares1(array))
}
