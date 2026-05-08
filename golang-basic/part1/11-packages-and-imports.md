# 11 — Packages & imports

**Goal:** Split code into **packages**, import them, and know **exported** vs **unexported** names.

## What you are learning

- Every file starts with **`package packagename`**. Executable commands use **`package main`**.
- **Import path** string identifies the package (from module path + directory).
- **Exported** identifier: starts with **uppercase** — visible outside the package. **Lowercase** = package-private.
- **`go mod init module/path`** creates the module; sibling directories under the module each become import paths like **`module/path/pkgname`**.

## Try yourself first

1. In a module, create `greeting/greeting.go` with `package greeting` and a function **`func Hello() string`** (exported).
2. In `main.go`, **`import "your/module/greeting"`** and print `greeting.Hello()`.
3. Add an unexported helper **`func tag(s string) string`** in the same package and use it inside `Hello`.

## Reference snippets

### `main.go` importing a local package

Assume module is `example.com/learn` and you have folder `greeting/` with `package greeting`.

```go
package main

import (
	"fmt"

	"example.com/learn/greeting"
)

func main() {
	fmt.Println(greeting.Hello())
}
```

### `greeting/greeting.go`

```go
package greeting

func Hello() string {
	return tag("world")
}

func tag(s string) string {
	return "Hello, " + s
}
```

### Dot import (rare; avoid until you know why you need it)

```go
package main

import (
	. "fmt"
)

func main() {
	Println("no fmt prefix")
}
```

### Blank import (register side effects, e.g. drivers)

```go
package main

import (
	_ "embed"
)
```

## Gotchas

- **Import cycle** is illegal: package A cannot import B if B imports A (directly or indirectly).
- Folder name does not have to match `package` clause, but **one directory = one package** (except `_test` packages).

## Compare

If `hello` is lowercase in `greeting`, `main` cannot call it — compiler error.
