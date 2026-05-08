# 06 — Pointers

**Goal:** Use **`&`**, **`*`**, know when values are copied, and avoid dereferencing **`nil`**.

## What you are learning

- **`&x`** takes the **address** of `x` (type `*T` if `x` is `T`).
- **`*p`** **dereferences** pointer `p` and gives the value it points to.
- **Zero value** of a pointer type is **`nil`**; dereferencing `nil` **panics**.
- Go passes **arguments by value**. Passing a pointer lets the callee mutate the caller’s variable (through the shared memory).
- Structs: **`&MyStruct{...}`** gives a pointer to a new struct literal.

## Try yourself first

1. Declare `x := 10`, pass `&x` to a function that increments through the pointer, print `x` after the call.
2. Show that passing `x` by value to a function does **not** change `x` in `main`.
3. Create a `*int` variable (nil), and use `if p != nil` before dereferencing.

## Reference snippets

### Pointer parameter mutates caller state

```go
package main

import "fmt"

func inc(p *int) {
	if p == nil {
		return
	}
	*p++
}

func main() {
	x := 10
	inc(&x)
	fmt.Println(x) // 11
}
```

### Value copy does not mutate

```go
package main

import "fmt"

func tryInc(n int) {
	n++
}

func main() {
	x := 10
	tryInc(x)
	fmt.Println(x) // still 10
}
```

### Pointer to new struct

```go
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	p := &Point{X: 3, Y: 4}
	fmt.Println(p.X, (*p).Y) // dot works on pointers to structs too
}
```

### `new(T)` allocates zeroed `T` and returns `*T`

```go
package main

import "fmt"

func main() {
	p := new(int)
	*p = 7
	fmt.Println(*p)
}
```

## Gotchas

- For **pointer receivers** on methods, see topic 10; rule of thumb: if the method must mutate the receiver or it is large, use `*T`.
- **`p.Field`** is allowed when `p` is `*Struct`; Go simplifies pointer dereference for fields.

## Compare

After `inc(&x)`, `x` should change; after `tryInc(x)`, it should not.
