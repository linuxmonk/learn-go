package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dropPattern()
}

func dropPattern() {

	cap := 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("Received signal: ", p)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Second)
		}
	}()

	workers := 20
	for w := 0; w < workers; w++ {
		select {
		case ch <- "paper":
			fmt.Println("Sent signal from: ", w)
		default:
			fmt.Println("Dropped signal for: ", w)
		}
	}

	close(ch)
	time.Sleep(time.Second)
	fmt.Println("----------------------------------------------")
}
