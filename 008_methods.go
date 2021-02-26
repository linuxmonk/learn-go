package main

import "fmt"

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

type entry struct {
	name  string
	email string
}

// NOTE
// In the signature below -
// func (u user) notify() is a method of type user.
// - 'u' is the receiver.
// - this method uses a VALUE SEMANTIC. 'u' is 'value type'
func (e entry) notify() {
	fmt.Printf("Email sent to %s<%s>\n", e.name, e.email)
}

func (e *entry) changeEmail(newEmail string) {
	e.email = newEmail
}

func main() {
	sep("Basics: Pointer vs. Value Semantics. Read Design notes")
	basic()
	sep("Example of a wrong semantic for a structure and range")
	wrongSemantic()
	sep("Type aliases and behaviour carry forward")
	typeAliasAndBehaviours()
	sep("Assigning method to function pointer and invoking (value vs pointer semantic)")
	firstClassFunctionsValueSemantics()
}

func basic() {
	/*
	 * DESIGN NOTE
	 *
	 * When do we use pointer semantic vs. value semantic ?
	 *
	 * For builtin-types (string, int, float, bool) use VALUE SEMANTIC
	 *
	 * For reference types (slice, map, channel, interfaces) use POINTER SEMANTIC
	 *
	 * For struct type this has to be decided based on the type of data.
	 *  - if you are not sure then always better to use pointer semantic.
	 *
	 * Ask these questions to yourself -
	 * - If adding / modifying the state of the data changes
	 *   the value to a completely new value. e.g. adding 3
	 *   seconds to a time is a new time. Then use value
	 *   semantics. Else use pointer semantics.
	 *
	 * Standard library uses this logic to determine which semantic
	 * to use. The factory function for a type determines the semantic
	 * all the methods adhere to.
	 *
	 * NOTE Semantic Violation: Donot switch the semantics for a type
	 * unless you are passing the value to one of these methods -
	 * - Unmarshal
	 * - Decode
	 * Since these need the pointer to the data.
	 */

	e := entry{"Hanumantha", "hanumantha@email.com"}
	fmt.Printf("Initial setting: %+v\n", e)
	e.notify()
	chEmail := "hanumantha@gmail.com"
	fmt.Printf("Change email to: %+v\n", chEmail)
	e.changeEmail(chEmail)
	e.notify()

	// Methods are just syntactic sugar. Go internally calls the method like this -

	fmt.Println(">>> Invocation directly instead of via syntactic sugar")
	entry.notify(e)
	(*entry).changeEmail(&e, "hanumantha@ch.secure.mail.com")
	fmt.Printf("Change email to: %+v\n", chEmail)
}

func inspectEntrySlice(entries []entry) {
	// NOTE langs still points to the same backing array. So a modification here
	// can be seen in the caller.
	fmt.Printf("Address: [%p], Length: [%d], Capacity: [%d]\n", &entries, len(entries), cap(entries))
	for i, e := range entries {
		fmt.Printf("[%d] %p [name: %s, email: %s]\n",
			i,
			&entries[i],
			e.name, e.email)
	}
}

func wrongSemantic() {
	entries := []entry{
		{"Hanumantha", "hanumantha@gmail.com"},
		{"Rama", "rama@gmail.com"},
	}
	inspectEntrySlice(entries)
	fmt.Println(">>> This produces incorrect results:")
	chEmail := "anonymous@anon.com"
	for _, e := range entries {
		fmt.Printf("Changing email of %s to %s\n", e.name, chEmail)
		e.changeEmail(chEmail)
	}
	fmt.Println(">>> Changing emails done.")
	inspectEntrySlice(entries)
}

func typeAliasAndBehaviours() {

	// The "person" is another type based on "entry"
	type person entry

	doc := `
The behaviour or methods are not carried forward
to the type upon which the derived type is based on.

Example -

	type entry struct {
	  name string
	}

	func (e entry) notify() {
	  fmt.Println("Sent notification to", e.name)
	}

	type person entry
	var p person

CANNOT CALL - 

	p.notify()
`
	// The behaviour of entry such as "notify" or "changeEmail"
	// are not carried forward to the type person.

	var p person
	p.name = "Some Person"
	p.email = "someperson@gmail.com"
	// p.notify undefined (type person has no field or method notify)
	// p.notify()
	fmt.Println(doc)
}

func firstClassFunctionsValueSemantics() {

	bill := entry{"Bill", "bill@gmail.com"}
	// f1 is a function that points to the notify method
	// on bill's data.
	fmt.Println("***** PTR TO FN - VALUE SEMANTIC *****")
	fmt.Printf(">>> init: %+v\n", bill)
	f1 := bill.notify
	fmt.Printf(">>> calling f1()\n")
	f1()

	// f1
	// |-------|     |----->| code of notify|
	// | f1 (*)|-----|      |---------------|
	// |-------|     |----->| copy of bill  |
	//
	// so f1 will only work on copy of bill because
	// "notify" is a VALUE SEMANTIC..
	// so calling bill.changeEmail() will not be seen
	// when calling f1()
	//

	// NOTE here since `notify` is a method that uses
	// value semantic, f1 will use a copy of bill to
	// invoke notify.

	fmt.Printf(">>> bill.changeEmail()\n")
	bill.changeEmail("bill@mydomain.mail.com")
	fmt.Printf(">>> calling f1()\n")
	f1()

	fmt.Println("***** PTR TO FN - POINTER SEMANTIC *****")
	fmt.Printf(">>> init: %+v\n", bill)
	f2 := bill.changeEmail

	// f2
	// |-------|     |----->| code of changeEmail |
	// | f2 (*)|-----|      |---------------------|
	// |-------|     |----->| pointer to bill     |
	//
	// so f2 will work bill because "changeEmail" works on
	// POINTER SEMANTIC.
	// so calling f2() and then printing bill will
	// show the changes of f2.
	//

	f2("bill@ch.secure.mail")
	fmt.Printf(">>> after f2() changeEmail: %+v\n", bill)
}
