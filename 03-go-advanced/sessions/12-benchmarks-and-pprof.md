# Part 3 — 12 — Benchmarks and `pprof`

**Goal:** Measure performance with benchmarks and find bottlenecks with profiles.

## Problem

Performance guesses are usually wrong. You need evidence before optimizing.

## Solution

Write `BenchmarkXxx` tests, run `go test -bench`, then collect CPU or memory profiles with `-cpuprofile` and `-memprofile`.

## Code snippet

```go
package textutil

import (
	"strings"
	"testing"
)

func JoinWords(words []string) string {
	var b strings.Builder
	for _, w := range words {
		b.WriteString(w)
	}
	return b.String()
}

func BenchmarkJoinWords(b *testing.B) {
	words := []string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		_ = JoinWords(words)
	}
}
```

Run:

```text
go test -bench=. -benchmem
go test -bench=. -cpuprofile=cpu.out
go tool pprof cpu.out
```

## Conclusion

Benchmark before and after every optimization.

## What improves if you use this

You spend time on real bottlenecks instead of cosmetic changes.

## Final things to know

**Q: What does `b.N` mean?**  
A: The testing package chooses it to get stable timing.

**Q: Why use `-benchmem`?**  
A: It shows allocations per operation.

**Q: What should be optimized first?**  
A: The hottest measured bottleneck that matters to users.
