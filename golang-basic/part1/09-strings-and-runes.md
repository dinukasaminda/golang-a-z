# 09 — Strings & runes

**Goal:** Treat a Go **`string`** as **UTF-8 bytes**, use **`rune`** for Unicode code points, and iterate without breaking multi-byte characters.

## What you are learning

- **`string`** is read-only bytes; **`len(s)`** is **byte length**, not character count.
- **`rune`** is alias for **`int32`** and holds one Unicode code point.
- **Range over string** by **`for _, r := range s`** gives **runes**, not raw bytes.
- Package **`unicode/utf8`**: **`utf8.RuneCountInString`**, **`utf8.DecodeRuneInString`**, etc.
- Conversions: **`[]byte(s)`** and **`string(b)`** copy data; use carefully in hot paths.

## Try yourself first

1. Print `len` of `"Hello"` and of `"こんにちは"` — compare byte length vs perceived “length”.
2. Loop with `for i := 0; i < len(s); i++` on a string containing multi-byte characters and observe wrong cuts; then use `range` over runes.
3. Use `utf8.RuneCountInString` on the same string.

## Reference snippets

### Byte length vs rune iteration

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "こんにちは"
	fmt.Println(len(s)) // bytes
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%d: %U %c\n", i, r, r)
	}
}
```

### Building strings efficiently (preview: use `strings.Builder` in real code)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("a")
	b.WriteString("b")
	fmt.Println(b.String())
}
```

### Runes and string from runes

```go
package main

import "fmt"

func main() {
	rs := []rune{'A', 'Ω', '世'}
	fmt.Println(string(rs))
}
```

## Gotchas

- **Indexing `s[i]`** gives a **byte**, not a rune; for arbitrary Unicode, index by byte position only when you know layout or use **`range`** / `utf8` helpers.
- Comparing strings is lexicographic by **bytes** (UTF-8 order), usually what you want for sorting UTF-8 text.

## Compare

Your `len` on Japanese should be **much larger than 5** in bytes; `RuneCountInString` should be **5** for five kana in the example.
