package main

import "fmt"

const i = 10      // constants of a kind
const pi = 3.1415 // constants of a kind

const j int = 40 // constants of a type (int)

// constants of a kind can be implicitly converted.
// NOTE as per spec constants of a kind can have a
// precision to 256 bits. Constants of a type have
// precision of 64 bits only.

const pi2 = 22.0 / 7.0

const zero = 1 / 3 // kind(1) / kind(3) => 0

// Looking at source of time constants
// https://golang.org/pkg/time/
// https://golang.org/src/time/time.go

type Duration int64

const (
	Nanosecond  Duration = 1                  // Typed constant with value 1
	Microsecond          = 1000 * Nanosecond  // kind(1000) * typed(1) ==> kind(1000)
	Millisecond          = 1000 * Microsecond // .. ditto ..
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func main() {

	fmt.Printf("i:    %v, type %T\n", i, i)
	fmt.Printf("j:    %v, type %T\n", j, j)
	fmt.Printf("pi:   %v, type %T\n", pi, pi)
	fmt.Printf("pi2:  %v, type %T\n", pi2, pi2)
	fmt.Printf("zero: %v, type %T\n", zero, zero)
	fmt.Println()

	fmt.Printf("Hour:        %20v, type %T\n", Hour, Hour)
	fmt.Printf("Minute:      %20v, type %T\n", Minute, Minute)
	fmt.Printf("Second:      %20v, type %T\n", Second, Second)
	fmt.Printf("Millisecond: %20v, type %T\n", Millisecond, Millisecond)
	fmt.Printf("Microsecond: %20v, type %T\n", Microsecond, Microsecond)
	fmt.Printf("Nanosecond:  %20v, type %T\n", Nanosecond, Nanosecond)

	fmt.Println("iota tests")
	iotaTest()
}

func iotaTest() {
	const (
		A1 = iota // starts at 0
		A2        // increments to 1
		A3        // increments to 2
	)

	const (
		A4 = iota + 10 // starts at 10
		A5             // increments to 11
		A6             // increments to 12
	)

	const (
		One = 1 << iota
		Two
		Four
		Eight
		Sixteen
	)

	fmt.Println("A1, A2, A3", A1, A2, A3)
	fmt.Println("A4, A5, A6", A4, A5, A6)
	fmt.Println("One, Two, Four, Eight, Sixteen", One, Two, Four, Eight, Sixteen)
}
