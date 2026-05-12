package main

import (
	"fmt"
	"io"
	"strings"
)

func describe(i any) {
	_, ok := i.(int)

	if ok {
		fmt.Println("Type of Input is: Int")
		return
	}

	_, ok = i.(string)
	if ok {
		fmt.Println("Type of Input is String")
		return
	}

	fmt.Println("Type of Input is Unknown")

}

func describe_switch(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("Int:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Other")
	}
}
func main() {
	describe(10.9)
	describe("Hello")

	describe_switch(10.9)
	describe_switch("Hello")

	var r io.Reader = strings.NewReader("ABC")
	if sr, ok := r.(io.ReadSeeker); ok {
		fmt.Println("Support seek", sr)
	}

}
