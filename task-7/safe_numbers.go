package main

import (
	"fmt"
	"log"
	"sync"
)

type SafeNumbers struct {
	mux     sync.RWMutex
	numbers map[int]int
}

func (s *SafeNumbers) Add(key, value int) {

	//lock to write so there is no race
	s.mux.Lock()
	defer s.mux.Unlock()

	//check value for uniqueness
	_, ok := s.numbers[key]
	if ok {
		log.Println("this number already exists")
		return
	}

	s.numbers[key] = value
	fmt.Printf("number written key %d value %d\n", key, value)
}

func main() {
	wg := &sync.WaitGroup{}
	writer := SafeNumbers{
		mux:     sync.RWMutex{},
		numbers: make(map[int]int),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 5; i++ {
			writer.Add(i, i*i)
		}
	}()

	wg.Wait()
}
