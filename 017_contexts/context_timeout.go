package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(250 * time.Millisecond)
		ch <- "some work"
	}()

	select {
	case s := <-ch:
		fmt.Println("work complete: ", s)
	case <-ctx.Done():
		fmt.Println("work timed out")
	}
}
