package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {

	var wg sync.WaitGroup

	nGoroutines := 4
	loopCount := 2
	wg.Add(nGoroutines)
	fmt.Println("Initial Counter Value: ", counter)
	for i := 0; i < nGoroutines; i++ {
		go func() {
			for i := 0; i < loopCount; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value %d (Must be %d)\n", counter, nGoroutines*loopCount)
}
