# Part 3 — 17 — False sharing and cache lines

**Goal:** Recognize when independent variables slow each other down because they share a CPU cache line.

## Problem

Two goroutines updating different fields can still fight if those fields sit on the same cache line.

## Solution

Keep frequently written independent counters apart in highly contended low-level code. Prefer simpler designs unless profiling points here.

## Code snippet

```go
package main

type Counter struct {
	value int64
	_     [56]byte // illustrative padding for a 64-byte cache line
}

type Sharded struct {
	a Counter
	b Counter
}

func main() {
	_ = Sharded{}
}
```

## Conclusion

False sharing is an advanced performance topic. Do not start here; arrive here from profiles and benchmarks.

## What improves if you use this

Very hot concurrent counters can scale better across CPU cores.

## Final things to know

**Q: Is padding always good?**  
A: No. It increases memory use.

**Q: Where does false sharing appear?**  
A: Metrics counters, queues, ring buffers, and low-level runtimes.

**Q: How do I prove it?**  
A: Benchmark under realistic concurrency.
