package main

import (
	"fmt"
	"sync"
)

var scores = make(map[string]int)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("Scores: A = %d, B = %d\n", scores["A"], scores["B"])
}
