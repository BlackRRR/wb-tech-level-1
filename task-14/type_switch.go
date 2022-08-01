package main

import "fmt"

func CheckType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("%d = type int\n", i)
	case string:
		fmt.Printf("%s = type string\n", i)
	case bool:
		fmt.Printf("%v = type bool\n", i)
	case chan int:
		fmt.Printf("%v = type channel\n", i)
	}

}

func main() {
	CheckType("Hello world")
	CheckType(15)
	CheckType(false)
	CheckType(make(chan int))
}
