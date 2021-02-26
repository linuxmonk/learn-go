package main

import (
	"fmt"
	"unicode/utf8"
)

// Slice is a reference type unlike Arrays / structs.
// They are like channels, maps, interfaces
// Zero value of a reference type is nil.

func main() {
	sep("Create Slice: Short and Var declarations")
	createSliceDeclaration()
	sep("Slice: Nil Slice, Empty Slice, Empty Struct")
	nilEmptySliceStruct()
	sep("Append Slice: Shows how capacity of backing array changes")
	appendSlice()
	sep("Append To Slice: The right way")
	appendSliceDont()
	sep("Slice of Slice")
	sliceOfSlice()
	sep("Slice side-effects")
	sliceSideEffects()
	sep("Slice copy-on-write")
	sliceCOWWithAppend()
	sep("Slice copy")
	sliceCopy()
	sep("Slices and References")
	slicesAndReferences()
	sep("Slices and Strings")
	slicesAndStrings()
	sep("Slices Range Mechanics: Value vs. Pointer")
	slicesRangeMechanics()
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

func createSliceDeclaration() {
	//
	// Slice is a 3-Word data structure internally.
	// languages := make([]string, 5)
	// Creates a slice with a length of 5 elements of
	// string type.
	// languages
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ Backing Array Like in 005_Arrays.go ]
	// |--------------|
	// | 5            | (Length)
	// |--------------|
	// | 5            | (Capacity)
	// |--------------|
	//
	languages := make([]string, 5)
	languages[0] = "C"
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Rust"
	languages[4] = "Bash"
	// languages[5] = "Something" --> results in a panic
	fmt.Printf("languages: %v, addr: %p, len: %v, cap: %v\n",
		languages, &languages, len(languages), cap(languages))

	var fruits []string // zero value slice: set to nil
	// NOTE: This is a nil slice
	// |--------------|
	// | nil          |
	// |--------------|
	// | 0            | (Length)
	// |--------------|
	// | 0            | (Capacity)
	// |--------------|
	fmt.Printf("fruits: %v, addr: %p, len: %v, cap: %v\n",
		fruits, &fruits, len(fruits), cap(fruits))

	shoppingList := make([]string, 5, 10) // 5 is the length, 10 is capacity
	fmt.Printf("shoppingList: %v, addr: %p, len: %v, cap: %v\n",
		shoppingList, &shoppingList, len(shoppingList), cap(shoppingList))

	fmt.Println("Calling inspectSlice: languages")
	// This will pass the 3-word data structure to inspectSlice
	//
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ Backing Array Like in 005_Arrays.go ]
	// |--------------|
	// | 5            | (Length)
	// |--------------|
	// | 5            | (Capacity)
	// |--------------|
	inspectSlice(languages)
}

func sliceOfSlice() {
	languages := make([]string, 5, 8)
	languages[0] = "C"
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Rust"
	languages[4] = "Bash"

	fmt.Println("Full Slice of Languages:")
	inspectSlice(languages)

	//
	// s := make([]int, 5, 8)
	// This slice 's' is represented as a 3-word data structure
	// in memory.
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ 1 | 2 | 3 | 4 | 5 | x | x | x ]
	// |--------------|
	// | 5            | (Length)
	// |--------------|
	// | 8            | (Capacity)
	// |--------------|
	//
	// A new slice based off of 's'
	// ns := s[1:4] will look like this
	// NOTE: The new capacity will be the total capacity from
	// the new starting point
	// |--------------|
	// | 0xf1b2cfd8e5 | --> [ 2 | 3 | 4 | x | x | x | x ]
	// |--------------|
	// | 3            | (Length)
	// |--------------|
	// | 7            | (Capacity)
	// |--------------|

	// slice of slice [a:b)
	// NOTE:
	// a = start index
	// b = end index (not including it)
	// The best way to avoid errors when passing slices to
	// functions would be to never
	// let 'b' exceed the length of the slice.
	fmt.Println("Slice of Slice languages[1:4)")
	inspectSlice(languages[1:4])
}

func sliceSideEffects() {

	languages := make([]string, 5, 8)
	languages[0] = "C"
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Rust"
	languages[4] = "Bash"

	favorites := languages[0:3] // C, Go, Python

	fmt.Println("***Initial State of Slices***")
	fmt.Println("languages")
	inspectSlice(languages)
	fmt.Println("favorites")
	inspectSlice(favorites)
	fmt.Println("*****************************")
	favorites[2] = "Rust"
	fmt.Println("**********************************")
	fmt.Println("      favorites[2] = Rust")
	fmt.Println("languages")
	inspectSlice(languages)
	fmt.Println("favorites")
	inspectSlice(favorites)
	fmt.Println("**********************************")

	// NOTE Append has similar side effects too
}

func sliceCOWWithAppend() {

	languages := make([]string, 5, 8)
	languages[0] = "C"
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Rust"
	languages[4] = "Bash"

	// indicates the new slice favorites
	// starts at 0, till index 2
	// slice has capacity of 3
	favorites := languages[0:2:2] // C, Go
	fmt.Println("***Initial State of Slices***")
	inspectSlice(languages)
	inspectSlice(favorites)
	fmt.Println("*****************************")

	favorites = append(favorites, "Tcl")
	fmt.Println("**********************************")
	inspectSlice(languages)
	inspectSlice(favorites)
	fmt.Println("***Change is seen only in favorites**")
}

func sliceCopy() {

	languages := make([]string, 5, 8)
	languages[0] = "C"
	languages[1] = "Go"
	languages[2] = "Python"
	languages[3] = "Rust"
	languages[4] = "Bash"

	lang2 := make([]string, 5, 10)
	copy(lang2, languages)

	inspectSlice(languages)
	inspectSlice(lang2)
}

type user struct {
	name  string
	likes int
}

func slicesAndReferences() {
	// NOTE maintaining a reference / pointer to a slice that will be modified
	// and might change it's size is a bad idea. This can lead to erroneous
	// behaviour and memory leaks.
	users := []user{
		{
			name:  "Tenali Ramanan",
			likes: 0,
		},
		{
			name:  "Motu Patlu",
			likes: 0,
		},
		{
			name:  "Chota Bheem",
			likes: 0,
		},
	}

	fmt.Println(">> Initial:")
	inspectUserSlice(users)
	motu_fan := &users[1]
	motu_fan.likes++
	fmt.Println(">> After (motu_fan likes):")
	inspectUserSlice(users)
	fmt.Println(">> Adding new TV series [Hanuman].")
	// This will change the backing array and address of the users
	users = append(users, user{name: "Hanuman", likes: 0})
	inspectUserSlice(users)
	fmt.Println(">> motu_fan gives his episode one more vote")

	// NOTE motu_fan causes a memory leak by keeping a reference to an older
	// backing array while the users moves on to use a new backing array
	// with extended capacity.
	motu_fan.likes++
	fmt.Printf(">> After (motu_fan likes again)")
	inspectUserSlice(users)
}

func slicesAndStrings() {

	var buf [utf8.UTFMax]byte

	s := "अ आ इ ई उ ऊ ऋ ॠ ऌ ॡ ए ऐ ओ औ अं "
	fmt.Println("Sanskrit Svaras / Vowels", s)

	// NOTE range syntax can be used on a string
	// to iterate.
	// i is the index position for every code point
	// r is the rune (typed int32) to hold the code point
	for i, r := range s {
		// Number of bytes in a rune (rune length)
		rl := utf8.RuneLen(r)
		// Determine the end index
		si := i + rl
		// NOTE
		// copy works on arrays and string. Since string is immutable
		// string has to always be the source.
		copy(buf[:], s[i:si])
		// print some information
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}

func slicesRangeMechanics() {

	languages := []string{"C", "Go", "Python", "Rust"}
	// Say at - 0xc0b1343216
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ Backing Array to strings ]
	// |--------------|
	// | 4            | (Length)
	// |--------------|
	// | 4            | (Capacity)
	// |--------------|
	//

	// for iterate over with VALUE SEMANTICS
	// NOTE this loop should have iterated only 2 times but
	// it iterates all the values. Because behind the scenes
	// "range languages" works on a copy of the slice thus
	// preventing side effects.
	// "_, v := range languages" internally has something like this
	//
	// Say at - 0xf0f1c1b216
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ Backing Array to strings ]
	// |--------------|
	// | 4            | (Length)
	// |--------------|
	// | 4            | (Capacity)
	// |--------------|
	//
	fmt.Println(">>> VALUE SEMANTICS <<<")
	inspectSlice(languages)
	for _, v := range languages {
		languages = languages[0:2]
		fmt.Println(">> languages = languages[0:2]")
		inspectSlice(languages)
		fmt.Println(v)
	}

	fmt.Println(">>> POINTER SEMANTICS <<<")
	languages = []string{"C", "Go", "Python", "Rust"}
	inspectSlice(languages)
	// for iterate over with POINTER SEMANTICS
	// NOTE this loop will iterate panic because
	//
	// As above initially languages is set to
	//
	// Say at - 0xc0b1443220
	// |--------------|
	// | 0xa1b2c3d4e5 | --> [ Backing Array to strings ]
	// |--------------|
	// | 4            | (Length)
	// |--------------|
	// | 4            | (Capacity)
	// |--------------|
	//
	// "i := range languages" behind the scenes uses
	// pointer semantics to point to the same languages
	// above at 0xc0b1443220. Cutting that short to 2
	// and then iterating over it and it panics!!
	for i := range languages {
		languages = languages[:2]
		fmt.Println(">> languages = languages[0:2]")
		inspectSlice(languages)
		fmt.Printf("%d: %s\n", i, languages[i])
	}
}

func nilEmptySliceStruct() {
	var nilSlice []string
	// NOTE: This is a nil slice
	// |--------------|
	// | nil          |
	// |--------------|
	// | 0            | (Length)
	// |--------------|
	// | 0            | (Capacity)
	// |--------------|
	fmt.Printf("nilSlice: %+v, len: %d, cap: %d\n",
		nilSlice, len(nilSlice), cap(nilSlice))

	var emptyStruct struct{}
	// NOTE emptyStruct is a Zero allocation type. No matter how many
	// empty structs are created they all point to a 8 byte value in a Go Runtime
	fmt.Printf("emptyStruct: %v\n", emptyStruct)

	emptySlice := []string{}
	// NOTE: This is an empty slice
	// |--------------|
	// | *emptyStruct | (EmptyStruct)
	// |--------------|
	// | 0            | (Length)
	// |--------------|
	// | 0            | (Capacity)
	// |--------------|
	//
	fmt.Printf("emptySlice: %+v, len: %d, cap: %d\n",
		emptySlice, len(emptySlice), cap(emptySlice))

}

func inspectSlice(langs []string) {
	// NOTE langs still points to the same backing array. So a modification here
	// can be seen in the caller.
	fmt.Printf("Address: [%p], Length: [%d], Capacity: [%d]\n", &langs, len(langs), cap(langs))
	for i, lang := range langs {
		fmt.Printf("[%d] %p %s\n",
			i,
			&langs[i],
			lang)
	}
}

func inspectUserSlice(users []user) {
	// NOTE langs still points to the same backing array. So a modification here
	// can be seen in the caller.
	fmt.Printf("Address: [%p], Length: [%d], Capacity: [%d]\n", &users, len(users), cap(users))
	for i, u := range users {
		fmt.Printf("[%d] %p [name: %s, likes: %d]\n",
			i,
			&users[i],
			u.name, u.likes)
	}
}

func appendSliceDont() {
	// This will cause appending the values, C, C++, Java
	// after the 5th element.
	langs := make([]string, 5, 5)
	langs = append(langs, "C")
	langs = append(langs, "C++")
	langs = append(langs, "Java")
	for i, l := range langs {
		fmt.Printf("i: %v, value: %v\n", i, l)
	}
}

func appendSlice() {
	// Program shows how the backing array capacity is increased as
	// we append 100,000 elements to the slice

	// When we don't know how much data would go into a slice, start
	// with a nil slice.
	var data []string

	// Setting the capacity to 1e5 will not result in periodic
	// allocation of data
	// Will set the capacity to 1e5
	//
	// data := make([]string, 0, 1e5)

	lastCap := cap(data)
	for i := 0; i < 1e5; i++ {
		rec := fmt.Sprintf("Rec: %d", i)
		data = append(data, rec)
		if cap(data) != lastCap {
			capChange := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				i,
				cap(data),
				capChange)
		}
	}
}
