# 08 — Maps

**Goal:** Create maps, set and get keys, test existence, delete, and iterate.

## What you are learning

- Map type: **`map[K]V`** — keys must be **comparable** (`==` works); `K` cannot be a slice, map, or function.
- **Zero value** of a map is **`nil`**; writing to a `nil` map **panics**; use **`make(map[K]V)`** or a map literal.
- **Get:** `v := m[k]` — if key missing, `v` is the **zero value** for `V` (you cannot tell missing from “stored zero” with one value).
- **Comma-ok:** **`v, ok := m[k]`** — `ok` is `true` if the key exists.
- **Delete:** **`delete(m, k)`**.
- **Iteration order** over maps is **randomised** — do not depend on order.

## Try yourself first

1. Build a `map[string]int` counting how many times each word appears in a small slice of strings.
2. Use `v, ok := m["missing"]` and branch on `ok`.
3. Delete a key and show behaviour when reading it again.

## Reference snippets

### Map literal and `make`

```go
package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["c"] = 3

	n := make(map[int]string)
	n[1] = "one"
	fmt.Println(m, n)
}
```

### Comma-ok idiom

```go
package main

import "fmt"

func main() {
	m := map[string]int{"x": 10}
	if v, ok := m["x"]; ok {
		fmt.Println("found", v)
	}
	if _, ok := m["y"]; !ok {
		fmt.Println("missing y")
	}
}
```

### `delete` and iteration

```go
package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	delete(m, "a")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

### Word count pattern

```go
package main

import "fmt"

func main() {
	words := []string{"go", "go", "rust", "go"}
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}
	fmt.Println(count)
}
```

## Gotchas

- **Concurrency:** default maps are **not** safe for concurrent write (or read while write); use `sync.Map` or a mutex later.
- **`nil` map:** `var m map[string]int` then `m["k"]=1` **panics** — call `make` first.

## Compare

Your word count should use **`count[w]++`** safely after `make`.
