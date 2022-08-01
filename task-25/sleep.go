package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	for i := 0; i <= 10; i++ {
		sleep(time.Second * 1)
		fmt.Println(i)
	}
}
