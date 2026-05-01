package main

import (
	"errors"
	"fmt"
)

func main() {
	add(3, 4)
	v, err := div(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of a/b:", v)
	}
	sum(1, 2, 3, 4, 5, 6, 7)

	var nslice = []int{3, 5, 7, 9}
	sum(nslice...)

	fmt.Println(split(16))

	// Function value

	var f func(int) int
	f = func(x int) int {
		return x * 2
	}

	fmt.Println(f(4))

}

func add(a, b int) {
	fmt.Println(a + b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cnanot division by zero")
	}
	return a / b, nil
}

// Variadic function
func sum(vals ...int) {
	s := 0

	for _, v := range vals {
		s += v
	}
	fmt.Println("Sum is: ", s)
}

// Named results
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
