package main

import (
	"fmt"
	"log"
)

func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 3

	if n > len(arr) {
		log.Println("value to remove does not exist in array")
		return
	}

	fmt.Println(arr)
	fmt.Println(remove(arr, 1))
}
