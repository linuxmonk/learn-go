package main

import (
	"errors"
	"fmt"
)

var ErrPermission = errors.New("Permission Error")
var ErrBackendDown = errors.New("Backend inaccessible")

type AppError struct {
	code int
	err  error
}

func (a *AppError) Error() string {
	return fmt.Sprintf("code: %d, err: %v\n", a.code, a.err)
}

func callerPermission() {
	if err := errCall(1); err != nil && errors.Is(err, ErrPermission) {
		fmt.Println("Call failed due to permission error")
		return
	}
	fmt.Println("Call OK")
}

func callAppError() {
	var aerr *AppError

	if err := errCall(3); err != nil && errors.As(err, &aerr) {
		fmt.Println("Error of type AppError found")
		fmt.Printf("Reason: %+v\n", aerr)
		return
	}
	fmt.Println("Call OK")
}

func errCall(i int) error {
	switch i {
	case 1:
		return fmt.Errorf("Network error: %w", ErrPermission)
	case 2:
		return errors.New("invalid argument error")
	case 3, 4, 5, 6:
		return &AppError{
			code: i,
			err:  fmt.Errorf("Application Error: %w", ErrBackendDown),
		}
	default:
		return nil
	}
}

func main() {
	sep(`Test Wrapping with %w and errors.Is to compare if error value matches`)
	callerPermission()
	sep(`Test Wrapping with %w and errors.As to compare if error type matches`)
	callAppError()
}

func sep(heading string) {
	var heading_len int
	heading_len = len(heading)
	if len(heading) == 0 {
		heading_len = 25
	}

	for i := 0; i < heading_len; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Println(heading)
	for i := 0; i < heading_len; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
