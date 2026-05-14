# Part 3 — 04 — `context`

**Goal:** Carry cancellation, deadlines, and request-scoped values across API boundaries.

## Problem

When a client disconnects or a request times out, downstream work should stop. Without cancellation, your service keeps doing useless work.

## Solution

Pass `context.Context` as the first parameter. Use `context.WithTimeout`, `context.WithCancel`, and check `ctx.Done()`.

## Context types

Every context starts from a root and is then *derived* into a new context that adds cancellation, a deadline, or a value. You never mutate a context — you wrap it.

- `context.Background()` — the empty root context. Use it at the top of `main`, in `init`, tests, and incoming requests where no parent context exists. It is never cancelled and has no deadline.
- `context.TODO()` — same behavior as `Background()`, but signals "I haven't decided which context belongs here yet." Use it as a placeholder so static analyzers (and reviewers) can spot unfinished plumbing.
- `context.WithCancel(parent)` — returns a child and a `cancel` function. Calling `cancel()` closes `ctx.Done()` so downstream goroutines can stop. Always `defer cancel()`.
- `context.WithTimeout(parent, d)` — child that cancels itself after duration `d`. Equivalent to `WithDeadline(parent, time.Now().Add(d))`.
- `context.WithDeadline(parent, t)` — child that cancels at an absolute time `t`. Useful when you already know the SLA cutoff.
- `context.WithValue(parent, key, val)` — child that carries a request-scoped value (request ID, auth user, trace span). Use sparingly and only for values that truly cross API boundaries.

Cancelling a parent cancels every child. That's how a single client disconnect tears down the whole request tree.

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

## Cancellation inside a `for` loop

Long-running loops (workers, pollers, streaming consumers) must check the context on every iteration so they exit promptly when the parent cancels. The idiomatic pattern is `select` with a `default` for non-blocking work, or a dedicated `<-ctx.Done()` case alongside whatever event the loop is waiting on.

```go
func worker(ctx context.Context, jobs <-chan int) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case j, ok := <-jobs:
			if !ok {
				return nil
			}
			process(j)
		}
	}
}
```

For a polling loop with a ticker, the same shape applies:

```go
func poll(ctx context.Context) error {
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			tick()
		}
	}
}
```

Two rules to internalize:

1. Check `ctx.Done()` in the *same* `select` that consumes work — never after a blocking call you can't interrupt.
2. Return `ctx.Err()` so callers can distinguish `context.Canceled` from `context.DeadlineExceeded`.

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
