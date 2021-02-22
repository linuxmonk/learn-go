package main

import (
	"fmt"
)

func main() {

	var b byte
	var i int
	var j = 20
	var f float32
	var f2 float64
	var ok bool
	var s string
	var s2 = "hello, world"
	var r rune // is typed to int32
	var rune_literal = '✓'

	fmt.Printf("b  = %v, type = %T\n", b, b)
	fmt.Printf("i  = %v, type = %T\n", i, i)
	fmt.Printf("j  = %v, type = %T\n", j, j)
	fmt.Printf("f  = %v, type = %T\n", f, f)
	fmt.Printf("f2 = %v, type = %T\n", f2, f2)
	fmt.Printf("ok = %v, type = %T\n", ok, ok)
	fmt.Printf("s = %v, type = %T\n", s, s)
	fmt.Printf("s2 = %v, type = %T\n", s2, s2)
	fmt.Printf("rune = %v, type = %T\n", r, r)
	fmt.Printf("rune_literal = %q, type = %T\n", rune_literal, rune_literal)
}
