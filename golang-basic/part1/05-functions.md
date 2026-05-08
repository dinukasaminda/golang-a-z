# 05 — Functions

**Goal:** Define functions with parameters, return values, **multiple returns**, and **named results**.

## What you are learning

- Syntax: **`func name(params) (results) { body }`**.
- **Multiple return values** are idiomatic for **`(result, error)`**.
- **Named return values** declare names in the signature; a naked **`return`** returns the current values of those names (easy to misuse — use when they clarify, not when they obscure).
- **Variadic** parameters: **`nums ...int`** collects arguments into a slice inside the function.
- Functions are **first-class values** (can assign to variables, pass around) — a small preview used in snippets.

## Try yourself first

1. Write `add(a, b int) int`.
2. Write `div(a, b int) (int, error)` returning an error when `b == 0` (use `errors.New` from topic 12, or return `(0, fmt.Errorf(...))` if you already import `fmt`).
3. Write a variadic `sum(vals ...int) int`.

## Reference snippets

### Simple function and multiple returns

```go
package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(add(2, 3))
	q, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q)
}
```

### Named results

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

### Variadic function

```go
package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	s := []int{10, 20}
	fmt.Println(sum(s...))
}
```

### Function value

```go
package main

import "fmt"

func main() {
	var f func(int) int
	f = func(x int) int { return x * 2 }
	fmt.Println(f(21))
}
```

## Gotchas

- **Parameter types** can be shortened: **`a, b int`** instead of `a int, b int`.
- Empty `return` with named results returns their **current** values — always initialise them or return explicitly to avoid subtle bugs.

## Compare

Your `div` should return a non-nil `error` on divide-by-zero; variadic call with a slice must use **`slice...`**.
