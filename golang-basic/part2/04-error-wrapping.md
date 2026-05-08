# Part 2 — 04 — Error wrapping

**Goal:** Chain errors with **`fmt.Errorf` + `%w`**, unwrap with **`errors.Is`** and **`errors.As`**.

## What you are learning

- **`fmt.Errorf("...: %w", err)`** attaches `err` into a **wrap chain** (Go 1.13+).
- **`errors.Is(err, target)`** reports whether `err` or any unwrap equals **`target`** (good for sentinel errors).
- **`errors.As(err, &ptr)`** finds the first error in the chain assignable to `*ptr` (good for custom error types).
- **`errors.Unwrap(err)`** moves one link down the chain; rarely needed directly.

## Try yourself first

1. Define **`var ErrNotFound = errors.New("not found")`**.
2. Write `find(id int) error` that returns **`fmt.Errorf("lookup %d: %w", id, ErrNotFound)`** for a sentinel case.
3. In `main`, use **`errors.Is`** to detect `ErrNotFound`.

## Reference snippets

### `errors.Is` with wrapped error

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func find(id int) error {
	if id < 0 {
		return fmt.Errorf("find %d: %w", id, ErrNotFound)
	}
	return nil
}

func main() {
	err := find(-1)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("handled not found")
	}
}
```

### `errors.As` with custom type

```go
package main

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	Code int
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http %d", e.Code)
}

func doRequest() error {
	return fmt.Errorf("request: %w", &HTTPError{Code: 404})
}

func main() {
	err := doRequest()
	var he *HTTPError
	if errors.As(err, &he) {
		fmt.Println("status", he.Code)
	}
}
```

### Unwrap manually

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("outer: %w", errors.New("inner"))
	fmt.Println(errors.Unwrap(err))
}
```

## Gotchas

- **`%v`** in `fmt.Errorf` does **not** wrap — use **`%w`** only once per `Errorf` and only for the error being wrapped.
- Do not compare wrapped errors with **`==`** to a sentinel; use **`errors.Is`**.

## Compare

Your `errors.Is(find(-1), ErrNotFound)` should be **true**.
