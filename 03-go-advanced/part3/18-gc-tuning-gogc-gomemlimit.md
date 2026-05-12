# Part 3 — 18 — GC tuning with `GOGC` and `GOMEMLIMIT`

**Goal:** Understand Go garbage collection knobs and when to tune them.

## Problem

Allocation-heavy services can spend too much CPU in GC or exceed memory limits in containers.

## Solution

Measure allocation rate and heap behavior. Tune `GOGC` for collection frequency and `GOMEMLIMIT` for soft memory limits.

## Code snippet

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("heap alloc", m.HeapAlloc)
	fmt.Println("num gc", m.NumGC)
}
```

Run examples:

```text
GOGC=50 go run .
GOMEMLIMIT=512MiB go run .
```

## Conclusion

Tune GC only with memory and latency data. Most improvements come first from reducing unnecessary allocations.

## What improves if you use this

You can control the trade-off between memory usage, CPU spent on GC, and latency.

## Final things to know

**Q: What does lower `GOGC` do?**  
A: Collects more often, usually using less memory and more CPU.

**Q: Is `GOMEMLIMIT` a hard limit?**  
A: It is a soft runtime limit, not an OS memory guarantee.

**Q: What should I inspect first?**  
A: Heap profiles and allocation benchmarks.
