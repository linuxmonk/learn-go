package main

import (
	"fmt"
	"os"
)

type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

type pipe struct {
	name string
}

// NOTE The concrete file type implements the 'reader' interface
// by using value symantics
// NOTE - 2: It is important to have this API / behaviour
// accept []byte as argument with a certain size rather than
// read trying return the []byte. This leads to bad API
// design as -
// - The API now has to know how much data to allocate
// - The allocated slice would have to be moved to heap as it is passed
//   back up the stack.
func (f file) read(b []byte) (int, error) {
	s := "This git repo has sample Go code written while listening to Go Training from Ardan Labs"
	n := copy(b, s)
	return n, nil
}

// pipe concrete type also implements the reader interface
// via value semantics
func (p pipe) read(b []byte) (int, error) {
	s := `
{
  name: "bill",
  profession: "programmer",
  languages: ["C", "C++", "Go", "Python"]
}`
	n := copy(b, s)
	return n, nil
}

func retrieve(r reader) error {
	data := make([]byte, 64)
	r.read(data)
	fmt.Printf("Retrieved data from %T: %s\n", r, string(data))
	return nil
}

// ==================== Method Sets ======================
// MethodSet is the set of methods defined by an interface.
// A type can implement an interface using pointer semantic
// or value semantic.
//
// NOTE - not all values are addressable. For example - constants are not addressable.
//
// 1. If you have a *T you can call methods that have a receiver type of *T as well as methods that have a receiver type of T (the passage you quoted, Method Sets).
// 2. If you have a T and it is addressable you can call methods that have a receiver type of *T as well as methods that have a receiver type of T, because the method call t.Meth() will be equivalent to (&t).Meth() (Calls).
// 3. If you have a T and it isn't addressable, you can only call methods that have a receiver type of T, not *T.
// 4. If you have an interface I, and some or all of the methods in I's method set are provided by methods with a receiver of *T (with the remainder being provided by methods with a receiver of T), then *T satisfies the interface I, but T doesn't. That is because *T's method set includes T's, but not the other way around (back to the first point again).

type notifier interface {
	notify1()
	notify2()
}

type user struct {
	name  string
	email string
}

type vUser int

func (u *user) notify1() {
	fmt.Printf("User %s<%s> has been notified - 1\n", u.name, u.email)
}

func (u *user) notify2() {
	fmt.Printf("User %s<%s> has been notified - 2\n", u.name, u.email)
}

func (v vUser) notify1() {
	fmt.Printf("vUser ID %d has been notified - 1\n", v)
}

func (v *vUser) notify2() {
	if v != nil {
		fmt.Printf("vUser ID %d has been notified - 2\n", *v)
	} else {
		fmt.Println("nil reciever on notify2")
	}
}

func main() {
	sep("File implements reader interface")
	readerInterfaceTest()
	sep("Implement polymorphism using interfaces")
	testPolymorphism()
	sep("MethodSets: Rules for implementing interface by pointer and value semantics")
	methodSets()
}

func readerInterfaceTest() {
	var f file
	buf := make([]byte, 512)
	fmt.Println("Type 'file' implementing a reader interface using value semantics")
	if _, err := f.read(buf); err != nil {
		fmt.Println("Error reading from file.")
		os.Exit(1)
	}
}

func testPolymorphism() {
	var f file
	var p pipe
	retrieve(f)
	retrieve(p)
}

const vID vUser = 41315

func methodSets() {

	u1 := user{
		name:  "Bob",
		email: "bob@mail.com",
	}

	fmt.Println("user implements notifer using pointer semantics.")
	fmt.Println("Calling with value:")
	u1.notify1()
	u1.notify2()
	fmt.Println("Calling with pointer:")
	p1 := &u1
	p1.notify1()
	p1.notify2()
	fmt.Println()

	v1 := vUser(12345)
	p2 := &v1
	fmt.Println("vUser implements notifer one method pointer semantic and one method value semantic")
	fmt.Println("Calling with value:")
	v1.notify1()
	v1.notify2()
	fmt.Println("Calling with pointer semantic (because v1 is addressable too)")
	p2.notify1()
	p2.notify2()

	fmt.Println()
	fmt.Println("vID of type vUser is a constant. Cannot be used with pointer semantic")
	vID.notify1()
	fmt.Println("NOTE vID which is constant cannot invoke 'notify2'")
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
