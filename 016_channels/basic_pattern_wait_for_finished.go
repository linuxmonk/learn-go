package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This is signalling without data. Better suited with WaitGroups. But this can also be done with channels

func main() {
	waitForFinished()
}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("sent signal from goroutine")
	}()

	// When signalling without data there is no data to receive. So the _ on first arg.
	// second flag is boolean.
	// wd is true if there is data
	// wd is false if there is no data
	_, wd := <-ch
	fmt.Println("received signal: ", wd)
	time.Sleep(time.Second)
	fmt.Println("-----------------------------------")
}
