# Benchmarking for microlevel optimizations

Using an example ([from Ardan Lab's performance chapter](https://play.golang.org/p/Cm92cvurEnE)) of streaming data and io package of Go library to examine 2 algorithms that implement the same logic and compare their performance.

To run the benchmarks -

- This command generates a memory profile
`go test -run non -bench -benchmem -memprofile m.out` 
- Then launch the `pprof` to read the memory profile and look at the numbers / hotspots.
- This command shows the allocations
`go test pprof -alloc_space m.out
- This can be examined in conjunction with escape analysis to view why certain allocations occurred.
- This command passes extra flags to generate escape analysis information before running benchmarks. Useful to compare findings from benchmarks against code's memory allocations/escapes
`go test -gcflags "-m -m" -run none -benchtime 3s -benchmem -memprofile m.out -bench .`

- This command generates CPU profile -
`go test -run none -bench . -benchtime 3s -cpuprofile c.out`
- To analyse the CPU profile data run pprof -
`go test pprof c.out`

  Sub-commands like (top, list <func>) show details of profiling
