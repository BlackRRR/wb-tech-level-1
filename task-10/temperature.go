package main

import "fmt"

func temperature(arr []float64, sets map[float64][]float64) {
	var i float64
	//go through from -100 to 100 in increments of 10 and write numbers to the set
	for i = -100; i <= 100; i += 10 {
		for j := range arr {
			if i < 0 {
				if arr[j] <= i && arr[j] > i-10 {
					sets[i] = append(sets[i], arr[j])
				}
			} else {
				if arr[j] >= i && arr[j] < i+10 {
					sets[i] = append(sets[i], arr[j])
				}
			}
		}
	}
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.1}
	sets := make(map[float64][]float64)

	temperature(arr, sets)

	for key, val := range sets {
		fmt.Println("key:", key, "val:", val)
	}
}
