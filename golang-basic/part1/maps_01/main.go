package main

import "fmt"

func main() {

	// Create and initialize a map
	mp := map[string]int{
		"A": 343,
		"B": 34,
	}

	fmt.Println(mp)

	// Make map
	mydata := make(map[string]int)
	mydata["Hello"] = 23
	mydata["Brawo"] = 43
	fmt.Println(mydata)

	v := mydata["Hello"]
	fmt.Println("value for Hello key: ", v)

	v, ok := mydata["X"]
	fmt.Println("Value : ", v)
	fmt.Println("Existance : ", ok)

	delete(mydata, "Hello")
	fmt.Println(mydata)

	mydata["Hello2"] = 46

	for k, v := range mydata {
		fmt.Println(k, v)
	}
}
