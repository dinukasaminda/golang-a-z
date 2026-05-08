# 12 — Errors — basics

**Goal:** Return and check **`error`**, create errors with **`errors.New`**, and use **`fmt.Errorf`** for formatted messages (wrapping comes in Part 2).

## What you are learning

- **`error`** is an interface: **`type error interface { Error() string }`**.
- **Idiom:** functions return **`(T, error)`**; **`nil` error** means success.
- **`errors.New("msg")`** creates a simple error value.
- **`fmt.Errorf("format", args...)`** builds an error string; with **`%w`** you wrap another error (Part 2).
- Always handle errors explicitly — **`if err != nil { return ..., err }`** or log/return up the stack.

## Try yourself first

1. Write `sqrt(x float64) (float64, error)` returning an error for negative `x`.
2. In `main`, call it twice (ok and error case) and print results or `err.Error()`.
3. Return **`fmt.Errorf("sqrt: negative %v", x)`** for the bad case.

## Reference snippets

### `errors.New` and `if err != nil`

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func main() {
	v, err := sqrt(9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	if _, err := sqrt(-1); err != nil {
		fmt.Println(err)
	}
}
```

### `fmt.Errorf` (message only)

```go
package main

import (
	"fmt"
)

func parseID(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty id")
	}
	return 42, nil
}

func main() {
	_, err := parseID("")
	fmt.Println(err)
}
```

### Sentinel error (named var to compare with `errors.Is` later)

```go
package main

import "errors"

var ErrNotFound = errors.New("not found")
```

## Gotchas

- **`panic`** is not for normal errors; use **`error`** return values.
- Comparing errors with **`==`** works for simple values like **`errors.New`** only if it is the **same variable**; for wrapped errors use **`errors.Is`** (Part 2).

## Compare

Your negative-input path must return **non-nil** `error` and a sensible zero result (often `0`).
