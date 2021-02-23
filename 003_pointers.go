package main

import "fmt"

// NOTE Goroutine minimum stack frame size is 2K.
// NOTE Any stack frames from where control has returned is invalid

const size int = 1024

type User struct {
	name  string
	email string
}

func main() {
	passByValue()
	fmt.Println("---")
	passByPointer()
	fmt.Println("---")
	escapeAnalysis()
	fmt.Println("---")
	stackGrowth()
}

func passByValue() {

	var count = 10
	fmt.Println("passByValue:")
	fmt.Printf("count: value [%v], address: [%v]\n", count, &count)
	increment(count)
	fmt.Printf("count: value [%v], address: [%v]\n", count, &count)
}

func increment(count int) {
	count += 1
}

func passByPointer() {

	var count = 10

	fmt.Println("passByPointer:")
	fmt.Printf("count: value [%v], address: [%v]\n", count, &count)
	increment2(&count)
	fmt.Printf("count: value [%v], address: [%v]\n", count, &count)
}

func increment2(incr *int) {
	// count stores the address of an integer
	*incr++
	fmt.Printf("incr:  value [%v], address: [%v], value points to: [%v]\n", incr, &incr, *incr)
}

func escapeAnalysis() {
	fmt.Println("Escape Analysis:")
	u := createUserV1()
	fmt.Printf("caller: u:  [%+v], address: [%p]\n", u, &u)
	u2 := createUserV2()
	fmt.Printf("caller: u2: [%+v], address: [%p]\n", *u2, u2)
}

//go: noinline
func createUserV1() User {
	u := User{
		name:  "Bill",
		email: "bill.s@mail.com",
	}
	fmt.Printf("createUserV1: u: %+v, address: [%p]\n", u, &u)
	return u
}

//go: noinline
func createUserV2() *User {
	u := User{
		name:  "Bill",
		email: "bill.s@mail.com",
	}
	fmt.Printf("createUserV2: u: %+v, address: [%p]\n", u, &u)
	// The compiler's static analysis module allocates 'u' on the
	// heap instead of the stack. - Escape Analysis.
	return &u
}

func stackGrowth() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {

	if c%100 == 0 {
		fmt.Println(c, s, *s)
	}

	c++
	if c == 1000 {
		return
	}
	stackCopy(s, c, a)
}
