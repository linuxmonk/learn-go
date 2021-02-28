package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func init() {
	// setting the value to 1 runs everything in order
	runtime.GOMAXPROCS(2)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Start Goroutines:")

	go func() {
		lowercase()
		// NOTE commenting the line below causes deadlock in `wg.Wait()`
		// go runtime raises a panic
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Println("Finished scheduling goroutines. Waiting...")
	// commenting this line below quits before goroutines run
	wg.Wait()
	fmt.Println("Terminating program.")
}

func lowercase() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: ", i)
		for r := 'A'; r <= 'Z'; r++ {
			c := strings.ToLower(string(r))
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func uppercase() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: ", i)
		for r := 'a'; r <= 'z'; r++ {
			c := strings.ToUpper(string(r))
			fmt.Print(c)
		}
		fmt.Println()
	}
}
