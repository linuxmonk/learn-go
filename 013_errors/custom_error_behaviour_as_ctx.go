package main

import (
	"fmt"
	"io"
	"net"
)

// Example looking at the net package OpError type.
// this type implements interface like 'temporary'
// So instead of calling / checking if the error is
// of a specific type we can check if the error
// implements the 'temporary' interface.

func (c *client) TypeAsContext() {

	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.ParseError:
				fmt.Println("Parse Error")
			case *net.OpError:
				fmt.Println("Op Error")
			default:
				if err == io.EOF {
					fmt.Println("Client leaving chat. <EOF>")
				}
			}
		}
	}
}

func (c *client) BehaviourAsContext() {

	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case temporary:
				if !e.Temporary() {
					fmt.Println("Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					fmt.Println("Client leaving chat. <EOF>")
				}
			}
		}
	}
}
