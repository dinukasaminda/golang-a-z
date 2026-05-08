package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str1 := "Hello"
	fmt.Println("Length of str1: ", len(str1))

	str2 := "こんにちは"
	fmt.Println("Length of str2: ", len(str2)) // This is counting all the bytes not the characters

	fmt.Println("Characters in str2: ", utf8.RuneCountInString(str2))

	// This is print each byte
	i := 0
	for i < len(str2) {
		fmt.Println("each byte: ", str2[i])
		i++
	}

	// range of str is giving each characters
	for i, r := range str2 {
		fmt.Printf("%d: %U %c\n", i, r, r)
	}

	// Building a string
	var b strings.Builder
	b.WriteString("a")
	b.WriteString("b")

	fmt.Println(b.String())

	// string from Runes
	rs := []rune{'A', 'Ω', '世'}
	fmt.Println(string(rs))

}
