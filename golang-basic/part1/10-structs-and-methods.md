# 10 — Structs & methods

**Goal:** Group fields with **`struct`**, attach **methods** with value vs **pointer receivers**.

## What you are learning

- **`type Name struct { ... }`** defines a new struct type.
- **Embedded fields** (no field name) are a composition feature — covered again in Part 2.
- **Methods:** **`func (receiver Type) MethodName(...) ...`** — `receiver` is like `self` but explicit.
- **Pointer receiver** **`*Type`**: method can mutate the struct and avoids copying large structs; call still works as **`v.Method()`** — Go adds `&` for you when possible.
- **Value receiver**: method gets a copy; mutations do not affect the original.

## Try yourself first

1. Define `type User struct { Name string; Age int }`, create a value, print a field.
2. Add `func (u User) Label() string` returning a formatted string.
3. Add `func (u *User) IncAge()` and call it — observe why pointer receiver matters.

## Reference snippets

### Struct literal and fields

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Ada", Age: 36}
	fmt.Println(u.Name, u.Age)

	p := &User{Name: "Lin", Age: 28}
	fmt.Println(p.Name) // same as (*p).Name
}
```

### Value receiver (copy)

```go
package main

import "fmt"

type Counter struct{ n int }

func (c Counter) Value() int { return c.n }

func main() {
	c := Counter{n: 3}
	fmt.Println(c.Value())
}
```

### Pointer receiver (mutate)

```go
package main

import "fmt"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

func main() {
	c := &Counter{}
	c.Inc()
	fmt.Println(c.n)
}
```

### Method on non-struct type (named type)

```go
package main

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", c)
}

func main() {
	t := Celsius(36.6)
	fmt.Println(t.String())
}
```

## Gotchas

- **Consistency:** if one method needs `*T`, usually **all** methods on `T` use `*T` for a clean API.
- **`nil` pointer receiver:** methods can be called on `nil` if the method handles `nil` — useful for some types, dangerous if forgotten.

## Compare

After `Inc`, the counter behind `*Counter` should increase; a value-receiver `Inc` would not change the original unless you reassign (pattern not idiomatic).
