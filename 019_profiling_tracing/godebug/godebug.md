## GODEBUG

`GODEBUG` environment variable can be set before runnin a go binary to generate certain information about CPU / memory usage to perform any macro level optimizations.

### Scheduler Tracing

Setting `GODEBUG=schedtrace=1000 ./run/your/binary` will display a scheduler trace every 1000 milliseconds (1sec). Here if we do a load test we can see the output of how many goroutines are in idle state, how many are in spinning state, and number of logical processors and the the state of the local run queue.

### Memory Tracing

Setting `GODEBUG=gctrace=1 ./run/your/binary` will display information on when a GC runs during your program's cycle. It displays operations such has size of heap, it's growth and if it was brought down in that cycle.
