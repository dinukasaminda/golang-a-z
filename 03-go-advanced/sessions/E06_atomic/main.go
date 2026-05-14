package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	atomicInt := atomic.Int32{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicInt.Add(1)
		}()

	}
	wg.Wait()

	fmt.Println("Value: ", atomicInt.Load())

}
