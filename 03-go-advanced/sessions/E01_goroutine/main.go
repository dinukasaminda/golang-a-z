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
