package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	var wg sync.WaitGroup
	var mux sync.Mutex

	nGoroutines := 4
	loopCount := 2
	wg.Add(nGoroutines)
	fmt.Println("Initial Counter Value: ", counter)
	for i := 0; i < nGoroutines; i++ {
		go func() {
			for i := 0; i < loopCount; i++ {
				mux.Lock()
				v := counter
				v++
				counter = v
				mux.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value %d (Must be %d)\n", counter, nGoroutines*loopCount)
}
