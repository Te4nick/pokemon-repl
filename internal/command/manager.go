package command

import (
	"fmt"

	"github.com/chrxn1c/pokemon-repl/internal/entity"
	"github.com/chrxn1c/pokemon-repl/pkg/pokectx"
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
		case "explore":
			cmd.Callback = ExploreCallback
		case "inspect":
			cmd.Callback = InspectCallback
		case "catch":
			cmd.Callback = CatchCallback
		case "pokedex":
			cmd.Callback = PokedexCallback
		default:
			cmd.Callback = commander.notImplemented
		}
		commander.commands[cmd.Name] = cmd
	}

	return commander
}

func (c *Commander) notImplemented(_ *pokectx.Te4nickPokeCTX, _ string) (output string, err error) {
	return "not implemented", nil
}

func (c *Commander) helpCallback(_ *pokectx.Te4nickPokeCTX, arg string) (output string, err error) {
	if len(arg) > 0 {
		fmt.Println("Detected argument to help command which is not supported")
		return "", nil
	}
	helpStr := "\nFor now you can do the following:\n"
	i := 1
	for _, cmd := range c.commands {
		helpStr += fmt.Sprintf("%d) %s - %s\n", i, cmd.Name, cmd.Description)
		i++
	}
	return helpStr, nil
}

func (c *Commander) Exec(cmd, arg string, ctx *pokectx.Te4nickPokeCTX) (output string, err error) {
	command, ok := c.commands[cmd]
	if !ok {
		return "Given command is not supported. Use \"help\" if necessary.", nil
	}

	return command.Callback(ctx, arg)
}
