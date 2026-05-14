package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		result <- "done"
	}()

	select {
	case v := <-result:
		fmt.Println("v:", v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

}
