package command

import (
	"fmt"

	callbacks "github.com/chrxn1c/pokemon-repl/internal/command/callbacks"
	"github.com/chrxn1c/pokemon-repl/internal/entity"
	"github.com/chrxn1c/pokemon-repl/internal/user_context"
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
			cmd.Callback = commander.helpCallback()
		case "exit":
			cmd.Callback = callbacks.ExitCallback
		case "map":
			cmd.Callback = callbacks.MapCallback
		case "mapb":
			cmd.Callback = callbacks.MapbCallback
		default:
			cmd.Callback = commander.notImplemented()
		}
		commander.commands[cmd.Name] = cmd
	}

	return commander
}

func (c *Commander) notImplemented() entity.Callback {
	return func(ctx *user_context.UserContext) (output string, err error) {
		return "not implemented", nil
	}
}

func (c *Commander) helpCallback() entity.Callback {
	return func(_ *user_context.UserContext) (output string, err error) {
		helpStr := "\nFor now you can do the following:\n"
		i := 1
		for _, cmd := range c.commands {
			helpStr += fmt.Sprintf("%d) %s - %s\n", i, cmd.Name, cmd.Description)
			i++
		}
		return helpStr, nil
	}
}

func (c *Commander) unknownCommand() entity.Callback {
	return func(_ *user_context.UserContext) (output string, err error) {
		return "Given command is not supported. Use \"help\" if necessary.", nil
	}
}

func (c *Commander) Exec(cmd string, ctx *user_context.UserContext) (output string, err error) {
	command, ok := c.commands[cmd]
	if !ok {
		return c.unknownCommand()(ctx)
	}

	return command.Callback(ctx)
}
