# Part 3 — 01 — Goroutines

**Goal:** Run work concurrently with `go`, understand that goroutines are lightweight, and wait for them correctly.

## Problem

A program often needs to do more than one slow task: call APIs, read files, process jobs, or handle many requests. Doing everything one by one wastes time while the program waits.

## Solution

Start a goroutine with `go f()`. The function runs concurrently with the caller. Because `main` can exit before goroutines finish, use synchronization such as `sync.WaitGroup`.

## Code snippet

```go
package main

import (
	"fmt"
	"sync"
)

func fetch(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fetched user", id)
}

func main() {
	var wg sync.WaitGroup
	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go fetch(id, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}
```

## Conclusion

Goroutines are for concurrent work, not automatic speed. You still need a way to wait, cancel, and protect shared data.

## What improves if you use this

Independent tasks can overlap, improving throughput and responsiveness.

## Final things to know

**Q: Is a goroutine the same as an OS thread?**  
A: No. The Go runtime schedules many goroutines over fewer OS threads.

**Q: Can goroutines share variables?**  
A: Yes, but shared mutation must be protected with channels, mutexes, or atomics.

**Q: What is the first bug to watch for?**  
A: Starting goroutines and letting `main` exit before they finish.
