# Part 3 — 09 — Generics

**Goal:** Use type parameters to write reusable, type-safe functions and data structures.

## Problem

Without generics, you duplicate functions for different types or use `any` and lose type safety.

## Solution

Add type parameters in square brackets and constrain what operations are allowed.

## Code snippet

```go
package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5}))
}
```

## Conclusion

Use generics when the algorithm is the same across types and the type parameter makes the API clearer.

## What improves if you use this

You remove duplication while keeping compile-time type checking.

## Final things to know

**Q: What does `~int` mean?**  
A: Any type whose underlying type is `int`.

**Q: Should every helper become generic?**  
A: No. Prefer simple concrete code unless reuse is real.

**Q: Are generics runtime reflection?**  
A: No. They are checked at compile time.
