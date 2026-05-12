# Part 3 — 15 — Escape analysis

**Goal:** Understand when values move from stack to heap.

## Problem

Heap allocations are useful but not free. Too many allocations increase GC work and latency.

## Solution

Use compiler escape reports to see why values allocate. Focus on hot paths measured by benchmarks.

## Code snippet

```go
package main

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func main() {
	_ = NewUser("Ada")
}
```

Inspect:

```text
go build -gcflags="-m" .
```

## Conclusion

Returning a pointer is not automatically bad. It is bad only when allocation cost matters and measurement says so.

## What improves if you use this

You can reduce allocations in hot code without guessing.

## Final things to know

**Q: What means "escapes to heap"?**  
A: The value must live beyond the current stack frame or cannot be proven local.

**Q: Should I avoid all heap allocations?**  
A: No. Optimize only meaningful hot paths.

**Q: Does using pointers always reduce copying?**  
A: It may add allocations and indirection. Measure.
