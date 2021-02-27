package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// If we need more context on what happened in the error then
// we need custom error types. The json package's Unmarshal
// function returns such custom errors.

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NOTE
// Type-as-context
// This kind of implementation where the return error type is
// used to check for the kind of error leads to lots of cases
// and more coupled way of error handling.
//
// instead the better way to approach error handling is to use
// Behaviour-as-context as the approach
//
func main() {
	var u user

	data := `
{
  "name": "Bob",
  "email": "bob@gmail.com"
}
`
	// NOTE the correct code would pass &u instead of u below
	err := json.Unmarshal([]byte(data), u)
	if err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *json.InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
	fmt.Println("Unmarshal successful", u)
}
