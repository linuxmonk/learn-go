package main

func main() {
	example2(true, false, true, 25)
}

//go:noinline
func example2(b1, b2, b3 bool, i int) {
	panic("want a stack trace")
}
