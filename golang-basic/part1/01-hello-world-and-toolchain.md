# 01 — Hello World & toolchain

**Goal:** Run a minimal Go program and know the usual commands. **You** create `main.go` in a folder; this file only explains and shows reference snippets.

## What you are learning

- Every Go program starts in **`package main`** and needs a **`func main()`** as the entry point.
- **`go run .`** compiles and runs the package in the current directory (good while learning).
- **`go build`** produces an executable binary (often named after the folder on Unix/macOS).
- **`go fmt ./...`** formats your code to the standard style; most editors run it on save.
- **`go mod init example.com/hello`** (or any module path) creates **`go.mod`** so the folder is a proper module.

## Try yourself first

1. Create a folder, run `go mod init` with a module path you choose.
2. Add `main.go` with `package main`, import `fmt`, and print one line from `main`.
3. Run with `go run .` and build with `go build`.

## Reference snippets

### Minimal program

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go.")
}
```

### Same program with a grouped import (idiomatic when you have several imports)

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go.")
}
```

### `go.mod` (after `go mod init example.com/hello`)

```text
module example.com/hello

go 1.22
```

(Your `go` line may differ depending on the Go version installed.)

## Gotchas

- File must be **`package main`** and contain **`func main()`** for a runnable command; other packages use `package somethingelse` and are built as libraries.
- The **module path** in `go mod init` can be a real URL you own or a placeholder like `example.com/learn`; it must be unique if you publish the module.

## Compare

After you code, check: correct package, imports only what you use, `main` has no parameters and no return values.
