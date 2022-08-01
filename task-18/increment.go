package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	counter int
}

func (c *counter) Inc() {
	//lock the mutex to avoid races
	c.Lock()
	defer c.Unlock()
	c.counter++
}

func main() {
	wg := &sync.WaitGroup{}

	c := &counter{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.counter)
}
