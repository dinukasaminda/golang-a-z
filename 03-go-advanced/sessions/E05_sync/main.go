package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func main() {
	c := Counter{
		mu: sync.Mutex{},
		n:  0,
	}

	for i := 0; i < 10; i++ {
		go c.Inc()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final value: ", c.n)
}
