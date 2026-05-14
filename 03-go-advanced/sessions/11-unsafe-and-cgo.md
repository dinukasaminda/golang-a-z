# Part 3 — 11 — `unsafe` and cgo

**Goal:** Understand when Go's safety boundaries can be bypassed and why that should be rare.

## Problem

Sometimes Go code must interact with memory layouts, syscalls, or C libraries that normal Go types cannot express directly.

## Solution

Use `unsafe` or cgo only at narrow boundaries. Hide it behind small, tested functions with normal Go APIs.

## Code snippet

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int64 = 10
	fmt.Println(unsafe.Sizeof(n))
}
```

## Conclusion

`unsafe` can break memory safety, portability, and future compatibility. Most Go code should not need it.

## What improves if you use this

When unavoidable, it enables low-level integration while keeping the unsafe area contained.

## Final things to know

**Q: Does `unsafe` mean faster?**  
A: Not automatically. It often makes code harder to optimize safely.

**Q: What does cgo cost?**  
A: Build complexity, cross-compilation friction, and call overhead.

**Q: What is the best pattern?**  
A: Small unsafe wrapper, normal safe public API.
