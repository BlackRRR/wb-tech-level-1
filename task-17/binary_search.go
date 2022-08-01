package main

import "fmt"

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	//take a number from the middle of the array
	mid := (high + low) / 2

	for low <= high {
		//value found
		if arr[mid] == x {
			return mid
		}

		mid = (high + low) / 2
		// if the number from the center is more to the right than the given one, then we narrow the border to the right
		if arr[mid] > x {
			high = mid - 1
		}
		// if the number from the center is more to the left than the given one, then we narrow the border to the left
		if arr[mid] < x {
			low = mid + 1
		}
	}
	//if there is no such number, return -1
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}
