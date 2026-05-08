# Part 2 — 02 — Type assertions & type switches

**Goal:** Recover concrete types from an **`interface{}` / `any`** or a **wider interface** safely.

## What you are learning

- **Type assertion:** **`v := x.(T)`** — panics if `x` is not `T`. **Safe form:** **`v, ok := x.(T)`** — `ok` is false if not `T`.
- **Type switch:** **`switch x := i.(type) { case int: ... case string: ... }`**
- Use this at **boundaries** (JSON decode, plugin APIs) where the static type is lost.

## Try yourself first

1. Write `func describe(i any)` that prints whether `i` is `int`, `string`, or something else (type switch).
2. Use the two-value assertion to try `float64` and fall back to a default message.

## Reference snippets

### Safe type assertion

```go
package main

import "fmt"

func main() {
	var i any = "hello"
	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	}
	_, ok = i.(int)
	if !ok {
		fmt.Println("not an int")
	}
}
```

### Type switch

```go
package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Printf("other %T\n", v)
	}
}

func main() {
	describe(42)
	describe("go")
	describe(3.14)
}
```

### Assertion on interface with methods

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("abc")
	if sr, ok := r.(io.ReadSeeker); ok {
		fmt.Println("supports seek", sr)
	}
}
```

## Gotchas

- **Never** use the one-value assert on untrusted dynamic data — use **`ok`** form.
- In a type switch, **`v`** has the type of each **case** in that clause.

## Compare

Your `describe` should handle at least three kinds without panicking.
