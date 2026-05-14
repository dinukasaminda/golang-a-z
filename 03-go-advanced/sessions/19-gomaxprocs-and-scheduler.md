# Part 3 — 19 — `GOMAXPROCS` and scheduler

**Goal:** Understand how Go schedules goroutines onto OS threads.

## Problem

Concurrency does not mean unlimited parallel CPU execution. CPU-bound work depends on available processors and scheduler behavior.

## Solution

Use `runtime.GOMAXPROCS` to understand or set the number of logical processors used for running Go code.

## Code snippet

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gomaxprocs", runtime.GOMAXPROCS(0))
}
```

## Conclusion

The scheduler multiplexes goroutines over OS threads. For most programs, keep the default and focus on blocking, contention, and cancellation.

## What improves if you use this

You can reason about CPU parallelism and diagnose scheduler-related performance issues.

## Final things to know

**Q: Does increasing `GOMAXPROCS` always help?**  
A: No. It can increase contention.

**Q: What is the default?**  
A: Usually the available CPU quota/core count.

**Q: What blocks a goroutine?**  
A: Channel operations, locks, syscalls, timers, network I/O, or explicit waits.
