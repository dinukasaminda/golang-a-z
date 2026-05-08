# 04 — Control flow

**Goal:** Use `if`, `for`, and `switch` the Go way. Go has **only `for`** for loops (no `while` keyword).

## What you are learning

- **`if`** may include a short statement before the condition: `if x := f(); x > 0 { ... }` — `x` is scoped to the `if`/`else`.
- **`for`** forms:
  - C-style: `for i := 0; i < n; i++`
  - Condition only: `for condition` (like `while`)
  - Infinite: `for { }`
  - Range: `for i, v := range slice` (or `for k, v := range map`)
- **`switch`** does not fall through by default; use `fallthrough` if you really want that.
- **`switch` with no tag** can replace long `if-else` chains.

## Try yourself first

1. Print numbers `0..4` with a `for` loop.
2. Sum a slice using `range`.
3. Use `switch` on a string for a few commands (`"hello"`, `"bye"`, default).

## Reference snippets

### `if` with short declaration

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	if s := strings.TrimSpace("  go  "); s != "" {
		fmt.Println(s)
	}
}
```

### `for` loops

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	n := 0
	for n < 3 {
		fmt.Println("while-like", n)
		n++
	}

	for {
		break // exit; or use return from main
	}
}
```

### `range` on slice and map

```go
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

### `switch`

```go
package main

import "fmt"

func main() {
	cmd := "ping"
	switch cmd {
	case "ping":
		fmt.Println("pong")
	case "quit", "exit":
		fmt.Println("bye")
	default:
		fmt.Println("unknown")
	}

	x := 42
	switch {
	case x < 0:
		fmt.Println("neg")
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("pos")
	}
}
```

## Gotchas

- **`for range`**: if you only need the value, use **`for _, v := range`**; if only index/key, **`for k := range m`**.
- `break` only exits the **innermost** `for`/`switch`; labels exist for nested cases (learn later if needed).

## Compare

Ensure you did not use a non-existent `while` keyword — in Go it is `for cond { }`.
