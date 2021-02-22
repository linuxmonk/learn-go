package main

import (
	"fmt"
	"unsafe"
)

type V struct {
	ok  bool
	i16 int16
	f32 float32
}

// Though V's total bytes is 7 bytes the data is padded for alignment
// | 1-byte | 2-bytes | 4-bytes |
// 1-byte ok can start at an address
//  - Say this starts at 0
// 2-byte i16 can start only at addresses divisible by 2
//  - This starts at 2
// 4-byte f32 can start only at addresses divisible by 4
//  - This starts at 4
// Making it a total of 8 bytes

// If i16 was replaced with i32 (int32) then the padding would increase
// to 3 bytes. To avoid extra padding try to order the struct elements
// from largest type to smallest type in the struct. NOTE this is only
// important / needs to be done when size of struct is a problem. Else
// readability takes precedence.

// NOTE The struct itself too gets aligned to the size of the largest member's
// size. Again not important unless it matters

func main() {

	var (
		val  V
		val3 V
	)

	fmt.Printf("val = %+v, type = %T, size = %v\n", val, val, unsafe.Sizeof(val))

	// struct literal
	val2 := V{
		ok:  true,
		i16: 234,
		f32: 3.1415,
	}
	fmt.Printf("val2 = %+v, type = %T, size = %v\n", val2, val2, unsafe.Sizeof(val2))

	// anonymous struct
	var anonStruct struct {
		ok  bool
		i16 int16
		f32 float32
	}
	fmt.Printf("anonStruct: %+v\n", anonStruct)

	// literal anon struct
	anonStruct2 := struct {
		ok  bool
		i16 int16
		f32 float32
	}{
		ok:  true,
		i16: 32000,
		f32: 3.1415,
	}
	fmt.Printf("anonStruct2: %+v\n", anonStruct2)

	// Go does not allow implicit conversion
	// https://play.golang.org/p/mFjapVOvH3Z

	// but this can only be explicitly state it via
	// type conversion
	val3 = V(anonStruct)
	fmt.Printf("val3: %+v\n", val3)
	val3 = anonStruct2
	fmt.Printf("val3: %+v\n", val3)
}
