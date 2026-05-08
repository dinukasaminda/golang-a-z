package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("X is negative")
	}
	return math.Sqrt(x), nil
}

func main() {
	xc, err := sqrt(-10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(xc)
}
