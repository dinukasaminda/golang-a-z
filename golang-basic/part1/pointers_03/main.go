package main

import "fmt"

// Allocate address
func main() {
	p := new(int)
	*p = 7

	fmt.Println(*p)
}
