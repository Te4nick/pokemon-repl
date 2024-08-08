package command

import (
	"fmt"

	"github.com/chrxn1c/pokemon-repl/internal/entity"
)

type Commander struct {
	commands map[string]*entity.Command
}

func NewCommander(cmds []*entity.Command) *Commander {
	commander := &Commander{
		commands: make(map[string]*entity.Command),
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

func (c *Commander) notImplemented() entity.Callback {
	return func(ctx *entity.UserContext) (output string, err error) {
		return "not implemented", nil
	}
}

func (c *Commander) helpCallback(_ *entity.UserContext) (output string, err error) {
	helpStr := "\nFor now you can do the following:\n"
	i := 1
	for _, cmd := range c.commands {
		helpStr += fmt.Sprintf("%d) %s - %s\n", i, cmd.Name, cmd.Description)
		i++
	}
	return helpStr, nil
}

func (c *Commander) Exec(cmd string, ctx *entity.UserContext) (output string, err error) {
	command, ok := c.commands[cmd]
	if !ok {
		return "Given command is not supported. Use \"help\" if necessary.", nil
	}

	return command.Callback(ctx)
}
