package main

import (
	"fmt"
	"os"
)

func doRequest() {
	defer fmt.Println("Defer1")
	defer fmt.Println("Defer2")
	fmt.Println("Sending request")
}

func main() {
	doRequest() // Defer work LIFO

	// File close using defer
	f, err := os.CreateTemp("./", "sample")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	fmt.Fprintln(f, "hello")

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)

		}
	}()
	panic("boom")
}
