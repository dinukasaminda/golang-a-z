package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Label() string {
	return fmt.Sprint("User Name: ", u.Name, " Age: ", u.Age)
}

func (u *User) IncAge() {
	u.Age++
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", c)
}

func main() {
	u := &User{
		Name: "dinuka",
		Age:  31,
	}
	fmt.Println("User : ", u.Name, u.Age)

	fmt.Println(u.Label())

	u.IncAge()
	fmt.Println(u.Label())

	t := Celsius(27.8)
	fmt.Println(t.String())
}
