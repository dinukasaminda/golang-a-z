# Part 3 — 02 — Channels

**Goal:** Communicate between goroutines using typed channels.

## Problem

Concurrent functions need to pass results safely. Sharing memory directly can cause data races and confusing ownership.

## Solution

Use `chan T` to send and receive values of type `T`. A channel gives one goroutine a safe handoff point to another goroutine.

## Code snippet

```go
package main

import "fmt"

func worker(out chan<- string) {
	out <- "report ready"
}

func main() {
	ch := make(chan string)
	go worker(ch)
	msg := <-ch
	fmt.Println(msg)
}
```

## Conclusion

Channels are best when you need communication or ownership transfer. They are not a replacement for every shared variable.

## What improves if you use this

Your concurrency design becomes easier to reason about because data moves through explicit paths.

## Final things to know

**Q: What does `chan<- string` mean?**  
A: Send-only channel parameter.

**Q: What does `<-chan string` mean?**  
A: Receive-only channel parameter.

**Q: Does sending block?**  
A: On an unbuffered channel, yes, until another goroutine receives.
