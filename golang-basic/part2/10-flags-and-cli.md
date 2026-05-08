# Part 2 — 10 — Flags & CLI

**Goal:** Parse flags with **`flag`** and read raw **`os.Args`**.

## What you are learning

- **`flag.String`**, **`flag.Int`**, etc. return **pointers** to variables filled by **`flag.Parse()`**.
- **`flag.StringVar`** binds an existing variable.
- Non-flag arguments: **`flag.Args()`** after `Parse`.
- **`os.Args[0]`** is the program name; user args start at index 1.

## Try yourself first

1. Define `-name` string and `-n` int with defaults, parse, print values.
2. Print remaining positional args after flags.
3. Know that **`flag.Parse()`** calls **`os.Exit(2)`** on bad flags by default.

## Reference snippets

### Basic flags

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "guest", "your name")
	n := flag.Int("n", 1, "repeat count")
	flag.Parse()

	for i := 0; i < *n; i++ {
		fmt.Println(*name)
	}
	fmt.Println("args:", flag.Args())
}
```

### `Var` bindings

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.Parse()
	fmt.Println(port)
}
```

### Raw `os.Args`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		fmt.Println(i, a)
	}
}
```

## Gotchas

- Call **`flag.Parse()`** once, typically near the start of `main`, after defining flags.
- For subcommands (`git commit`-style), the stdlib `flag` package is low-level; many projects use **`cobra`** later.

## Compare

Running `go run . -name=Go -n 2 a b` should print `Go` twice and `args: [a b]`.
