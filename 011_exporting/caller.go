package main

import (
	"fmt"

	"github.com/gusaki/learn-go/011_exporting/counter"
	"github.com/gusaki/learn-go/011_exporting/user"
)

func main() {
	alertCounter := counter.AlertCounter(10)
	fmt.Println("Imported from package counter, AlertCounter = ", alertCounter)

	u := user.User{
		Name: "Bob",
		ID:   12345,
	}
	fmt.Println("Imported from package user, user = ", u)
}
