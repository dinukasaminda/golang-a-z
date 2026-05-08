package main

import (
	"fmt"
	"strings"
)

func main() {
	// IF and IF Else
	c := 30

	if c > 10 {
		fmt.Println("C greater than 10!")
	}

	if s := strings.TrimSpace("  go  "); s == "go" {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed")
	}

	// Range on slice
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	// Range on map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Switch example 01

	cmd := "cancel"
	switch cmd {
	case "ping":
		{
			fmt.Println("PING")
			fmt.Println("PING...")
		}
	case "quit", "cancel":
		fmt.Println("Quit")
	default:
		fmt.Println("Unknown")
	}

	// Switch example 02
	x := 42
	switch {
	case x < 0:
		fmt.Println("neg")
	case x == 0:
		fmt.Println("zero")
	case x > 10:
		fmt.Println("x greater than 10")
	case x > 15:
		fmt.Println("x greater than 15")
	default:
		fmt.Println("pos")
	}

}
