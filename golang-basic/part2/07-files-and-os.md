# Part 2 — 07 — Files & `os`

**Goal:** Create, open, read, write, and remove files; read **environment** and **arguments**.

## What you are learning

- **`os.Open`**, **`os.Create`**, **`os.OpenFile`** with flags.
- **`os.ReadFile` / `os.WriteFile`** (simple whole-file helpers, Go 1.16+).
- **`os.Getenv`**, **`os.Args`**.
- Always handle **`error`** from I/O; **`defer f.Close()`** after successful open.

## Try yourself first

1. Write a string to a temp file with **`os.CreateTemp`**, read it back with **`os.ReadFile`**.
2. Print **`os.Args`** and one environment variable.
3. Use **`os.MkdirTemp`** and clean up with **`os.RemoveAll`**.

## Reference snippets

### Temp file write and read

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "demo-*.txt")
	if err != nil {
		panic(err)
	}
	path := f.Name()
	defer os.Remove(path)
	defer f.Close()

	if _, err := f.Write([]byte("data")); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
```

### `ReadFile` / `WriteFile`

```go
package main

import "os"

func main() {
	path := "sample.txt"
	_ = os.WriteFile(path, []byte("hi"), 0o644)
	_, _ = os.ReadFile(path)
	_ = os.Remove(path)
}
```

### Args and env

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("HOME"))
}
```

## Gotchas

- File permissions: **`0o644`** is common for user-writable files; secrets should not be world-readable.
- On Windows, paths still use **`string`**; use **`filepath`** package for cross-platform joining (next step in your learning).

## Compare

Ensure you **`Close`** before **`ReadFile`** the same path if you need data flushed (or use `WriteFile`).
