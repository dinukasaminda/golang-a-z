# Part 3 — 21 — Lock-free programming

**Goal:** Understand lock-free techniques and when not to use them.

## Problem

In extreme hot paths, locks can become bottlenecks. But replacing locks with atomics can introduce subtle correctness bugs.

## Solution

Use lock-free programming only for narrow, measured problems. Build on atomic operations and simple invariants.

## Code snippet

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var flag atomic.Bool
	if flag.CompareAndSwap(false, true) {
		fmt.Println("won initialization race")
	}
}
```

## Conclusion

Lock-free code is a specialist tool. Most application code is clearer and fast enough with channels or mutexes.

## What improves if you use this

The hottest contention points can avoid kernel blocking and reduce latency spikes.

## Final things to know

**Q: Does lock-free mean wait-free?**  
A: No. Lock-free guarantees system-wide progress, not that every goroutine finishes quickly.

**Q: What is CAS?**  
A: Compare-and-swap: update only if the current value is what you expected.

**Q: What is the biggest risk?**  
A: Correctness bugs that tests rarely expose.
