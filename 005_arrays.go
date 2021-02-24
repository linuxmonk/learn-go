package main

// Go has only 3 data structures
// - Arrays
// - Slices
// - Maps

// Arrays are
// - a series of contiguous memory locations.
// - fixed size
//
// This allows faster memory operations because of CPU's
// prefetch module can pre-load memory into cache lines
// to improve performance.
//
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/arrays/README.md

import "fmt"

func main() {
	sep("Array Declaration: var")
	arrayDeclarationVar()
	sep("Array Declaration: short")
	arrayShortDeclaration()
	sep("Array of different sizes are not same type")
	arrayDifferentSizes()
	sep("Array stores data in contiguous memory locations")
	arrayShowAllocationIsContiguous()
	sep("Range Semantics and Mechanics")
	rangeSemantics()
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

func arrayDeclarationVar() {
	// declare and initialize array with zero values
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Watermelon"
	fruits[3] = "Pineapple"
	fruits[4] = "Mango"

	// fruits array string will be stored like this
	// in memory.
	//
	//  |------||------||------||------||------|
	//  | nil  || nil  || nil  || nil  || nil  |
	//  |------||------||------||------||------|
	//  |  0   ||  0   ||  0   || 0    || 0    |
	//  |------||------||------||------||------|
	//
	// String is immutable. Go keeps the value of "Apple"
	// in memory and the two word data structure refers to
	// this from other references. For -
	//
	// p = fruits[0]
	// Will lead to the representation below -
	//
	// for i, fruit := range fruits { ... }
	//
	// The variable 'fruit' which also points to 'Apple' will get
	// a 2 Word data structure like this -
	//
	//  p    || 0xabcd1234 | 5 || --->   |------||
	// fruit || 0xfeffa234 | 5 || --->   | 0x0a || -----> Apple
	//                                   |------||
	//                                   |  5   ||
	//                                   |------||
	//
	// When passing 'fruit' to fmt.Println, a 2 Word data structure is
	// passed to Println. Which again points to the above data in heap
	// from the stack below.
	fmt.Println("Fruits:")
	for i, fruit := range fruits {
		if i == 1 {
			fruits[1] = "Jackfruit"
		}
		fmt.Println(i, fruit)
	}
}

func arrayShortDeclaration() {
	fmt.Println("Fibs:")
	// Short declaration
	fibs := [4]int{1, 2, 3, 5}
	// iterate array using index
	for i := 0; i < len(fibs); i++ {
		fmt.Println(fibs[i])
	}
}

func arrayDifferentSizes() {
	// Arrays of different sizes are not the same type
	// even if they hold same data type
	var five [5]int
	four := [4]int{1, 2, 3, 4}
	// Throws compilation error -
	//  > cannot use four (type [4]int) as type [5]int in assignment
	// five = four
	fmt.Println("five: ", five)
	fmt.Println("four: ", four)
}

func arrayShowAllocationIsContiguous() {
	fmt.Println("Show how array is allocated on contiguous memory:")
	languages := [5]string{"C", "Go", "Python", "Bash", "Rust"}
	for i, lang := range languages {
		fmt.Printf("i = [%v], language: %10v, address: %p\n", i, lang, &languages[i])
	}
}

func rangeSemantics() {

	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	languages := [5]string{"C", "C++", "Go", "Python", "Bash"}
	// will have a backing array like this
	//
	//  |------||------||------||------||------|
	//  | 0xa  || 0xb  || 0xc  || 0xd  || 0xe  |
	//  |------||------||------||------||------|
	//  |  1   ||  3   ||  2   || 6    || 4    |
	//  |------||------||------||------||------|
	//     |       |       |      |       |
	//     V       V       V      V       V
	//     C      C++     Go    Python   Bash
	//
	// taking only one value from range gives the index
	fmt.Println("Range: index-only")
	for i := range languages {
		fmt.Println(i, languages[i])
	}

	// NOTE (**)  The code
	// `for i, lang := range languages { ... }`
	// The `range languages` will result in a new backing array in memory
	// like this -
	//
	//  |------||------||------||------||------|
	//  | 0x0  || 0x1  || 0x2  || 0x3  || 0x4  |
	//  |------||------||------||------||------|
	//  |  1   ||  3   ||  2   || 6    || 4    |
	//  |------||------||------||------||------|
	//
	// lang holds a copy (2-word data struct) from languages
	// reference counting to the actual string.

	fmt.Println("Range: index, value")
	fmt.Println("Before: (Value, Addres): ", languages[1], &languages[1])
	for i, lang := range languages {
		languages[1] = "Rust"
		if i == 1 {
			fmt.Printf("In loop: lang: %v, languages[1]: %v\n", lang, languages[1])
		}
	}
	fmt.Println("After: (Value, Addres): ", languages[1], &languages[1])

	// Pointer semantic on range. -- DON"T DO THIS --
	fmt.Printf("Range: %v for idx, ptr := range &array %v\n", colorGreen, colorReset)
	fmt.Printf("%v DON'T DO THIS %v\n", colorRed, colorReset)
	fmt.Println("Before (Ptr): (Value, Addres): ", languages[1], &languages[1])
	for i, lang := range &languages {
		languages[1] = "Erlang"
		if i == 1 {
			fmt.Printf("In loop (Ptr): lang: %v, languages[1]: %v\n", lang, languages[1])
		}
	}
	fmt.Println("After (Ptr): (Value, Addres): ", languages[1], &languages[1])
}
