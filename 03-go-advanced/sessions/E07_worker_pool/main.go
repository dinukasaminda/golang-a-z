package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("starting worker: ", id)
	for j := range jobs {
		fmt.Println("Worker: ", id, " Job: ", j)
		results <- j * j
	}

	fmt.Println("stopping worker: ", id)

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
