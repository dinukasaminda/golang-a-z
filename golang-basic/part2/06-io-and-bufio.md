# Part 2 — 06 — `io` & `bufio`

**Goal:** Read and write through **`io.Reader`** and **`io.Writer`**, and buffer with **`bufio`**.

## What you are learning

- **`io.Reader`**: `Read(p []byte) (n int, err error)` — streaming bytes; **`io.EOF`** means end of input.
- **`io.Writer`**: `Write(p []byte) (n int, err error)`.
- Helpers: **`io.Copy`**, **`io.ReadAll`** (read whole reader into memory — mind size).
- **`bufio.Scanner`** for line-by-line text; **`bufio.Reader`** for peek/read strings with delimiter.

## Try yourself first

1. Use **`strings.NewReader`** as an `io.Reader` and **`io.ReadAll`** to get a `[]byte`, convert to string.
2. Use **`bufio.NewScanner`** on a string reader and print lines from a multi-line string.
3. Write bytes to **`bytes.Buffer`** (it implements `Writer`).

## Reference snippets

### `io.Copy` and `io.ReadAll`

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello")
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, r)
	fmt.Println(buf.String())

	data, err := io.ReadAll(strings.NewReader("world"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```

### `bufio.Scanner` lines

```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "line1\nline2\n"
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
}
```

### `bufio.Reader` read string until delimiter

```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := bufio.NewReader(strings.NewReader("hello,world"))
	part, err := r.ReadString(',')
	if err != nil {
		panic(err)
	}
	fmt.Print(part)
}
```

## Gotchas

- **`Scanner`** has max token size; huge lines need **`bufio.Reader`** or custom split.
- Always check **`sc.Err()`** after the loop.

## Compare

Your line loop should print two lines for the two-line string example.
