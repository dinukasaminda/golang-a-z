package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello")
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, r)
	fmt.Println(buf.String())

	data, err := io.ReadAll(strings.NewReader("world"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//
	r2 := bufio.NewReader(strings.NewReader("hello,world"))
	part, err := r2.ReadString(',')
	if err != nil {
		panic(err)
	}
	fmt.Print(part)
}
