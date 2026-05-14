# Part 3 — 07 — Worker pools

**Goal:** Limit concurrency while processing many jobs.

## Problem

Starting one goroutine per job can overload databases, APIs, CPU, or memory when job count is large.

## Solution

Create a fixed number of workers. Send jobs through a channel and collect results separately.

## Code snippet

```go
package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job", j)
		results <- j * j
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// we have to set a reciver for the results channel before starting sending data
	// otherwise worker will be blocked -> result to deadlock
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-results)
		}
	}()

	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)
}
```

## Conclusion

A worker pool gives you backpressure: only a controlled number of jobs run at once.

## What improves if you use this

You protect downstream systems and make throughput more predictable.

## Final things to know

**Q: What decides worker count?**  
A: CPU count for CPU-bound work; external capacity for I/O-bound work.

**Q: Who closes the jobs channel?**  
A: The sender.

**Q: Who closes the results channel?**  
A: Usually a coordinator after all workers finish.
