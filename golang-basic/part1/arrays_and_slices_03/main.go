package main

import "fmt"

// Copy
func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 10)

	n := copy(dst, src)

	fmt.Println(n, dst)

}
