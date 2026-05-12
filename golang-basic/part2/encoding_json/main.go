package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	// Converting struct to json string
	u := User{Name: "Dinuka", Age: 30}

	b, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Converting json string to struct
	var v User
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}
	fmt.Println(v)

	// Converting to map
	dec := json.NewDecoder(bytes.NewReader([]byte(`{"n":1.5,"i":42}`)))
	dec.UseNumber()
	var m map[string]any
	if err := dec.Decode(&m); err != nil {
		panic(err)
	}
	i, err := m["i"].(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
