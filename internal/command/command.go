package command

import (
	callbacks "github.com/chrxn1c/pokemon-repl/internal/command/callbacks"
	descriptions "github.com/chrxn1c/pokemon-repl/internal/command/descriptions"
	names "github.com/chrxn1c/pokemon-repl/internal/command/names"
	context "github.com/chrxn1c/pokemon-repl/internal/user_context"
)

type Command struct {
	Name        string
	description string
	Callback    func(ctx *context.UserContext) (output string, err error)
}

var exitCommand Command = Command{
	Name:        names.EXIT_NAME,
	description: descriptions.EXIT_DESCRIPTION,
	Callback:    callbacks.ExitCallback,
}

var helpCommand Command = Command{
	Name:        names.HELP_NAME,
	description: descriptions.HELP_DESCRIPTION,
	Callback:    callbacks.HelpCallback,
}

var mapCommand Command = Command{
	Name:        names.MAP_NAME,
	description: descriptions.MAP_DESCIPTION,
	Callback:    callbacks.MapCallback,
}

var mapbCommand Command = Command{
	Name:        names.MAPB_NAME,
	description: descriptions.MAPB_DESCIPTION,
	Callback:    callbacks.MapbCallback,
}

var Commands map[string]Command = map[string]Command{
	names.EXIT_NAME: exitCommand,
	names.HELP_NAME: helpCommand,
	names.MAP_NAME:  mapCommand,
	names.MAPB_NAME: mapbCommand,
}
