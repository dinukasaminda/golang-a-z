package main

import "fmt"

func worker(out chan<- string) {
	out <- "report ready"
}

func main() {
	// Unbuffered channel
	ch := make(chan string)
	go worker(ch)

	msg := <-ch
	fmt.Println(msg)
}
