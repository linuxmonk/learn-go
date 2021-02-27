package main

// NOTE If a function wants to return a error message
// (which isn't meant to be parsed but just checked and logged)
// then use `errors.New()` to return the error message. As shown
// below

import (
	"errors"
	"fmt"
	"os"
)

func someCall() error {
	return errors.New("IO Error")
}

func main() {

	if err := someCall(); err != nil {
		fmt.Println("Error invoking someCall due to", err)
		os.Exit(1)
	}
	fmt.Println("All good")
}
