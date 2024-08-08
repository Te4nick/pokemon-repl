package command

import (
	"fmt"

	"github.com/chrxn1c/pokemon-repl/internal/user_context"
)

type Commander struct {
	commands map[string]*Command
}

func NewCommander(cmds []*Command) *Commander {
	commander := &Commander{
		commands: make(map[string]*Command),
	}

	for _, cmd := range cmds {
		switch cmd.Name {
		case "help":
			cmd.Callback = commander.helpCallback
		case "exit":
			cmd.Callback = ExitCallback
		case "map":
			cmd.Callback = MapCallback
		case "mapb":
			cmd.Callback = MapbCallback
		default:
			cmd.Callback = commander.notImplemented()
		}
		commander.commands[cmd.Name] = cmd
	}

	return commander
}

func (c *Commander) notImplemented() Callback {
	return func(ctx *user_context.UserContext) (output string, err error) {
		return "not implemented", nil
	}
}

func (c *Commander) helpCallback(_ *user_context.UserContext) (output string, err error) {
	helpStr := "\nFor now you can do the following:\n"
	i := 1
	for _, cmd := range c.commands {
		helpStr += fmt.Sprintf("%d) %s - %s\n", i, cmd.Name, cmd.Description)
		i++
	}
	return helpStr, nil
}

func (c *Commander) Exec(cmd string, ctx *user_context.UserContext) (output string, err error) {
	command, ok := c.commands[cmd]
	if !ok {
		return "Given command is not supported. Use \"help\" if necessary.", nil
	}

	return command.Callback(ctx)
}
