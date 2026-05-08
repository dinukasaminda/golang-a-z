package main

import "fmt"

func main() {
	// This is an slice
	a := []int{1, 2, 3}
	a = append(a, 5, 6, 7)
	fmt.Println(a)

	// 3 values with default
	// 8 capacity
	b := make([]int, 3, 8)

	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	c := [3]int{5, 6, 7}
	// c = append(c, 3)  <- getting runtime error , cannot add to array
	fmt.Println(c)

}
