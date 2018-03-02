package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("N: %d\n", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}(ch)

	wg.Wait()
}
