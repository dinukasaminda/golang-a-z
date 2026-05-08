package main

import "fmt"

// Pointers With Structs

type Point struct{ X, Y int }

func main() {
	p := &Point{
		X: 10,
		Y: 20,
	}
	fmt.Println(p.X, (*p).Y)
}
