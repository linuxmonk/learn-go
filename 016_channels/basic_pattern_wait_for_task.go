package main

import (
	"fmt"
	"time"
)

func main() {
	waitForTask()
}

func waitForTask() {

	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("Received signal: ", p)
	}()

	time.Sleep(time.Duration(500) * time.Millisecond)
	ch <- "paper"
	fmt.Println("Sent signal")

	time.Sleep(time.Second)
	fmt.Println("------------------------")
}
