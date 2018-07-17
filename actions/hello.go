package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

//SayHelloAction Implementation of HelloAction
type SayHelloAction struct {
}

//SayHello is a method to say hello to someone
func (action SayHelloAction) SayHello(c *cli.Context) error {
	fmt.Printf("Hello %q\n", c.Args().Get(0))
	return nil
}
