# Part 3 — 16 — Allocations and struct layout

**Goal:** Reduce unnecessary allocations and understand field alignment.

## Problem

Small layout choices can waste memory when a struct is stored millions of times.

## Solution

Group fields to reduce padding, preallocate slices when sizes are known, and check allocations with benchmarks.

## Code snippet

```go
package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	A bool
	B int64
	C bool
}

type Better struct {
	B int64
	A bool
	C bool
}

func main() {
	fmt.Println(unsafe.Sizeof(Bad{}))
	fmt.Println(unsafe.Sizeof(Better{}))
}
```

## Conclusion

Struct layout matters most for large collections, caches, and hot memory paths.

## What improves if you use this

You reduce memory footprint and GC pressure.

## Final things to know

**Q: What is padding?**  
A: Extra bytes inserted for CPU alignment.

**Q: Should readability lose to layout?**  
A: Only in measured hot or large data structures.

**Q: How do I see allocation counts?**  
A: Use `go test -bench=. -benchmem`.
