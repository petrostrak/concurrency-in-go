package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("%s <%s>", u.name, u.email)
}

func main() {
	u := user{
		name:  "Petros Trak",
		email: "pit.trak@gmail.com",
	}
	fmt.Println(u)
}
