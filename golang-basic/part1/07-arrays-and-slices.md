# 07 — Arrays & slices

**Goal:** Know **arrays** (fixed size, value type) vs **slices** (dynamic view over an array), and use **`append`**, **`len`**, **`cap`**.

## What you are learning

- **Array:** **`[N]T`** — length is part of the type; copying an array copies all elements.
- **Slice:** **`[]T`** — descriptor with pointer, length, capacity; **slice literals** like `[]int{1,2,3}`.
- **`make([]T, len, cap)`** allocates a slice backed by a new array.
- **`append(s, x)`** may allocate a new backing array if capacity is exceeded; always assign the result: **`s = append(s, x)`**.
- **Slicing:** **`s[low:high]`** — half-open interval; **`s[:]`** full length; **`s[1:]`** drop first.
- Two slices can **share** the same backing array — mutating an element visible through both can surprise you.

## Try yourself first

1. Create a slice of three ints, `append` two more, print length and capacity (explore how `cap` grows).
2. Use `copy` to copy one slice into another.
3. Show a subslice sharing memory: change `s[0]` and see it reflected in a sub-slice (or the reverse).

## Reference snippets

### Slice literal, `make`, `append`

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 0, 8)
	b = append(b, 10)
	fmt.Println(b, len(b), cap(b))
}
```

### Array vs slice type

```go
package main

import "fmt"

func main() {
	var arr [3]int // array: size fixed
	arr[0] = 1

	s := arr[:] // slice view of array
	fmt.Println(arr, s)
}
```

### `copy`

```go
package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	n := copy(dst, src)
	fmt.Println(n, dst)
}
```

### Shared backing array (careful)

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	t := s[1:3]
	t[0] = 99
	fmt.Println(s) // s[1] is now 99
}
```

## Gotchas

- **`nil` slice**: `var s []int` is `nil`; `len`/`cap` are 0; `append` works. A `nil` slice is often fine.
- Reading **`s[i]`** with `i >= len(s)` **panics**, not an error.

## Compare

Always **`s = append(s, ...)`**; check your mental model of **shared subslices** before mutating.
