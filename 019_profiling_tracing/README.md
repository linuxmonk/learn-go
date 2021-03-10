# Profiling and Tracing

## How a profiler works ?

A profiler runs your program and configures the OS to interrupt the process at regular intervals. This is done by sending SIGPROF to the program being profiled which suspends and transfers execution to the profiler. The profiler then grabs the program counter for each executing thread and restarts the program.

## Types of Profiling

- CPU profiling
- Memory profiling
- Block profiling

## Some of the issues

- Latency (Networking, Disk, etc.)
- Internal Latency (Synchronization)
- Allocation (on heap)
- Accessing data (like bringing things to cache lines etc.)
- Algorithm efficiencies

## Tools

- 'hey' to generate load (HTTP load generator) - https://github.com/rakyll
- 'ngrok' for load testing and capture traffic
- 'time' unix/linux utility to display how much time a process spend on user, system activities
- 'perf' on linux

## Basic Go Profiling

- Stack Traces and Core Dumps 
  - Read about [stack traces here](./stack_traces_and_coredumps/stack_trace.md)
  - Read about [core dumps here](./stack_traces_and_coredumps/core_dump.md)
- Read about [Benchmarking for micro level optimizations](./benchmarks/benchmarks.md)
- GODEBUG
- Memory and CPU profiling
- pprof profiling (http/pprof)
- Blocking profiling
- Mutex profiling
- Tracing


