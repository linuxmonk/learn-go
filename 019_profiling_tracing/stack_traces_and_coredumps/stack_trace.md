## Stack Traces and Core Dumps

### How to Read a Stack Trace

```
func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) {
	panic("Want a stack trace")
}
```


A slice is a 3 word data structure 

```
 +---------+
 |    *    | -> pointer to backing array
 |    2    | -> len
 |    4    | -> cap
 +---------+
```

A string is a 2 word data structure 

```
 +---------+
 |    *    | -> pointer to backing array
 |    5    | -> len
 +---------+
```

An int is one word data structure

*Stack Trace dumped. Example-1*

```
panic: want a stack trace

goroutine 1 [running]:
main.example(0xc000050738, 0x2, 0x4, 0x1071f87, 0x5, 0xa)
        /Users/saikiran/src/learn-go/019_profiling_tracing/stack_traces_and_coredumps/stack_trace.go:9 +0x39
main.main()
        /Users/saikiran/src/learn-go/019_profiling_tracing/stack_traces_and_coredumps/stack_trace.go:4 +0x72
exit status 2
```

In the above output -

`main.example(0xc000050738, 0x2, 0x4, 0x1071f87, 0x5, 0xa)` represents the 5 words of data the `main()` sends to `example()`.

First 3 args are the 3 word string -

```
 +---------+
 |    *    | -> 0x000050738
 |  len    | -> 2
 |  cap    | -> 4
 +---------+
```

The next 2 arguments are 2 word string

```
 +---------+
 |    *    | -> 0x1071f87
 |  len    | -> 5
 +---------+
```

The next argument (`0xa`) is 10 in hex.

----

*Stack Trace dumped. Example-2*

```
func main() {
    example2(true, false, true, 25)
}

func example(b1, b2, b3 bool, i int) {
	panic("want a stack trace")
}
```

For this the stack trace looks like -

```
panic: want a stack trace

goroutine 1 [running]:
main.example2(0x1010001, 0x19)
        /Users/saikiran/src/learn-go/019_profiling_tracing/stack_traces_and_coredumps/stack_trace2.go:9 +0x39
main.main()
        /Users/saikiran/src/learn-go/019_profiling_tracing/stack_traces_and_coredumps/stack_trace2.go:4 +0x36
exit status 2
```

`main.example(0x1010001, 0x19)` represents the 4 arguments passed. The arguments are compacted as a boolean value
is represented using a byte.

```
0x1010001, 0x19

Reading 0x1010001 right-to left for Little Endian architecture -

01 -> true
00 -> false
01 -> true
1  -> unused/junk

0x19 -> hex value for 25
```
