package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	var wg sync.WaitGroup

	nGoroutines := 4
	loopCount := 2
	wg.Add(nGoroutines)
	fmt.Println("Initial Counter Value: ", counter)
	for i := 0; i < nGoroutines; i++ {
		go func() {
			for i := 0; i < loopCount; i++ {
				v := counter
				fmt.Println(counter)
				v++
				counter = v
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value %d (Must be %d)\n", counter, nGoroutines*loopCount)
}
