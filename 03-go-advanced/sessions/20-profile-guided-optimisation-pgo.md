# Part 3 — 20 — Profile-guided optimisation

**Goal:** Use production-like CPU profiles to help the compiler optimize hot code.

## Problem

The compiler cannot always know which paths are hot in real workloads.

## Solution

Collect a representative CPU profile and build with PGO using `-pgo`.

## Code snippet

```text
go test -bench=. -cpuprofile=default.pgo
go build -pgo=default.pgo ./cmd/server
```

For automatic use, place `default.pgo` in the main package directory when appropriate.

## Conclusion

PGO is useful after you already have benchmarks and profiles. It is not a replacement for good algorithms.

## What improves if you use this

Hot paths may get better inlining and layout decisions from the compiler.

## Final things to know

**Q: What profile should I use?**  
A: A representative CPU profile from real or realistic traffic.

**Q: Can stale profiles hurt?**  
A: Yes. Refresh profiles when workloads change.

**Q: Should libraries ship PGO profiles?**  
A: Usually applications benefit more because they know real workloads.
