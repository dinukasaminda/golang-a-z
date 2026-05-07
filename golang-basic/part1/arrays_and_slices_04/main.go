package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50} // create slice

	t := s[1:3] // create slice with 1st till 3rd indexes : will copy 20,30 from backing same s
	// t := make([]int, 3)
	// copy(t, s)
	fmt.Println(t)

	t[0] = 99
	fmt.Println(t)
	fmt.Println(s)

}
