# Part 3 — 03 — `select` and timeouts

**Goal:** Wait on multiple channel operations and avoid waiting forever.

## Problem

A goroutine might wait on a result that never arrives. In services, that becomes stuck requests, leaked goroutines, and exhausted resources.

## Solution

Use `select` to wait on multiple channel cases. Add `time.After` or context cancellation for timeouts.

## Code snippet

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		result <- "done"
	}()

	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}
```

## Conclusion

Every blocking channel wait in production code should have a clear reason it cannot wait forever.

## What improves if you use this

Services fail faster, recover capacity, and avoid goroutine leaks.

## Final things to know

**Q: What happens if multiple cases are ready?**  
A: Go chooses one pseudo-randomly.

**Q: Should I use `time.After` in a hot loop?**  
A: Prefer reusable timers in very hot paths to reduce allocations.

**Q: What is `default` for?**  
A: Non-blocking send or receive.
