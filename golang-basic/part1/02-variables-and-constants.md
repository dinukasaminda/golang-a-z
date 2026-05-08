# 02 — Variables & constants

**Goal:** Declare variables with `var` and `:=`, use `const`, and understand **zero values**.

## What you are learning

- **`var name T`** declares a variable of type `T`; without an initializer it gets the **zero value** for `T` (`0`, `0.0`, `false`, `""`, `nil` for pointers/slices/maps/channels/functions/interfaces).
- **`var name = value`** lets the compiler infer the type.
- **Short declaration:** **`name := value`** only inside functions; declares and assigns in one step.
- **`const`** defines compile-time constants (numbers, strings, booleans, or constant expressions).
- **`iota`** inside a `const` block generates incrementing untyped integer constants (often used for enums).

## Try yourself first

1. Declare an `int`, a `string`, and a `bool` with `var` and print them (see zero values).
2. Use `:=` to declare and assign in `main`.
3. Add a `const` for a string and a `const` block with `iota` for three named constants.

## Reference snippets

### `var` with zero value and with initializer

```go
package main

import "fmt"

func main() {
	var count int
	var name string
	fmt.Println(count, name) // 0

	var x = 42
	var y int = 10
	fmt.Println(x, y)
}
```

### Short declaration `:=`

```go
package main

import "fmt"

func main() {
	n := 100
	s := "short"
	fmt.Println(n, s)
}
```

### Multiple `var` and short declaration

```go
package main

import "fmt"

func main() {
	var a, b int = 1, 2
	c, d := 3, 4
	fmt.Println(a, b, c, d)
}
```

### `const` and `iota`

```go
package main

import "fmt"

const Pi = 3.14159

const (
	Red = iota // 0
	Green      // 1
	Blue       // 2
)

func main() {
	fmt.Println(Pi, Red, Green, Blue)
}
```

## Gotchas

- `:=` **declares** new variables; you cannot use it if every name on the left already exists in that scope (use `=` for assignment only).
- Package-level variables use `var`; **`:=` is not allowed** at package level.

## Compare

Your zero-value prints should match the type rules; `iota` values should be `0, 1, 2` in the example above.
