# Part 2 — 08 — `time` & duration

**Goal:** Use **`time.Time`**, **`time.Duration`**, parse and format, and simple scheduling.

## What you are learning

- **`time.Now()`** returns current local time; **`time.Now().UTC()`** for UTC.
- **`time.Duration`** is an `int64` nanoseconds; literals like **`5 * time.Second`**.
- **`time.Parse`** and **`Format`** use **reference time**: **`Mon Jan 2 15:04:05 MST 2006`** as layout **`2006-01-02 15:04:05`** style templates.
- **`time.After(d)`** returns `<-chan time.Time` once after `d`.
- **`time.NewTicker(d)`** ticks repeatedly; **`Stop()`** when done.

## Try yourself first

1. Print current time in RFC3339 (`time.RFC3339`).
2. Parse a string date with **`time.Parse`** in UTC.
3. Use **`time.Sleep`** for 100ms and measure elapsed with **`time.Since(start)`**.

## Reference snippets

### Format and parse

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(t.Format(time.RFC3339))

	s := "2026-05-01T12:00:00Z"
	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsed.Equal(t))
}
```

### Duration and `Since`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	fmt.Println(time.Since(start).Milliseconds())
}
```

### `Ticker`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(200 * time.Millisecond)
	defer t.Stop()
	for i := 0; i < 3; i++ {
		<-t.C
		fmt.Println("tick", i)
	}
}
```

## Gotchas

- **`time.Parse`** without location: use layouts that include zone or pass **`time.UTC`** via `ParseInLocation` when needed.
- **`Timer` / `Ticker`** must be **stopped** or they leak (for `Ticker`, always `defer t.Stop()`).

## Compare

Your RFC3339 string should round-trip parse to an equal instant when zones match.
