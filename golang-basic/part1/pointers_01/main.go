package main

import "fmt"

func inc(p *int) {
	// this is checking the address is null or not
	if p == nil {
		return
	}
	// otherwise,
	// dereferencing for nill address will panic

	// we increament the value of that int address by one
	*p++
}

func main() {
	x := 10
	fmt.Println("X: ", x)
	inc(&x)
	fmt.Println("New x: ", x)
}
