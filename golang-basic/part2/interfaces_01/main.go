package main

import "fmt"

type Speaker interface{ Say() string }

func Greet(s Speaker) {
	fmt.Println(s.Say())
}

type Dog struct{}

func (Dog) Say() string {
	return "Baw"
}

type Robot struct{}

func (Robot) Say() string {
	return "Beep"
}

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func main() {
	Greet(Dog{})
	Greet(Robot{})

	fmt.Println(Point{3, 4})
}
