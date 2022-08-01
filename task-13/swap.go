package main

import "fmt"

func swap(a, b int) {
	fmt.Println("before swap ", "a = ", a, "b = ", b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("after swap ", "a = ", a, "b = ", b)
}

func swapGo(a, b int) {
	fmt.Println("before Go swap ", "a = ", a, "b = ", b)

	a, b = b, a

	fmt.Println("after Go swap ", "a = ", a, "b = ", b)
}

func main() {
	swapGo(5, 6)
	swap(5, 6)
}
