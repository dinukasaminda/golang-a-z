# Part 3 — 05 — `sync` primitives

**Goal:** Use `Mutex`, `RWMutex`, `WaitGroup`, `Once`, and related primitives correctly.

## Problem

Multiple goroutines sometimes need shared state. Without coordination, you get data races and broken invariants.

## Solution

Use `sync.Mutex` to protect critical sections, `WaitGroup` to wait, `Once` for one-time initialization, and `RWMutex` when reads dominate writes.

## Code snippet

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func main() {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.n)
}
```

## Conclusion

Locks are simple when the protected data and the lock live together.

## What improves if you use this

Shared state stays correct under concurrency.

## Final things to know

**Q: Should I copy a struct containing a mutex?**  
A: No. Copying a used mutex is a bug.

**Q: Is `RWMutex` always faster?**  
A: No. Use it when reads heavily outnumber writes and contention matters.

**Q: What does `sync.Once` solve?**  
A: Safe one-time initialization across goroutines.
