## Scheduler

### OS - pre-emptive scheduler

* Thread states: 
  - Running
  - Runnable
  - Waiting

Understand your workload. CPU vs IO bound work.

*NOTE* for CPU bound work more goroutines / threads means its just lot more context switching. More goroutines are good for IO bound tasks.

### GO scheduler

```
 /__M__\  OS Thread
    | <-------------------------- Go program's goroutine
 =======
|   P   | Logical Processor
 =======
```

* 1 P per core is allotted to a program.
* The Go scheduler is a co-operative scheduler which runs in userspace. It schedules go routines on to P's.
  - Go scheduler uses function calls / certain statements to determine the points of context switches between goroutines


## Cache Coherency and False Sharing

* *Cache Coherence* In a multi-core system, when each core has its own copy of the cache line (from RAM) then this could lead to cache coherence problem.
* *False Sharing* This happens when suppose 4 goroutines are modifiying their own piece of data in an array/slice then the array entries being modified might be in the cache line. Each processor updating anything in the line (its own entry) would still cause cache to be invalidated and repeated memory operations to keep things in sync. This is called Memory Thrashing due to False Sharing.

## Data Race

* `go build -race` can be used to build a program with race detection on. And it panics at runtime if there is a race condition.
* `go test -race` runs tests and tries to detect race condition
* Maps in Go are not thread safe. The versions of go after 1.11 have builtin race detection for maps. Which detect race conditions and panics.
