package main

import "fmt"

func main() {

	words := []string{"go", "go", "rust", "python"}

	counts := make(map[string]int)

	for _, w := range words {
		counts[w]++
	}

	fmt.Println(counts)
}
