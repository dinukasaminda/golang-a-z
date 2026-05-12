# Part 3 — 13 — Race detector

**Goal:** Detect unsynchronized shared memory access.

## Problem

Data races can pass tests for weeks and fail under production timing.

## Solution

Run tests or programs with `-race`. Fix races with ownership, channels, mutexes, or atomics.

## Code snippet

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	n := 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			n++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
```

Run:

```text
go test -race ./...
go run -race .
```

## Conclusion

Make `-race` part of your regular test habit for concurrent code.

## What improves if you use this

You catch dangerous concurrency bugs before users do.

## Final things to know

**Q: Does `-race` prove no races exist?**  
A: No. It detects races in executed code paths.

**Q: Is it slower?**  
A: Yes, so use it for testing, not normal production builds.

**Q: What is a race?**  
A: Concurrent access to the same memory where at least one access writes and there is no synchronization.
