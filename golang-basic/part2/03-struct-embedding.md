# Part 2 — 03 — Struct embedding

**Goal:** Compose types by **embedding** structs — promoted fields and methods.

## What you are learning

- **Embedding:** anonymous field **`Inner`** inside **`Outer`** — `Outer` gets **promoted** fields/methods like **`outer.Field`** if unambiguous.
- This is **not inheritance** — there is no subtype polymorphism; it is **composition** with syntactic sugar.
- Useful for **wrapping** and **layering** (e.g. `http.ResponseWriter` wrappers).

## Try yourself first

1. Embed a struct with `Name string` inside `Person` and set `person.Name` from `main`.
2. Embed a type that has a method and call it on the outer struct value.

## Reference snippets

### Promoted field

```go
package main

import "fmt"

type Metadata struct {
	ID   int
	Tags []string
}

type Document struct {
	Metadata
	Title string
}

func main() {
	d := Document{
		Metadata: Metadata{ID: 1, Tags: []string{"draft"}},
		Title:    "Notes",
	}
	fmt.Println(d.ID, d.Title)
}
```

### Promoted method

```go
package main

import "fmt"

type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("log:", msg)
}

type Service struct {
	Logger
	Name string
}

func main() {
	s := Service{Name: "api", Logger: Logger{}}
	s.Log("starting")
}
```

### Name clash (outer wins; inner must be qualified)

```go
package main

import "fmt"

type Inner struct{ Name string }
type Outer struct {
	Inner
	Name string
}

func main() {
	o := Outer{Inner: Inner{Name: "inner"}, Name: "outer"}
	fmt.Println(o.Name)
	fmt.Println(o.Inner.Name)
}
```

## Gotchas

- Promotion can make APIs **harder to read** if overused — sometimes a **named field** is clearer.
- Embedding **pointer** `*T` is allowed; `nil` embedded pointer can panic if method called without nil check.

## Compare

You should access embedded fields both as **`d.Metadata.ID`** and **`d.ID`** when unambiguous.
