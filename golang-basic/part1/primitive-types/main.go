package main

import "fmt"

func main() {
	var a int = 10
	var b int64 = int64(a)

	var x int = 3
	var y int = 2
	fmt.Println(x / y) // integer division

	fmt.Println(float64(x) / float64(y)) // floating point division

	fmt.Println(b)

	// Explicit width types
	// Explictly setting the width 16,32,64
	var u uint = 255
	var i int32 = -23
	fmt.Println(u)
	fmt.Println(i)

	// Boolean

	ok := true
	fmt.Println("Boolean : ", ok)
}
