package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	if u == nil {
		return "Nil Value"
	}
	return fmt.Sprintf("Name: %s, Email: %s", u.name, u.email)
}

func main() {
	u := user{
		name:  "Bob",
		email: "bob@mail.com",
	}
	fmt.Println(u)  // prints via default behaviour
	fmt.Println(&u) // prints via overridden stringer
}
