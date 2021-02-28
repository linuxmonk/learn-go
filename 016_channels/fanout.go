package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Multiple goroutines writing to one buffered channel
	fanout()
}

func fanout() {

	var emps = 20
	ch := make(chan string, emps)

	// Signalling on the send side never blocks.
	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Printf("Employee %d sent paper\n", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Printf("Manager received %s\n", p)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("----------------------------------------")
}
