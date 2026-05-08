# Part 2 — 01 — Interfaces

**Goal:** Define **behaviour-only** types, rely on **implicit implementation**, and prefer **small interfaces**.

## What you are learning

- An **interface** lists method signatures. A type **satisfies** the interface if it has those methods — **no `implements` keyword**.
- **Empty interface** **`interface{}`** or **`any`** (Go 1.18+) means “any value”; use sparingly at boundaries.
- **Accept interfaces, return concrete types** — functions should take `io.Reader`, not `*os.File`, when possible.
- **`nil` interface value** vs **`nil` concrete value in an interface** — a subtle gotcha (see snippet).

## Try yourself first

1. Define `type Speaker interface { Say() string }` and two structs that implement it.
2. Write `func Greet(s Speaker)` and pass different concrete types.
3. Read about the nil interface gotcha in the reference snippet and reproduce it mentally.

## Reference snippets

### Small interface and implicit satisfaction

```go
package main

import "fmt"

type Speaker interface {
	Say() string
}

type Dog struct{}

func (Dog) Say() string { return "woof" }

type Robot struct{}

func (Robot) Say() string { return "beep" }

func Greet(s Speaker) {
	fmt.Println(s.Say())
}

func main() {
	Greet(Dog{})
	Greet(Robot{})
}
```

### Interface with common stdlib: `fmt.Stringer`

```go
package main

import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func main() {
	fmt.Println(Point{3, 4})
}
```

### Nil interface gotcha (concrete `*T` is nil but interface value is not “typed nil”)

```go
package main

import "fmt"

type Speaker interface {
	Say() string
}

type Cat struct{}

func (c *Cat) Say() string {
	if c == nil {
		return "quiet cat"
	}
	return "meow"
}

func main() {
	var c *Cat
	var s Speaker = c
	fmt.Println(s.Say()) // ok if method handles nil *Cat
}
```

## Gotchas

- Pointer receiver means only **`*T`** satisfies the interface unless you only use value receivers.
- Large interfaces are hard to mock and reuse — **split** them.

## Compare

If you used value receiver on a type but held `*T` in an interface, ensure method set matches what you expect.
