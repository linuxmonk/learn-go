package main

import (
	"errors"
	"fmt"
)

// If a function would need to return more than one single
// type of an error then use error values as shown below.

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrPageMoved  = errors.New("Page Moved")
)

func main() {

	if err := someCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Error")
		case ErrPageMoved:
			fmt.Println("Page has moved.")
		default:
			fmt.Println("Unknown error.")
		}
	}
}

func someCall(b bool) error {
	if b == true {
		return ErrBadRequest
	}
	return ErrPageMoved
}
