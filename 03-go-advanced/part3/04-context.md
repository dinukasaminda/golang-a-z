# Part 3 — 04 — `context`

**Goal:** Carry cancellation, deadlines, and request-scoped values across API boundaries.

## Problem

When a client disconnects or a request times out, downstream work should stop. Without cancellation, your service keeps doing useless work.

## Solution

Pass `context.Context` as the first parameter. Use `context.WithTimeout`, `context.WithCancel`, and check `ctx.Done()`.

## Code snippet

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func query(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("query complete")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println(query(ctx))
}
```

## Conclusion

`context` is the standard cancellation contract for servers, clients, database calls, and goroutines tied to a request.

## What improves if you use this

You reduce wasted work, leaked goroutines, and slow shutdowns.

## Final things to know

**Q: Where does `context.Context` go in a function signature?**  
A: First parameter, usually named `ctx`.

**Q: Should context store optional function parameters?**  
A: No. Use it only for cancellation, deadlines, and request-scoped values.

**Q: Why call `cancel()`?**  
A: It releases timer resources and signals children early.
