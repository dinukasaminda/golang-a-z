# Part 3 — 08 — Pipelines

**Goal:** Build staged data processing with channels.

## Problem

Data often moves through steps: read, parse, validate, enrich, store. Mixing all steps in one function makes cancellation and testing harder.

## Solution

Represent each stage as a function that receives from one channel and sends to another.

## Code snippet

```go
package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	for v := range square(gen(1, 2, 3)) {
		fmt.Println(v)
	}
}
```

## Conclusion

Pipelines make data flow explicit. Add cancellation before using them for long-running production work.

## What improves if you use this

Each stage becomes easier to test, replace, and scale.

## Final things to know

**Q: What closes an output channel?**  
A: The stage that sends on it.

**Q: What if a later stage stops early?**  
A: Earlier stages can block unless you add cancellation.

**Q: Are pipelines always needed?**  
A: No. Use them when staged concurrency clarifies the problem.
