package main

import (
	"fmt"
	"time"
)

func main() {

	const emps = 2
	ch := make(chan string)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("Employee %d: received : %s\n", emp, p)
			}
			fmt.Printf("Employee %d: received shutdown signal\n", emp)
		}(e)
	}

	workers := 10
	for i := 0; i < workers; i++ {
		ch <- "paper"
		fmt.Printf("Worker sent work item %d\n", i)
	}

	close(ch)
	time.Sleep(time.Second)
	fmt.Println("-----------------------------------")
}
