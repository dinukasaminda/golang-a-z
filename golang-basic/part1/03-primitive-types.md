# 03 — Primitive types

**Goal:** Use numeric types, `bool`, `string`, and **explicit conversions** (Go never silently converts).

## What you are learning

- Integers: **`int`**, **`int8`**, **`int16`**, **`int32`**, **`int64`**, and unsigned **`uint`** variants; **`byte`** is alias for **`uint8`**.
- Floats: **`float32`**, **`float64`** (default for floating literals in many contexts).
- **`bool`**: only `true` and `false`.
- **`string`**: immutable sequence of bytes (UTF-8 text); covered in more detail in topic 09.
- **Conversion:** **`T(value)`** converts `value` to type `T` when the rules allow; there is no implicit `int` → `int64` or `int` → `float64`.

## Try yourself first

1. Declare an `int` and an `int64`, assign numbers, and convert between them with `int64(x)` and `int(y)`.
2. Divide two integers and observe integer division; then use `float64` to get a fractional result.
3. Use `bool` in an `if`.

## Reference snippets

### Integer and float conversions

```go
package main

import "fmt"

func main() {
	var a int = 10
	var b int64 = int64(a)

	var x int = 3
	var y int = 2
	fmt.Println(x / y) // 1 integer division

	fmt.Println(float64(x) / float64(y)) // 1.5

	_ = b
}
```

### Explicit width types

```go
package main

import "fmt"

func main() {
	var u uint8 = 255
	var i int32 = -100
	fmt.Println(u, i)
}
```

### `bool`

```go
package main

import "fmt"

func main() {
	ok := true
	if ok {
		fmt.Println("ok")
	}
}
```

## Gotchas

- **`int` size** is either 32 or 64 bits depending on the platform; when you care about wire format or APIs, prefer **`int32`/`int64`** explicitly.
- Converting **`float64` to `int`** truncates toward zero; rounding is your responsibility.

## Compare

Check that you used **`T(v)`** everywhere you crossed types, not assignment without conversion.
