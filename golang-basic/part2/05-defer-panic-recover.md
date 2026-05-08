# Part 2 — 05 — `defer`, `panic`, `recover`

**Goal:** Schedule cleanup with **`defer`**, know when **`panic`** is appropriate, and use **`recover`** only inside **deferred** functions.

## What you are learning

- **`defer f()`** runs `f` when the surrounding function **returns** (normal or panic), in **LIFO** order.
- Deferred calls see **final** values of named results in some edge cases — advanced; start with simple cleanup.
- **`panic(v)`** aborts the goroutine, unwinding stacks until **`recover`** or program exit.
- **`recover()`** returns the panic value **only** when called from a **deferred** function while unwinding.
- **Do not** use panic/recover for normal errors — use **`error`**.

## Try yourself first

1. `defer` several `fmt.Println` calls and observe print order.
2. `defer` a function that closes a file after `os.Open` (or use `fmt.Println` to simulate “cleanup”).
3. Write `safe()` that defers `recover()` and catches a deliberate `panic` in an inner function.

## Reference snippets

### `defer` order

```go
package main

import "fmt"

func main() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")
}
```

### `defer` with closure capturing loop variable (Go 1.22+ fixed loop var; still understand the pattern)

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
```

### `defer` for file close

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	fmt.Fprintln(f, "hello")
}
```

### `recover`

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	panic("boom")
}
```

## Gotchas

- **`defer` in loop** without care can defer **many** calls until function exit — sometimes you want an **anonymous function** to scope defer per iteration.
- **`recover` outside defer** does nothing useful.

## Compare

Output order for stacked defers should be **last deferred runs first**.
