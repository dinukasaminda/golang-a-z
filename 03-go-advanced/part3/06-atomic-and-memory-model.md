# Part 3 — 06 — `sync/atomic` and memory model

**Goal:** Use atomic operations for small shared values and understand visibility between goroutines.

## Problem

Reading and writing a shared integer from multiple goroutines is a race, even if the operation looks tiny.

## Solution

Use `sync/atomic` for simple counters, flags, and pointers when a mutex would be too heavy or awkward.

## Code snippet

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}
```

## Conclusion

Atomics are powerful but narrow. Prefer mutexes until profiling or design pressure proves atomics are the right tool.

## What improves if you use this

Hot counters and flags can be updated safely with less lock contention.

## Final things to know

**Q: Are atomics easier than mutexes?**  
A: Usually no. They are easier to misuse.

**Q: Can atomics protect complex invariants?**  
A: Not cleanly. Use a mutex for multi-field state.

**Q: What is the memory model about?**  
A: It defines when writes in one goroutine become visible to another.
