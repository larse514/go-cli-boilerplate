package commands

import (
	"github.com/larse514/go-cli-boilerplate/actions"
	"github.com/urfave/cli"
)

//HelloAction interface to define hello command actions
type HelloAction interface {
	Hello(c *cli.Context) error
}

//Hello is a command to register a route to a service
func Hello() cli.Command {
	action := actions.SayHelloAction{}
	return cli.Command{

		Name:    "hello",
		Aliases: []string{"h"},
		Usage:   "use to say hello",
		Action:  action.SayHello,
	}
}
