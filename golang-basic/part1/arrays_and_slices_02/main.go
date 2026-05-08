package main

import "fmt"

func main() {
	var arr [3]int // This is array
	arr[0] = 1

	fmt.Println(arr)

	s := arr[:] // This is slice view of array arr
	s = append(s, 4)
	fmt.Println(s)

}
