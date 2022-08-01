package main

import "fmt"

func intersection(a []int, b []int, counter map[int]int) []int {
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] = 1
		}
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count == 1 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 2, 2}
	b := []int{4, 5, 6, 8, 1, 10, 11, 0, 1, 2, 2}
	counter := make(map[int]int)

	fmt.Printf("%v\n", intersection(a, b, counter))
}
